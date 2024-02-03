package tegrastats

import (
	"fmt"
	"time"
)

type TegraStats struct {
	Timestamp                 time.Time
	Ram, RamMax               int
	Swap, SwapMax, SwapCached int
	Cpus                      []TegraCpu
	EMCFreq                   int
	GR3DFreq                  int
	CpuTemp                   float64
	TBoardTemp                float64
	CV0Temp                   float64
	SOC2Temp                  float64
	DiodeTemp                 float64
	SOC0Temp                  float64
	CV1Temp                   float64
	GpuTemp                   float64
	TjTemp                    float64
	Soc1Temp                  float64
	CV2Temp                   float64
}

func (ts TegraStats) String() string {
	return fmt.Sprintf(
		"%v RAM(%vMB/%vMB) SWAP(%vMB/%vMB cached=%vMB) CPUS(%v) EMC_FREQ(%v%%) GR3D_FREQ(%v%%) CV0_TEMP(%vC) CPU_TEMP(%vC) BOARD_TEMP(%vC) SOC2_TEMP(%vC) DIODE_TEMP(%vC) SOC0_TEMP(%vC) CV1_TEMP(%vC) GPU_TEMP(%vC) TJ_TEMP(%vC) SOC1_TEMP(%vC) CV2_TEMP(%vC)",
		ts.Timestamp, ts.Ram, ts.RamMax, ts.Swap, ts.SwapMax, ts.SwapCached, ts.Cpus, ts.EMCFreq, ts.GR3DFreq, ts.CV0Temp, ts.CpuTemp, ts.TBoardTemp, ts.SOC2Temp, ts.DiodeTemp, ts.SOC0Temp, ts.CV1Temp, ts.GpuTemp, ts.TjTemp, ts.Soc1Temp, ts.CV2Temp)
}

type TegraCpu struct {
	Core       string
	Percentage float64
}

func (tc TegraCpu) String() string {
	return fmt.Sprintf("%v@%v%%", tc.Core, tc.Percentage)
}
