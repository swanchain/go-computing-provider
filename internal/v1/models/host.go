package models

type HostInfo struct {
	SwanProviderVersion string `json:"swan_miner_version"`
	OperatingSystem     string `json:"operating_system"`
	Architecture        string `json:"architecture"`
	CPUCores            int    `json:"cpu_cores"`
}
