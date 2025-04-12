package tegrastats

import (
	"errors"
	"regexp"
	"strconv"
)

var gr3dRe = regexp.MustCompile(`GR3D_FREQ (?P<percentage>\d+)%@?\[?(?P<frequency>\d+)?\]?`)
var errGR3DNotFound = errors.New("gr3d not found")

func parseGR3D(input string) (*TegraGR3D, error) {
	result := gr3dRe.FindStringSubmatch(input)

	if len(result) == 0 {
		return nil, errGR3DNotFound
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

	return &TegraGR3D{
		Percentage: percentage,
		Frequency:  frequency,
	}, nil
}
