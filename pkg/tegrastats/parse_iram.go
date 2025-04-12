package tegrastats

import (
	"errors"
	"regexp"
	"strconv"
)

var iramRe = regexp.MustCompile(`IRAM (?P<inUse>\d+)/(?P<total>\d+)kB`)
var errIRAMNotFound = errors.New("iram not found")

func parseIRAM(input string) (*TegraIRAM, error) {
	result := iramRe.FindStringSubmatch(input)

	if len(result) == 0 {
		return nil, errIRAMNotFound
	}

	inUse, err := strconv.Atoi(result[iramRe.SubexpIndex("inUse")])
	if err != nil {
		return nil, err
	}

	total, err := strconv.Atoi(result[iramRe.SubexpIndex("total")])
	if err != nil {
		return nil, err
	}

	return &TegraIRAM{
		InUse: inUse,
		Total: total,
	}, nil
}
