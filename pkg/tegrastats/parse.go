package tegrastats

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// https://docs.nvidia.com/drive/drive_os_5.1.6.1L/nvvib_docs/index.html#page/DRIVE_OS_Linux_SDK_Development_Guide/Utilities/util_tegrastats.html
var regexes = map[string]*regexp.Regexp{
	"time":       regexp.MustCompile(`\d+-\d+-\d+\s\d+:\d+:\d+`),
	"ram":        regexp.MustCompile(`RAM (\d+)/\d+MB`),
	"ramMax":     regexp.MustCompile(`RAM \d+/(\d+)MB`),
	"swap":       regexp.MustCompile(`SWAP (\d+)/\d+MB`),
	"swapMax":    regexp.MustCompile(`SWAP \d+/(\d)+MB`),
	"swapCached": regexp.MustCompile(`\(cached (\d+)MB\)`),
	"cpus":       regexp.MustCompile(`CPU \[(.+)\]`),
	"emcFreq":    regexp.MustCompile(`EMC_FREQ (\d+)%`),
	"gr3dFreq":   regexp.MustCompile(`GR3D_FREQ (\d+)%`),
	"cv0Temp":    regexp.MustCompile(`CV0@(-?[\-|\d|\.]+)C`),
	"cpuTemp":    regexp.MustCompile(`CPU@(-?[\d|\.]+)C`),
	"tboardTemp": regexp.MustCompile(`Tboard@(-?[\d|\.]+)C`),
	"soc2Temp":   regexp.MustCompile(`SOC2@(-?[\d|\.]+)C`),
	"diodeTemp":  regexp.MustCompile(`Tdiode@(-?[\d|\.]+)C`),
	"soc0Temp":   regexp.MustCompile(`SOC0@(-?[\d|\.]+)C`),
	"cv1Temp":    regexp.MustCompile(`CV1@(-?[\d|\.]+)C`),
	"gpuTemp":    regexp.MustCompile(`GPU@(-?[\d|\.]+)C`),
	"tjTemp":     regexp.MustCompile(`tj@(-?[\d|\.]+)C`),
	"soc1Temp":   regexp.MustCompile(`SOC1@(-?[\d|\.]+)C`),
	"cv2Temp":    regexp.MustCompile(`CV2@(-?[\d|\.]+)C`),
}

func ParseTegraStats(input string) (*TegraStats, error) {
	result := make(map[string]string)

	for k, v := range regexes {
		matches := v.FindStringSubmatch(input)
		result[k] = matches[len(matches)-1]
	}

	time, err := time.Parse("01-02-2006 15:04:05", result["time"])
	if err != nil {
		return nil, err
	}

	// parse result["ram"] and result["ramMax"] to int
	ram, err := strconv.Atoi(result["ram"])
	if err != nil {
		return nil, err
	}
	ramMax, err := strconv.Atoi(result["ramMax"])
	if err != nil {
		return nil, err
	}
	swap, err := strconv.Atoi(result["swap"])
	if err != nil {
		return nil, err
	}
	swapMax, err := strconv.Atoi(result["swapMax"])
	if err != nil {
		return nil, err
	}
	swapCached, err := strconv.Atoi(result["swapCached"])
	if err != nil {
		return nil, err
	}

	cpus := ParseTegraStatsCpus(result["cpus"])

	emcFreq, err := strconv.Atoi(result["emcFreq"])
	if err != nil {
		return nil, err
	}
	gr3dFreq, err := strconv.Atoi(result["gr3dFreq"])
	if err != nil {
		return nil, err
	}
	cv0Temp, err := strconv.ParseFloat(result["cv0Temp"], 64)
	if err != nil {
		return nil, err
	}
	cpuTemp, err := strconv.ParseFloat(result["cpuTemp"], 64)
	if err != nil {
		return nil, err
	}
	boardTemp, err := strconv.ParseFloat(result["tboardTemp"], 64)
	if err != nil {
		return nil, err
	}
	soc2Temp, err := strconv.ParseFloat(result["soc2Temp"], 64)
	if err != nil {
		return nil, err
	}
	diodeTemp, err := strconv.ParseFloat(result["diodeTemp"], 64)
	if err != nil {
		return nil, err
	}
	soc0Temp, err := strconv.ParseFloat(result["soc0Temp"], 64)
	if err != nil {
		return nil, err
	}
	cv1Temp, err := strconv.ParseFloat(result["cv1Temp"], 64)
	if err != nil {
		return nil, err
	}
	gpuTemp, err := strconv.ParseFloat(result["gpuTemp"], 64)
	if err != nil {
		return nil, err
	}
	tjTemp, err := strconv.ParseFloat(result["tjTemp"], 64)
	if err != nil {
		return nil, err
	}
	soc1Temp, err := strconv.ParseFloat(result["soc1Temp"], 64)
	if err != nil {
		return nil, err
	}
	cv2Temp, err := strconv.ParseFloat(result["cv2Temp"], 64)
	if err != nil {
		return nil, err
	}

	return &TegraStats{
		Timestamp:  time,
		Ram:        ram,
		RamMax:     ramMax,
		Swap:       swap,
		SwapMax:    swapMax,
		SwapCached: swapCached,
		Cpus:       cpus,
		EMCFreq:    emcFreq,
		GR3DFreq:   gr3dFreq,
		CV0Temp:    cv0Temp,
		CpuTemp:    cpuTemp,
		TBoardTemp: boardTemp,
		SOC2Temp:   soc2Temp,
		DiodeTemp:  diodeTemp,
		SOC0Temp:   soc0Temp,
		CV1Temp:    cv1Temp,
		GpuTemp:    gpuTemp,
		TjTemp:     tjTemp,
		Soc1Temp:   soc1Temp,
		CV2Temp:    cv2Temp,
	}, nil
}

func ParseTegraStatsCpus(input string) []TegraCpu {
	output := make([]TegraCpu, 0)

	rawCpus := strings.Split(input, ",")
	for i, rawCpu := range rawCpus {
		cpu := strings.Split(rawCpu, "@")
		percentage, _ := strconv.ParseFloat(strings.TrimSuffix(cpu[0], "%"), 64)
		output = append(output, TegraCpu{
			Core:       strconv.Itoa(i),
			Percentage: percentage,
		})
	}

	return output
}
