package tegrastats

import (
	"errors"
	"regexp"
	"strconv"
)

var swapRe = regexp.MustCompile(`SWAP (?P<inUse>\d+)/(?P<total>\d+)MB \(cached (?P<cached>\d+)MB\)`)
var errSwapNotFound = errors.New("swap not found")

func parseSwap(input string) (*TegraSwap, error) {
	result := swapRe.FindStringSubmatch(input)
	if len(result) == 0 {
		return nil, errSwapNotFound
	}

	inUse, err := strconv.Atoi(result[swapRe.SubexpIndex("inUse")])
	if err != nil {
		return nil, err
	}

	total, err := strconv.Atoi(result[swapRe.SubexpIndex("total")])
	if err != nil {
		return nil, err
	}

	cached, err := strconv.Atoi(result[swapRe.SubexpIndex("cached")])
	if err != nil {
		return nil, err
	}

	return &TegraSwap{
		InUse:  inUse,
		Total:  total,
		Cached: cached,
	}, nil
}
