package main

import (
	"bufio"
	"net/http"
	"os"
	"os/exec"

	"github.com/gdwr/jetson_exporter/pkg/tegrastats"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
)

var (
	ramUsageMegabytes = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_ram_usage_megabytes",
		Help: "The current RAM usage in megabytes",
	})
	ramMaxMegabytes = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_ram_max_megabytes",
		Help: "The maximum RAM in megabytes",
	})
	swapUsageMegabytes = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_swap_usage_megabytes",
		Help: "The current SWAP usage in megabytes",
	})
	swapMaxMegabytes = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_swap_max_megabytes",
		Help: "The maximum SWAP in megabytes",
	})
	swapCachedMegabytes = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_swap_cached_megabytes",
		Help: "The current SWAP cached in megabytes",
	})
	emcFreqPercent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_emc_freq_percent",
		Help: "The current EMC frequency in percent",
	})
	gr3dFreqPercent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_gr3d_freq_percent",
		Help: "The current GR3D frequency in percent",
	})
	cpuTempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_cpu_temp_celsius",
		Help: "The current temperature of the CPU in celsius",
	})
	tboardTempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_tboard_temp_celsius",
		Help: "The current temperature of the T(egra)Board in celsius",
	})
	soc2TempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_soc2_temp_celsius",
		Help: "The current temperature of the SOC2 in celsius",
	})
	diodeTempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_diode_temp_celsius",
		Help: "The current temperature of the diode in celsius",
	})
	soc0TempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_soc0_temp_celsius",
		Help: "The current temperature of the SOC0 in celsius",
	})
	cv1TempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_cv1_temp_celsius",
		Help: "The current temperature of the CV1 in celsius",
	})
	gpuTempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_gpu_temp_celsius",
		Help: "The current temperature of the GPU in celsius",
	})
	tjTempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_tj_temp_celsius",
		Help: "The current temperature of the TJ in celsius",
	})
	soc1TempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_soc1_temp_celsius",
		Help: "The current temperature of the SOC1 in celsius",
	})
	cv2TempCelsius = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tegrastats_cv2_temp_celsius",
		Help: "The current temperature of the CV2 in celsius",
	})
)

func main() {
	logger := promlog.New(&promlog.Config{})

	go func() {
		cmd := exec.Command("tegrastats")

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
			ramUsageMegabytes.Set(float64(stats.Ram))
			ramMaxMegabytes.Set(float64(stats.RamMax))
			swapUsageMegabytes.Set(float64(stats.Swap))
			swapMaxMegabytes.Set(float64(stats.SwapMax))
			swapCachedMegabytes.Set(float64(stats.SwapCached))
			emcFreqPercent.Set(float64(stats.EMCFreq))
			gr3dFreqPercent.Set(float64(stats.GR3DFreq))
			cpuTempCelsius.Set(stats.CpuTemp)
			tboardTempCelsius.Set(stats.TBoardTemp)
			soc2TempCelsius.Set(stats.SOC2Temp)
			diodeTempCelsius.Set(stats.DiodeTemp)
			soc0TempCelsius.Set(stats.SOC0Temp)
			cv1TempCelsius.Set(stats.CV1Temp)
			gpuTempCelsius.Set(stats.GpuTemp)
			tjTempCelsius.Set(stats.TjTemp)
			soc1TempCelsius.Set(stats.Soc1Temp)
			cv2TempCelsius.Set(stats.CV2Temp)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Jetson Exporter</title></head>
             <body>
             <h1>Jetson Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})

	addr := ":9012"
	level.Info(logger).Log("HTTP server listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}
