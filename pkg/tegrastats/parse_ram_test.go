package tegrastats

import (
	"testing"
)

func TestParseRAM(t *testing.T) {
	ParseTest(t, parseRAM, []TestParameter[TegraRAM]{
		{"RAM 707/7854MB", TegraRAM{InUse: 707, Total: 7854}},
		{"RAM 407/15823MB", TegraRAM{InUse: 407, Total: 15823}},
		{"RAM 2257/30536MB", TegraRAM{InUse: 2257, Total: 30536}},
		{"RAM 330/3964MB", TegraRAM{InUse: 330, Total: 3964}},
		{"RAM 4722/7844MB", TegraRAM{InUse: 4722, Total: 7844}},
		{"RAM 4937/14887MB", TegraRAM{InUse: 4937, Total: 14887}},
	})
}
