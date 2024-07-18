package conf

import (
	_ "embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/swanchain/go-computing-provider/build"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var config *ComputeNode

// ComputeNode is a compute node config
type ComputeNode struct {
	API      API
	UBI      UBI
	LOG      LOG
	HUB      HUB
	MCS      MCS
	Registry Registry
	RPC      RPC
	CONTRACT CONTRACT `toml:"CONTRACT,omitempty"`
}

type API struct {
	Port            int
	MultiAddress    string
	Domain          string
	NodeName        string
	WalletWhiteList string
	WalletBlackList string
}
type UBI struct {
	UbiEnginePk string
}

type LOG struct {
	CrtFile string
	KeyFile string
}

type HUB struct {
	ServerUrl        string
	AccessToken      string
	BalanceThreshold float64
	OrchestratorPk   string
	VerifySign       bool
}

type MCS struct {
	ApiKey     string
	BucketName string
	Network    string
}

type Registry struct {
	ServerAddress string
	UserName      string
	Password      string
}

type RPC struct {
	SwanChainRpc string `toml:"SWAN_CHAIN_RPC"`
}

type CONTRACT struct {
	SwanToken         string `toml:"SWAN_CONTRACT"`
	Collateral        string `toml:"SWAN_COLLATERAL_CONTRACT"`
	CpAccountRegister string `toml:"REGISTER_CP_CONTRACT"`
	TaskRegister      string `toml:"REGISTER_TASK_CONTRACT"`
	ZkCollateral      string `toml:"ZK_COLLATERAL_CONTRACT"`
}

func GetRpcByNetWorkName() (string, error) {
	if len(strings.TrimSpace(GetConfig().RPC.SwanChainRpc)) == 0 {
		return "", fmt.Errorf("You need to set SWAN_CHAIN_RPC in the configuration file")
	}
	return GetConfig().RPC.SwanChainRpc, nil
}

func InitConfig(cpRepoPath string, standalone bool) error {
	configFile := filepath.Join(cpRepoPath, "config.toml")

	if _, err := os.Stat(configFile); err != nil {
		return fmt.Errorf("not found %s repo, "+
			"please use `computing-provider init` to initialize the repo ", cpRepoPath)
	}

	metaData, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		return fmt.Errorf("failed load config file, path: %s, error: %w", configFile, err)
	}

	if standalone {
		if !requiredFieldsAreGivenForSeparate(metaData) {
			log.Fatal("Required fields not given")
		}
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("Required fields not given")
		}
	}

	networkConfig := build.LoadParam()
	for _, nc := range networkConfig {
		ncCopy := nc
		if ncCopy.Network == build.NetWorkTag {
			config.CONTRACT.SwanToken = ncCopy.Config.SwanTokenContract
			config.CONTRACT.Collateral = ncCopy.Config.OrchestratorCollateralContract
			config.CONTRACT.Register = ncCopy.Config.RegisterCpContract
			config.CONTRACT.ZkCollateral = ncCopy.Config.ZkCollateralContract
		}
	}
	return nil
}

func GetConfig() *ComputeNode {
	return config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"API"},
		{"LOG"},
		{"UBI"},
		{"HUB"},
		{"MCS"},
		{"Registry"},
		{"RPC"},

		{"API", "MultiAddress"},
		{"API", "Domain"},
		{"API", "NodeName"},

		{"LOG", "CrtFile"},
		{"LOG", "KeyFile"},

		{"UBI", "UbiEnginePk"},

		{"HUB", "ServerUrl"},
		{"HUB", "AccessToken"},

		{"MCS", "ApiKey"},
		{"MCS", "BucketName"},
		{"MCS", "Network"},

		{"RPC", "SWAN_CHAIN_RPC"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}

func requiredFieldsAreGivenForSeparate(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"API"},
		{"UBI"},
		{"RPC"},

		{"API", "MultiAddress"},
		{"API", "NodeName"},

		{"UBI", "UbiEnginePk"},

		{"RPC", "SWAN_CHAIN_RPC"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}

