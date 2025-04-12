package tegrastats

import (
	"errors"
	"regexp"
	"strconv"
)

var tempRe = regexp.MustCompile(`(?P<name>\w+)@(?P<temp>-?\d*\.?\d+)C`)
var errTempsNotFound = errors.New("temps not found")

func parseTemps(input string) ([]TegraTemp, error) {
	results := tempRe.FindAllStringSubmatch(input, -1)
	if len(results) == 0 {
		return nil, errTempsNotFound
	}

	output := make([]TegraTemp, len(results))

	for i, result := range results {
		temp, err := strconv.ParseFloat(result[tempRe.SubexpIndex("temp")], 32)
		if err != nil {
			return nil, err
		}

		output[i] = TegraTemp{
			Name: result[tempRe.SubexpIndex("name")],
			Temp: float32(temp),
		}
	}

	return output, nil
}
