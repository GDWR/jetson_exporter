package main

import (
	"bufio"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"

	"github.com/gdwr/jetson_exporter/pkg/tegrastats"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
)

var (
	tegrastatsTemperature = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "temperature_celsius",
			Help:      "Processor block temperature in degrees Celsius",
		},
		[]string{"block"},
	)

	tegrastatsCPULoad = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "cpu_load_percentage",
			Help:      "Load on each CPU core relative to the current running frequency, or off if a core is currently powered down",
		},
		[]string{"core"},
	)
	tegrastatsCPUFrequency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "cpu_frequency_megahertz",
			Help:      "CPU frequency in megahertz. Goes up or down dynamically depending on the CPU workload",
		},
		[]string{"core"},
	)

	tegrastatsRAMUsed = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "ram_used_megabytes",
			Help:      "Amount of RAM in use, specified in megabytes",
		},
	)
	tegrastatsRAMTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "ram_total_megabytes",
			Help:      "Total amount of RAM available for applications",
		},
	)

	tegrastatsSwapInUse = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "swap_used_megabytes",
			Help:      "Amount of SWAP in use, in megabytes",
		},
	)
	tegrastatsSwapTotal = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "swap_total_megabytes",
			Help:      "Total amount of SWAP available for applications",
		},
	)
	tegrastatsSwapCached = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "swap_cached_megabytes",
			Help:      "Amount of SWAP cached",
		},
	)

	tegrastatsGR3DUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "gr3d_usage_percentage",
			Help:      "Percentage of the GR3D that in use, relative to the current running frequency",
		},
	)
	tegrastatsGR3DFrequency = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "tegrastats",
			Name:      "gr3d_frequency_megahertz",
			Help:      "GR3D frequency in megahertz",
		},
	)
)

func main() {
	var (
		webConfig      = webflag.AddFlags(kingpin.CommandLine, ":9102")
		metricsPath    = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
		tegrastatsPath = kingpin.Flag("tegrastats.path", "Path to tegrastats binary, if not provided will attempt to use $PATH").Default("tegrastats").String()
	)

	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print("jetson-exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	prometheus.Register(tegrastatsTemperature)
	prometheus.Register(tegrastatsCPULoad)
	prometheus.Register(tegrastatsCPUFrequency)
	prometheus.Register(tegrastatsRAMUsed)
	prometheus.Register(tegrastatsRAMTotal)
	prometheus.Register(tegrastatsSwapInUse)
	prometheus.Register(tegrastatsSwapTotal)
	prometheus.Register(tegrastatsSwapCached)
	prometheus.Register(tegrastatsGR3DUsage)
	prometheus.Register(tegrastatsGR3DFrequency)

	level.Info(logger).Log("msg", "Starting jetson-exporter", "tegraPath", *tegrastatsPath)

	go func() {
		cmd := exec.Command(*tegrastatsPath)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			level.Error(logger).Log("Can't find tegrastats")
			panic(err)
		}

		err = cmd.Start()
		if err != nil {
			level.Error(logger).Log("Can't find tegrastats")
			panic(err)
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			content := scanner.Text()
			stats, err := tegrastats.ParseTegraStats(content)
			if err != nil {
				level.Error(logger).
					Log("msg", "Error parsing tegrastats, provide this log in a Github Issue for resolution. https://github.com/GDWR/jetson-exporter/issues/new",
						"content", content,
						"err", err)
				continue
			}
			level.Debug(logger).Log("message", "updating metrics")

			for _, temp := range stats.Temps {
				tegrastatsTemperature.WithLabelValues(temp.Name).Set(float64(temp.Temp))
			}

			if stats.RAM != nil {
				tegrastatsRAMUsed.Set(float64(stats.RAM.InUse))
				tegrastatsRAMTotal.Set(float64(stats.RAM.Total))
			}

			if stats.Swap != nil {
				tegrastatsSwapInUse.Set(float64(stats.Swap.InUse))
				tegrastatsSwapTotal.Set(float64(stats.Swap.Total))
				tegrastatsSwapCached.Set(float64(stats.Swap.Cached))
			}

			if stats.GR3D != nil {
				tegrastatsGR3DUsage.Set(float64(stats.GR3D.Percentage))
				tegrastatsGR3DFrequency.Set(float64(stats.GR3D.Frequency))
			}

			for i, cpu := range stats.CPUs {
				iStr := strconv.Itoa(i)
				tegrastatsCPULoad.WithLabelValues(iStr).Set(float64(cpu.Percentage))
				tegrastatsCPUFrequency.WithLabelValues(iStr).Set(float64(cpu.Frequency))
			}
		}
	}()

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Jetson Exporter</title></head>
             <body>
             <h1>Jetson Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
			 </body>
             </html>`))
	})

	srv := &http.Server{}
	if err := web.ListenAndServe(srv, webConfig, logger); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}