func GenerateAndUpdateConfigFile(cpRepoPath string, multiAddress, nodeName string, port int) error {
	fmt.Println("Checking if repo exists")

	if Exists(cpRepoPath) {
		return fmt.Errorf("repo at '%s' is already initialized", cpRepoPath)
	}

	var configTmpl ComputeNode

	configFilePath := path.Join(cpRepoPath, "config.toml")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		defaultComputeNode := generateDefaultConfig()
		networkConfig := build.LoadParam()
		for _, nc := range networkConfig {
			ncCopy := nc
			if ncCopy.Network == build.NetWorkTag {
				defaultComputeNode.UBI.UbiEnginePk = ncCopy.Config.ZkEnginePk

				defaultComputeNode.HUB.ServerUrl = ncCopy.Config.OrchestratorUrl
				defaultComputeNode.HUB.OrchestratorPk = ncCopy.Config.OrchestratorPk

				defaultComputeNode.RPC.SwanChainRpc = ncCopy.Config.ChainRpc
			}
		}

		configFile, err := os.Create(configFilePath)
		if err != nil {
			return fmt.Errorf("create %s file failed, error: %v", configFilePath, err)
		}
		if err = toml.NewEncoder(configFile).Encode(defaultComputeNode); err != nil {
			return fmt.Errorf("write data to %s file failed, error: %v", configFilePath, err)
		}

		configTmpl = defaultComputeNode
	} else {
		if _, err = toml.DecodeFile(configFilePath, &configTmpl); err != nil {
			return err
		}
	}

	os.Remove(configFilePath)
	configFile, err := os.Create(configFilePath)
	if err != nil {
		return err
	}

	if len(multiAddress) != 0 && !strings.EqualFold(multiAddress, strings.TrimSpace(configTmpl.API.MultiAddress)) {
		configTmpl.API.MultiAddress = multiAddress
	}

	if len(strings.TrimSpace(nodeName)) != 0 {
		configTmpl.API.NodeName = nodeName
	} else {
		hostname, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("get hostname failed, error: %v", err)
		}
		configTmpl.API.NodeName = hostname
	}

	if port != 0 {
		configTmpl.API.Port = port
	}

	if err = toml.NewEncoder(configFile).Encode(configTmpl); err != nil {
		return err
	}

	file, err := os.Create(path.Join(cpRepoPath, "provider.db"))
	if err != nil {
		return err
	}
	defer file.Close()

	if err = os.MkdirAll(path.Join(cpRepoPath, "keystore"), 0755); err != nil {
		return fmt.Errorf("failed to create keystore, error: %v", err)
	}

	fmt.Printf("Initialized CP repo at '%s'. \n", cpRepoPath)
	return nil
}

func generateDefaultConfig() ComputeNode {
	return ComputeNode{
		API: API{
			Port:            8085,
			MultiAddress:    "/ip4/<PUBLIC_IP>/tcp/<PORT>",
			Domain:          "",
			NodeName:        "<YOUR_CP_Node_Name>",
			WalletWhiteList: "",
			WalletBlackList: "",
		},
		UBI: UBI{
			UbiEnginePk: "",
		},
		LOG: LOG{
			CrtFile: "",
			KeyFile: "",
		},
		HUB: HUB{
			ServerUrl:        "",
			AccessToken:      "",
			BalanceThreshold: 0.1,
			OrchestratorPk:   "",
			VerifySign:       true,
		},
		MCS: MCS{
			ApiKey:     "",
			BucketName: "",
			Network:    "polygon.mainnet",
		},
		Registry: Registry{
			ServerAddress: "",
			UserName:      "",
			Password:      "",
		},
		RPC: RPC{
			SwanChainRpc: "",
		},
		CONTRACT: CONTRACT{
			SwanToken:    "",
			Collateral:   "",
			Register:     "",
			ZkCollateral: "",
		},
	}
}

func Exists(cpPath string) bool {
	_, err := os.Stat(filepath.Join(cpPath, "keystore"))
	KeyStoreNoExist := os.IsNotExist(err)
	err = nil
	_, err = os.Stat(filepath.Join(cpPath, "provider.db"))
	providerNotExist := os.IsNotExist(err)

	if KeyStoreNoExist && providerNotExist {
		return false
	}
	return true
}
