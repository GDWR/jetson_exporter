package tegrastats

import (
	"testing"
)

func TestParseIRAM(t *testing.T) {
	ParseTest(t, parseIRAM, []TestParameter[TegraIRAM]{
		{"IRAM 0/252kB", TegraIRAM{InUse: 0, Total: 252}},
	})
}
