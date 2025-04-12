package tegrastats

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var cpusRe = regexp.MustCompile(`CPU \[(.+)\]`)
var cpuRe = regexp.MustCompile(`(?P<percentage>\d+)%@(?P<frequency>\d+)`)
var errCPUsNotFound = errors.New("cpus not found")

func parseCPUs(input string) ([]TegraCPU, error) {
	result := cpusRe.FindStringSubmatch(input)
	if len(result) == 0 {
		return nil, errCPUsNotFound
	}

	rawCpus := strings.Split(input, ",")
	output := make([]TegraCPU, len(rawCpus))
	for i, rawCpu := range rawCpus {
		result := cpuRe.FindStringSubmatch(rawCpu)

		percentage, err := strconv.Atoi(result[cpuRe.SubexpIndex("percentage")])
		if err != nil {
			return nil, err
		}

		freq, err := strconv.Atoi(result[cpuRe.SubexpIndex("frequency")])
		if err != nil {
			return nil, err
		}

		output[i] = TegraCPU{
			Percentage: percentage,
			Frequency:  freq,
		}
	}
	return output, nil
}
