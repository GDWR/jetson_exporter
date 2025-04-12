package tegrastats

import (
	"testing"
)

func TestParseEMC(t *testing.T) {
	ParseTest(t, parseEMC, []TestParameter[TegraEMC]{
		{"EMC_FREQ 0%", TegraEMC{Percentage: 0, Frequency: 0}},
		{"EMC_FREQ 2%@1866", TegraEMC{Percentage: 2, Frequency: 1866}},
		{"EMC_FREQ 0%@1600", TegraEMC{Percentage: 0, Frequency: 1600}},
		{"EMC_FREQ 1%@665", TegraEMC{Percentage: 1, Frequency: 665}},
		{"EMC_FREQ 0%@2133", TegraEMC{Percentage: 0, Frequency: 2133}},
		{"EMC_FREQ 0%@204", TegraEMC{Percentage: 0, Frequency: 204}},
		{"EMC_FREQ 3%@204", TegraEMC{Percentage: 3, Frequency: 204}},
		{"EMC_FREQ 2%@1866", TegraEMC{Percentage: 2, Frequency: 1866}},
	})
}
