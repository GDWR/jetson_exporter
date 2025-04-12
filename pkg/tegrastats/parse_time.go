package tegrastats

import (
	"errors"
	"regexp"
	"time"
)

var timeRe = regexp.MustCompile(`\d+-\d+-\d+\s\d+:\d+:\d+`)
var errTimeNotFound = errors.New("time not found")

func parseTime(input string) (*time.Time, error) {
	result := timeRe.FindString(input)
	if len(result) == 0 {
		return nil, errTimeNotFound
	}

	time, err := time.Parse("01-02-2006 15:04:05", result)
	if err != nil {
		return nil, err
	}
	return &time, nil
}
