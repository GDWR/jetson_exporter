// https://docs.nvidia.com/jetson/archives/r34.1/DeveloperGuide/text/AT/JetsonLinuxDevelopmentTools/TegrastatsUtility.html#reported-statistics
package tegrastats

func ParseTegraStats(input string) (*TegraStats, error) {
	time, err := parseTime(input)
	if err != nil && err != errTimeNotFound {
		return nil, err
	}

	ram, err := parseRAM(input)
	if err != nil && err != errRAMNotFound {
		return nil, err
	}

	swap, err := parseSwap(input)
	if err != nil && err != errSwapNotFound {
		return nil, err
	}

	iram, err := parseIRAM(input)
	if err != nil && err != errIRAMNotFound {
		return nil, err
	}

	cpus, err := parseCPUs(input)
	if err != nil && err != errCPUsNotFound {
		return nil, err
	}

	emc, err := parseEMC(input)
	if err != nil && err != errEMCNotFound {
		return nil, err
	}

	gr3d, err := parseGR3D(input)
	if err != nil && err != errGR3DNotFound {
		return nil, err
	}

	temps, err := parseTemps(input)
	if err != nil && err != errTempsNotFound {
		return nil, err
	}

	return &TegraStats{
		Timestamp: time,
		RAM:       ram,
		Swap:      swap,
		IRAM:      iram,
		CPUs:      cpus,
		EMC:       emc,
		GR3D:      gr3d,
		Temps:     temps,
	}, nil
}
