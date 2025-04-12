package tegrastats

import (
	"testing"
)

func TestParseCPUs(t *testing.T) {
	ParseTestArr(t, parseCPUs, []TestParameterArr[TegraCPU]{
		{
			"CPU [8%@2265,8%@2265,17%@2265,7%@2265,2%@2265,3%@2265,30%@2265,17%@2265]",
			[]TegraCPU{
				{Percentage: 8, Frequency: 2265},
				{Percentage: 8, Frequency: 2265},
				{Percentage: 17, Frequency: 2265},
				{Percentage: 7, Frequency: 2265},
				{Percentage: 2, Frequency: 2265},
				{Percentage: 3, Frequency: 2265},
				{Percentage: 30, Frequency: 2265},
				{Percentage: 17, Frequency: 2265},
			},
		},
		{
			"CPU [17%@102,43%@102,59%@102,63%@102]",
			[]TegraCPU{
				{Percentage: 17, Frequency: 102},
				{Percentage: 43, Frequency: 102},
				{Percentage: 59, Frequency: 102},
				{Percentage: 63, Frequency: 102},
			},
		},
	})
}
