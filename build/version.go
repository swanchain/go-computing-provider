package build

import (
	_ "embed"
	"encoding/json"
	"log"
)

var CurrentCommit string

var NetWorkTag string

const BuildVersion = "0.6.2"

const UBITaskImageIntelCpu = "filswan/ubi-worker-cpu-intel:v2.0"
const UBITaskImageIntelGpu = "filswan/ubi-worker-gpu-intel:v2.0"
const UBITaskImageAmdCpu = "filswan/ubi-worker-cpu-amd:v2.0"
const UBITaskImageAmdGpu = "filswan/ubi-worker-gpu-amd:v2.0"
const UBIResourceExporterDockerImage = "filswan/resource-exporter:v11.2.8"

func UserVersion() string {
	return BuildVersion + "+" + NetWorkTag + CurrentCommit
}

//go:embed parameters.json
var netWorkConfig string

func LoadParam() []NetworkConfig {
	var config []NetworkConfig
	err := json.Unmarshal([]byte(netWorkConfig), &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

type NetworkConfig struct {
	Network string `json:"network"`
	Config  struct {
		SequencerUrl                   string `json:"sequencer_url"`
		ZkEnginePk                     string `json:"zk_engine_pk"`
		OrchestratorPk                 string `json:"orchestrator_pk"`
		ChainRpc                       string `json:"chain_rpc"`
		SwanTokenContract              string `json:"swan_token_contract"`
		OrchestratorCollateralContract string `json:"orchestrator_collateral_contract"`
		JobManagerContract             string `json:"job_manager_contract"`
		JobManagerContractCreated      uint64 `json:"job_manager_contract_created"`
		RegisterCpContract             string `json:"register_cp_contract"`
		ZkCollateralContract           string `json:"zk_collateral_contract"`
		RegisterTaskContract           string `json:"register_task_contract"`
		SequencerContract              string `json:"sequencer_contract"`
	} `json:"config"`
}
