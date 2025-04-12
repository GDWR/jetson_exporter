package tegrastats

import (
	"time"
)

// https://docs.nvidia.com/jetson/archives/r34.1/DeveloperGuide/text/AT/JetsonLinuxDevelopmentTools/TegrastatsUtility.html#reported-statistics
type TegraStats struct {
	Timestamp *time.Time
	RAM       *TegraRAM
	Swap      *TegraSwap
	IRAM      *TegraIRAM
	CPUs      []TegraCPU
	EMC       *TegraEMC
	GR3D      *TegraGR3D
	Temps     []TegraTemp
}

type TegraRAM struct {
	InUse int
	Total int
}

type TegraSwap struct {
	InUse  int
	Total  int
	Cached int
}

type TegraIRAM struct {
	InUse int
	Total int
}

type TegraCPU struct {
	Percentage int
	Frequency  int
}

type TegraEMC struct {
	Percentage int
	Frequency  int
}

type TegraGR3D struct {
	Percentage int
	Frequency  int
}

type TegraTemp struct {
	Name string
	Temp float32
}
