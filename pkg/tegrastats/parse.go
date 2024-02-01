package tegrastats

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// https://docs.nvidia.com/drive/drive_os_5.1.6.1L/nvvib_docs/index.html#page/DRIVE_OS_Linux_SDK_Development_Guide/Utilities/util_tegrastats.html
var regex = regexp.MustCompile(
	`(?P<time>\d+-\d+-\d+\s\d+:\d+:\d+) ` +
		`RAM (?P<ram>\d+)/(?P<ramMax>\d+)MB ` +
		`\(lfb \d+x\d+MB\) ` +
		`SWAP (?P<swap>\d+)/(?P<swapMax>\d+)MB ` +
		`\(cached (?P<swapCached>\d+)MB\) ` +
		`CPU \[(?P<cpus>.+)\] ` +
		`EMC_FREQ (?P<emcFreq>\d+)% ` +
		`GR3D_FREQ (?P<gr3dFreq>\d+)% ` +
		`CV0@(?P<cv0Temp>[\-|\d|\.]+)C ` +
		`CPU@(?P<cpuTemp>[\d|\.]+)C ` +
		`Tboard@(?P<tboardTemp>[\d|\.]+)C ` +
		`SOC2@(?P<soc2Temp>[\d|\.]+)C ` +
		`Tdiode@(?P<diodeTemp>[\d|\.]+)C ` +
		`SOC0@(?P<soc0Temp>[\d|\.]+)C ` +
		`CV1@(?P<cv1Temp>[\-|\d|\.]+)C ` +
		`GPU@(?P<gpuTemp>[\d|\.]+)C ` +
		`tj@(?P<tjTemp>[\d|\.]+)C ` +
		`SOC1@(?P<soc1Temp>[\d|\.]+)C ` +
		`CV2@(?P<cv2Temp>[\-|\d|\.]+)C`,
)

func ParseTegraStats(input string) (*TegraStats, error) {
	result := make(map[string]string)
	match := regex.FindStringSubmatch(input)
	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			if len(match) == 0 {
				return nil, fmt.Errorf("no match found for %v", name)
			}

			result[name] = match[i]
		}
	}

	time, err := time.Parse("01-02-2006 15:04:05", result["time"])
	if err != nil {
		return nil, err
	}

	// parse result["ram"] and result["ramMax"] to int
	ram, _ := strconv.Atoi(result["ram"])
	if err != nil {
		return nil, err
	}
	ramMax, _ := strconv.Atoi(result["ramMax"])
	if err != nil {
		return nil, err
	}
	swap, _ := strconv.Atoi(result["swap"])
	if err != nil {
		return nil, err
	}
	swapMax, _ := strconv.Atoi(result["swapMax"])
	if err != nil {
		return nil, err
	}
	swapCached, _ := strconv.Atoi(result["swapCached"])
	if err != nil {
		return nil, err
	}

	cpus := ParseTegraStatsCpus(result["cpus"])
	if err != nil {
		return nil, err
	}
	emcFreq, _ := strconv.Atoi(result["emcFreq"])
	if err != nil {
		return nil, err
	}
	gr3dFreq, _ := strconv.Atoi(result["gr3dFreq"])
	if err != nil {
		return nil, err
	}
	cv0Temp, _ := strconv.ParseFloat(result["cv0Temp"], 64)
	if err != nil {
		return nil, err
	}
	cpuTemp, _ := strconv.ParseFloat(result["cpuTemp"], 64)
	if err != nil {
		return nil, err
	}
	boardTemp, _ := strconv.ParseFloat(result["tboardTemp"], 64)
	if err != nil {
		return nil, err
	}
	soc2Temp, _ := strconv.ParseFloat(result["soc2Temp"], 64)
	if err != nil {
		return nil, err
	}
	diodeTemp, _ := strconv.ParseFloat(result["diodeTemp"], 64)
	if err != nil {
		return nil, err
	}
	soc0Temp, _ := strconv.ParseFloat(result["soc0Temp"], 64)
	if err != nil {
		return nil, err
	}
	cv1Temp, _ := strconv.ParseFloat(result["cv1Temp"], 64)
	if err != nil {
		return nil, err
	}
	gpuTemp, _ := strconv.ParseFloat(result["gpuTemp"], 64)
	if err != nil {
		return nil, err
	}
	tjTemp, _ := strconv.ParseFloat(result["tjTemp"], 64)
	if err != nil {
		return nil, err
	}
	soc1Temp, _ := strconv.ParseFloat(result["soc1Temp"], 64)
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
	return output
}
