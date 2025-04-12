package tegrastats

import (
	"testing"
)

func TestParseGR3D(t *testing.T) {
	ParseTest(t, parseGR3D, []TestParameter[TegraGR3D]{
		{"GR3D_FREQ 77%", TegraGR3D{Percentage: 77, Frequency: 0}},
		{"GR3D_FREQ 5%@1300", TegraGR3D{Percentage: 5, Frequency: 1300}},
		{"GR3D_FREQ 3%@998", TegraGR3D{Percentage: 3, Frequency: 998}},
		{"GR3D_FREQ 23%@140", TegraGR3D{Percentage: 23, Frequency: 140}},
		{"GR3D_FREQ 0%@1377", TegraGR3D{Percentage: 0, Frequency: 1377}},
		{"GR3D_FREQ 60%@0", TegraGR3D{Percentage: 60, Frequency: 0}},
		{"GR3D_FREQ 0%@[1377]", TegraGR3D{Percentage: 0, Frequency: 1377}},
	})
}
