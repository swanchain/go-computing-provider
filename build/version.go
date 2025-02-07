package build

import (
	_ "embed"
	"encoding/json"
	"log"
)

var CurrentCommit string

var NetWorkTag string

const BuildVersion = "1.0.2"

const UBITaskImageIntelCpu = "filswan/ubi-worker-cpu-intel:latest"
const UBITaskImageIntelGpu = "filswan/ubi-worker-gpu-intel:latest"
const UBITaskImageAmdCpu = "filswan/ubi-worker-cpu-amd:latest"
const UBITaskImageAmdGpu = "filswan/ubi-worker-gpu-amd:latest"
const UBIResourceExporterDockerImage = "filswan/resource-exporter:v12.0.0-beta"
const TraefikServerDockerImage = "traefik:v2.10"

const ResourceExporterVersion = "v12.0.0"

func UserVersion() string {
	return BuildVersion + "+" + NetWorkTag + CurrentCommit
}

var UpgradeVersionForMainnet = []UpgradeInfo{
	{
		NetworkName: "Mainnet_v1.0.0",
		Height:      3144197, // 2024-12-16 23:00:00 (UTC+8)
	},
}

var UpgradeVersionForTestnet = []UpgradeInfo{
	{
		NetworkName: "Testnet_v1.0.0",
		Height:      10238200,
	},
}

//go:embed parameters.json
var netWorkConfig string

func LoadParam() []NetworkConfig {
	var config []NetworkConfig
	err := json.Unmarshal([]byte(netWorkConfig), &config)
	if err != nil {
		log.Fatalln(err)
	}

	var mainnetLast, testnetLast UpgradeInfo
	if len(UpgradeVersionForMainnet) > 0 {
		mainnetLast = UpgradeVersionForMainnet[len(UpgradeVersionForMainnet)-1]
	}
	if len(UpgradeVersionForTestnet) > 0 {
		testnetLast = UpgradeVersionForTestnet[len(UpgradeVersionForTestnet)-1]
	}

	for i, networkConfig := range config {
		if networkConfig.Network == NetWorkTag && NetWorkTag == "mainnet" {
			config[i].BoundaryHeight = mainnetLast.Height
			config[i].UpgradeName = mainnetLast.NetworkName
		}
		if networkConfig.Network == NetWorkTag && NetWorkTag == "testnet" {
			config[i].BoundaryHeight = testnetLast.Height
			config[i].UpgradeName = testnetLast.NetworkName
		}
	}
	return config
}

type NetworkConfig struct {
	Network        string `json:"network"`
	BoundaryHeight uint64 `json:"boundary_height"`
	UpgradeName    string `json:"upgrade_name"`
	Config         struct {
		SequencerUrl   string `json:"sequencer_url"`
		ZkEnginePk     string `json:"zk_engine_pk"`
		OrchestratorPk string `json:"orchestrator_pk"`
		ChainRpc       string `json:"chain_rpc"`
		EdgeUrl        string `json:"edge_url"`
		LegacyContract struct {
			NetworkName                    string `json:"network_name"`
			SwanTokenContract              string `json:"swan_token_contract"`
			OrchestratorCollateralContract string `json:"orchestrator_collateral_contract"`
			JobManagerContract             string `json:"job_manager_contract"`
			JobManagerContractCreated      uint64 `json:"job_manager_contract_created"`
			RegisterCpContract             string `json:"register_cp_contract"`
			ZkCollateralContract           string `json:"zk_collateral_contract"`
			RegisterTaskContract           string `json:"register_task_contract"`
			SequencerContract              string `json:"sequencer_contract"`
			EdgeTaskPayment                string `json:"edge_task_payment"`
			EdgeTaskPaymentCreated         uint64 `json:"edge_task_payment_created"`
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
			EdgeTaskPaymentCreated         uint64 `json:"edge_task_payment_created"`
		} `json:"upgrade_contract"`
		OrchestratorCollateralUbiZeroContract string `json:"orchestrator_collateral_ubi_zero_contract"`
		ZkCollateralUbiZeroContract           string `json:"zk_collateral_ubi_zero_contract"`
	} `json:"config"`
}

type UpgradeInfo struct {
	Height      uint64
	NetworkName string
}
