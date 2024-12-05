package build

import (
	_ "embed"
	"encoding/json"
	"log"
)

var CurrentCommit string

var NetWorkTag string

const BuildVersion = "0.7.1"

const UBITaskImageIntelCpu = "filswan/ubi-worker-cpu-intel:latest"
const UBITaskImageIntelGpu = "filswan/ubi-worker-gpu-intel:latest"
const UBITaskImageAmdCpu = "filswan/ubi-worker-cpu-amd:latest"
const UBITaskImageAmdGpu = "filswan/ubi-worker-gpu-amd:latest"
const UBIResourceExporterDockerImage = "filswan/resource-exporter:v11.3.0"
const TraefikServerDockerImage = "traefik:v2.10"

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
	Network        string `json:"network"`
	BoundaryHeight uint64 `json:"boundary_height"`
	Config         struct {
		SequencerUrl   string `json:"sequencer_url"`
		ZkEnginePk     string `json:"zk_engine_pk"`
		OrchestratorPk string `json:"orchestrator_pk"`
		ChainRpc       string `json:"chain_rpc"`
		EdgeUrl        string `json:"edge_url"`
		LegacyContract struct {
			SwanTokenContract              string `json:"swan_token_contract"`
			OrchestratorCollateralContract string `json:"orchestrator_collateral_contract"`
			JobManagerContract             string `json:"job_manager_contract"`
			JobManagerContractCreated      uint64 `json:"job_manager_contract_created"`
			RegisterCpContract             string `json:"register_cp_contract"`
			ZkCollateralContract           string `json:"zk_collateral_contract"`
			RegisterTaskContract           string `json:"register_task_contract"`
			SequencerContract              string `json:"sequencer_contract"`
			EdgeTaskPayment                string `json:"edge_task_payment"`
			EdgeTaskPaymentCreated         int64  `json:"edge_task_payment_created"`
		} `json:"legacy_contract"`
		UpgradeContract struct {
			SwanTokenContract              string `json:"swan_token_contract"`
			OrchestratorCollateralContract string `json:"orchestrator_collateral_contract"`
			JobManagerContract             string `json:"job_manager_contract"`
			JobManagerContractCreated      uint64 `json:"job_manager_contract_created"`
			RegisterCpContract             string `json:"register_cp_contract"`
			ZkCollateralContract           string `json:"zk_collateral_contract"`
			RegisterTaskContract           string `json:"register_task_contract"`
			SequencerContract              string `json:"sequencer_contract"`
			EdgeTaskPayment                string `json:"edge_task_payment"`
			EdgeTaskPaymentCreated         int64  `json:"edge_task_payment_created"`
		} `json:"upgrade_contract"`
		OrchestratorCollateralUbiZeroContract string `json:"orchestrator_collateral_ubi_zero_contract"`
		ZkCollateralUbiZeroContract           string `json:"zk_collateral_ubi_zero_contract"`
	} `json:"config"`
}
