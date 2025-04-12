package tegrastats

import (
	"testing"
)

func TestParseSwap(t *testing.T) {
	ParseTest(t, parseSwap, []TestParameter[TegraSwap]{
		{"SWAP 494/7443MB (cached 1MB)", TegraSwap{InUse: 494, Total: 7443, Cached: 1}},
	})
}
