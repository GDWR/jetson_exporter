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
		{
			"04-11-2025 10:36:27 RAM 4321/62801MB (lfb 1759x4MB) SWAP 19/31400MB (cached 0MB) CPU [0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729] EMC_FREQ 0%@2133 GR3D_FREQ 0%@[0,0] VIC_FREQ 729 APE 174 CV0@-256C CPU@44.562C Tboard@32C SOC2@40.093C Tdiode@33C SOC0@41.093C CV1@-256C GPU@-256C tj@44.343C SOC1@41.25C CV2@-256C VDD_GPU_SOC 2340mW/2340mW VDD_CPU_CV 390mW/390mW VIN_SYS_5V0 4025mW/4025mW NC 0mW/0mW VDDQ_VDD2_1V8AO 301mW/301mW NC 0mW/0mW",
			[]TegraCPU{
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
				{Percentage: 0, Frequency: 729},
			},
		},
	})
}
