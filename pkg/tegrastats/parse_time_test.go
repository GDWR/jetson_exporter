package tegrastats

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	ParseTest(t, parseTime, []TestParameter[time.Time]{
		{"04-11-2025 19:52:25", time.Date(2025, 04, 11, 19, 52, 25, 0, time.UTC)},
		{"01-03-2023 16:10:22", time.Date(2023, 01, 03, 16, 10, 22, 0, time.UTC)},
	})
}
