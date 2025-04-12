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
		{
			"04-11-2025 10:36:27 RAM 4321/62801MB (lfb 1759x4MB) SWAP 19/31400MB (cached 0MB) CPU [0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729,0%@729] EMC_FREQ 0%@2133 GR3D_FREQ 0%@[0,0] VIC_FREQ 729 APE 174 CV0@-256C CPU@44.562C Tboard@32C SOC2@40.093C Tdiode@33C SOC0@41.093C CV1@-256C GPU@-256C tj@44.343C SOC1@41.25C CV2@-256C VDD_GPU_SOC 2340mW/2340mW VDD_CPU_CV 390mW/390mW VIN_SYS_5V0 4025mW/4025mW NC 0mW/0mW VDDQ_VDD2_1V8AO 301mW/301mW NC 0mW/0mW",
			[]TegraTemp{
				{"CV0", -256},
				{"CPU", 44.562},
				{"Tboard", 32},
				{"SOC2", 40.093},
				{"Tdiode", 33},
				{"SOC0", 41.093},
				{"CV1", -256},
				{"GPU", -256},
				{"tj", 44.343},
				{"SOC1", 41.25},
				{"CV2", -256},
			},
		},
	})
}
