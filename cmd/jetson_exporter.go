package main

import (
	"bufio"
	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	"net/http"
	"os"
	"os/exec"

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
			Name: "tegrastats_temperature_celsius",
			Help: "temperature in celsius gathered from tegrastats",
		},
		[]string{"component"},
	)
	tegrastatsRam = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tegrastats_ram_megabytes",
			Help: "ram metrics in megabytes gathered from tegrastats",
		},
		[]string{"type"},
	)
	tegrastatsSwap = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tegrastats_swap_megabytes",
			Help: "swap metrics in megabytes gathered from tegrastats",
		},
		[]string{"type"},
	)
	tegrastatsCpu = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tegrastats_cpu_percentage",
			Help: "cpu metrics in percentage gathered from tegrastats",
		},
		[]string{"core"},
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
	prometheus.Register(tegrastatsRam)
	prometheus.Register(tegrastatsSwap)
	prometheus.Register(tegrastatsCpu)

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

			tegrastatsTemperature.WithLabelValues("cpu").Set(stats.CpuTemp)
			tegrastatsTemperature.WithLabelValues("gpu").Set(stats.GpuTemp)
			tegrastatsTemperature.WithLabelValues("tboard").Set(stats.TBoardTemp)
			tegrastatsTemperature.WithLabelValues("diode").Set(stats.DiodeTemp)
			tegrastatsTemperature.WithLabelValues("tj").Set(stats.TjTemp)
			tegrastatsTemperature.WithLabelValues("soc0").Set(stats.SOC0Temp)
			tegrastatsTemperature.WithLabelValues("soc1").Set(stats.Soc1Temp)
			tegrastatsTemperature.WithLabelValues("soc2").Set(stats.SOC2Temp)
			tegrastatsTemperature.WithLabelValues("cv1").Set(stats.CV1Temp)
			tegrastatsTemperature.WithLabelValues("cv2").Set(stats.CV2Temp)

			tegrastatsRam.WithLabelValues("usage").Set(float64(stats.Ram))
			tegrastatsRam.WithLabelValues("max").Set(float64(stats.RamMax))

			tegrastatsSwap.WithLabelValues("current").Set(float64(stats.Swap))
			tegrastatsSwap.WithLabelValues("cached").Set(float64(stats.SwapCached))
			tegrastatsSwap.WithLabelValues("max").Set(float64(stats.SwapMax))

			for _, cpu := range stats.Cpus {
				tegrastatsCpu.WithLabelValues(cpu.Core).Set(cpu.Percentage)
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
