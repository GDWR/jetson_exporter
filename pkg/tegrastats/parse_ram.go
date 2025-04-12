package tegrastats

import (
	"errors"
	"regexp"
	"strconv"
)

var ramRe = regexp.MustCompile(`RAM (?P<inUse>\d+)/(?P<total>\d+)MB`)
var errRAMNotFound = errors.New("ram not found")

func parseRAM(input string) (*TegraRAM, error) {
	result := ramRe.FindStringSubmatch(input)
	if len(result) == 0 {
		return nil, errRAMNotFound
	}

	inUse, err := strconv.Atoi(result[ramRe.SubexpIndex("inUse")])
	if err != nil {
		return nil, err
	}

	total, err := strconv.Atoi(result[ramRe.SubexpIndex("total")])
	if err != nil {
		return nil, err
	}

	return &TegraRAM{
		InUse: inUse,
		Total: total,
	}, nil
}
