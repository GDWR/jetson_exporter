package tegrastats

import (
	"errors"
	"regexp"
	"strconv"
)

var emcRe = regexp.MustCompile(`EMC_FREQ (?P<percentage>\d+)%@?(?P<frequency>\d+)?`)
var errEMCNotFound = errors.New("emc not found")

func parseEMC(input string) (*TegraEMC, error) {
	result := emcRe.FindStringSubmatch(input)

	if len(result) == 0 {
		return nil, errEMCNotFound
	}

	percentage, err := strconv.Atoi(result[emcRe.SubexpIndex("percentage")])
	if err != nil {
		return nil, err
	}

	frequency := 0
	if result[emcRe.SubexpIndex("frequency")] != "" {
		frequency, err = strconv.Atoi(result[emcRe.SubexpIndex("frequency")])
		if err != nil {
			return nil, err
		}
	}

	return &TegraEMC{
		Percentage: percentage,
		Frequency:  frequency,
	}, nil
}
