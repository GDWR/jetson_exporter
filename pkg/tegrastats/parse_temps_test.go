package tegrastats

import (
	"testing"
)

func TestParseTemps(t *testing.T) {
	ParseTestArr(t, parseTemps, []TestParameterArr[TegraTemp]{
		{
			"AO@31C CPU@25C GPU@22C PLL@21.5C Tdiode@28.75C PMIC@100C Tboard@29C thermal@23.5C",
			[]TegraTemp{
				{"AO", 31},
				{"CPU", 25},
				{"GPU", 22},
				{"PLL", 21.5},
				{"Tdiode", 28.75},
				{"PMIC", 100},
				{"Tboard", 29},
				{"thermal", 23.5},
			},
		},
	})
}
