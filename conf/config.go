package conf

import (
	_ "embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var config *ComputeNode

const (
	MainnetNetwork = "mainnet"
	TestnetNetwork = "testnet"
)

// ComputeNode is a compute node config
type ComputeNode struct {
	API      API
	UBI      UBI
	LOG      LOG
	HUB      HUB
	MCS      MCS
	Registry Registry
	RPC      RPC
	CONTRACT CONTRACT
}

type API struct {
	Port            int
	MultiAddress    string
	Domain          string
	NodeName        string
	WalletWhiteList string
}
type UBI struct {
	UbiEnginePk      string
	AggregateCommits bool
	SequencerUrl     string
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
	ApiKey      string
	AccessToken string
	BucketName  string
	Network     string
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
	SwanToken    string `toml:"SWAN_CONTRACT"`
	Collateral   string `toml:"SWAN_COLLATERAL_CONTRACT"`
	Register     string `toml:"REGISTER_CP_CONTRACT"`
	TaskRegister string `toml:"REGISTER_TASK_CONTRACT"`
	ZkCollateral string `toml:"ZK_COLLATERAL_CONTRACT"`
	Sequencer    string `toml:"SEQUENCER_CONTRACT"`
}

func GetRpcByNetWorkName(netWorkName ...string) (string, error) {
	var netWork string
	if len(netWorkName) > 0 && netWorkName[0] != "" {
		netWork = netWorkName[0]
	} else {
		netWork, _ = os.LookupEnv("CP_NETWORK")
		if netWork == "" {
			netWork = MainnetNetwork
		}
	}

	if netWork != MainnetNetwork && netWork != TestnetNetwork {
		return "", fmt.Errorf("not support network: %s", netWorkName[0])
	}
	return GetConfig().RPC.SwanChainRpc, nil
}

func InitConfig(cpRepoPath string, standalone bool) error {
	netWork, ok := os.LookupEnv("CP_NETWORK")
	if !ok {
		netWork = MainnetNetwork
	}
	if netWork != MainnetNetwork && netWork != TestnetNetwork {
		return fmt.Errorf("not support network: %s", netWork)
	}

	configFile := filepath.Join(cpRepoPath, fmt.Sprintf("config-%s.toml", netWork))

	if _, err := os.Stat(configFile); err != nil {
		return fmt.Errorf("not found %s config file, please use `computing-provider init` to initialize a repo", configFile)
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
		{"CONTRACT"},

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

		{"CONTRACT", "SWAN_CONTRACT"},
		{"CONTRACT", "SWAN_COLLATERAL_CONTRACT"},
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
		{"HUB"},

		{"API", "MultiAddress"},
		{"API", "NodeName"},

		{"UBI", "UbiEnginePk"},

		{"RPC", "SWAN_CHAIN_RPC"},

		{"CONTRACT", "SWAN_CONTRACT"},
		{"CONTRACT", "SWAN_COLLATERAL_CONTRACT"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}

//go:embed config-testnet.toml.sample
var testnetConfigContent string

//go:embed config-mainnet.toml.sample
var mainnetConfigContent string

func GenerateAndUpdateConfigFile(cpRepoPath string, multiAddress, nodeName string, port int) error {
	var configTmpl ComputeNode
	var configContent string
	var configFile *os.File
	var err error

	netWork, ok := os.LookupEnv("CP_NETWORK")
	if !ok {
		netWork = MainnetNetwork
	}
	if netWork != MainnetNetwork && netWork != TestnetNetwork {
		return fmt.Errorf("not support network: %s", netWork)
	}

	switch netWork {
	case MainnetNetwork:
		configContent = mainnetConfigContent
	case TestnetNetwork:
		configContent = testnetConfigContent
	default:
		return fmt.Errorf("not support network: %s", netWork)
	}

	configName := fmt.Sprintf("config-%s.toml", netWork)
	configFilePath := path.Join(cpRepoPath, configName)
	if _, err = os.Stat(configFilePath); os.IsNotExist(err) {
		logs.GetLogger().Warnf("The configuration file %s not found, generating this configuration file", configFilePath)
		if _, err = toml.Decode(configContent, &configTmpl); err != nil {
			return fmt.Errorf("parse toml data failed, error: %v", err)
		}
		configFile, err = os.Create(configFilePath)
		if err != nil {
			return fmt.Errorf("create %s file failed, error: %v", configName, err)
		}
		if err = toml.NewEncoder(configFile).Encode(configTmpl); err != nil {
			return fmt.Errorf("write data to %s file failed, error: %v", configName, err)
		}
	}

	if _, err = toml.DecodeFile(configFilePath, &configTmpl); err != nil {
		return err
	}
	os.Remove(configFilePath)

	configFile, err = os.Create(configFilePath)
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
	return nil
}
