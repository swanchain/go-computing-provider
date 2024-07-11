package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	account2 "github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/internal/contract/ecp"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/contract/token"
	"github.com/swanchain/go-computing-provider/internal/initializer"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/swanchain/go-computing-provider/wallet"
	"github.com/urfave/cli/v2"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Start a cp process",
	Action: func(cctx *cli.Context) error {
		logs.GetLogger().Info("Start a computing provider client.")

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		initializer.ProjectInit(cpRepoPath)

		r := gin.Default()
		r.Use(cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			ValidateHeaders: false,
		}))
		pprof.Register(r)

		v1 := r.Group("/api/v1")
		cpManager(v1.Group("/computing"))

		shutdownChan := make(chan struct{})
		httpStopper, err := util.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port), true)
		if err != nil {
			logs.GetLogger().Fatal("failed to start cp-api endpoint: %s", err)
		}

		finishCh := util.MonitorShutdown(shutdownChan,
			util.ShutdownHandler{Component: "cp-api", StopFunc: httpStopper},
		)
		<-finishCh

		return nil
	},
}

func cpManager(router *gin.RouterGroup) {
	router.GET("/cp", computing.StatisticalSources)
	router.GET("/host/info", computing.GetServiceProviderInfo)
	router.POST("/lagrange/jobs", computing.ReceiveJob)
	router.POST("/lagrange/jobs/redeploy", computing.RedeployJob)
	router.DELETE("/lagrange/jobs", computing.CancelJob)
	router.POST("/lagrange/jobs/renew", computing.ReNewJob)
	router.GET("/lagrange/spaces/log", computing.GetSpaceLog)
	router.POST("/lagrange/cp/proof", computing.DoProof)
	router.GET("/lagrange/cp/whitelist", computing.WhiteList)
	router.GET("/lagrange/cp/blacklist", computing.BlackList)
	router.GET("/lagrange/job/:job_uuid", computing.GetJobStatus)

	router.POST("/cp/ubi", computing.DoUbiTaskForK8s)
	router.POST("/cp/receive/ubi", computing.ReceiveUbiProofForK8s)

}

var infoCmd = &cli.Command{
	Name:  "info",
	Usage: "Print computing-provider info",
	Action: func(cctx *cli.Context) error {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		localNodeId := computing.GetNodeId(cpRepoPath)
		k8sService := computing.NewK8sService()
		var count int
		if k8sService.Version == "" {
			count = 0
		} else {
			count, _ = k8sService.GetDeploymentActiveCount()
		}

		chainRpc, err := conf.GetRpcByNetWorkName()
		if err != nil {
			return err
		}
		client, err := ethclient.Dial(chainRpc)
		if err != nil {
			return err
		}
		defer client.Close()

		var netWork = ""
		chainId, err := client.ChainID(context.Background())
		if err != nil {
			return err
		}
		if chainId.Int64() == 254 {
			netWork = fmt.Sprintf("Mainnet(%d)", chainId.Int64())
		} else {
			netWork = fmt.Sprintf("Testnet(%d)", chainId.Int64())
		}

		var fcpCollateralBalance = "0.0000"
		var fcpEscrowBalance = "0.0000"
		var ecpCollateralBalance = "0.0000"
		var ecpEscrowBalance = "0.0000"
		var ownerBalance = "0.0000"
		var workerBalance = "0.0000"
		var contractAddress, ownerAddress, workerAddress, beneficiaryAddress, taskTypes, chainNodeId, version string
		var cpAccount models.Account

		cpStub, err := account2.NewAccountStub(client)
		if err == nil {
			cpAccount, err = cpStub.GetCpAccountInfo()
			if err != nil {
				err = fmt.Errorf("get cpAccount info on the chain failed, error: %v", err)
			}

			for _, taskType := range cpAccount.TaskTypes {
				taskTypes += models.TaskTypeStr(int(taskType)) + ", "
			}
			if taskTypes != "" {
				taskTypes = taskTypes[:len(taskTypes)-2]
			}

			contractAddress = cpStub.ContractAddress
			ownerAddress = cpAccount.OwnerAddress
			workerAddress = cpAccount.WorkerAddress
			beneficiaryAddress = cpAccount.Beneficiary
			chainNodeId = cpAccount.NodeId
			version = cpAccount.Version
		}

		ownerBalance, err = wallet.Balance(context.TODO(), client, ownerAddress)
		workerBalance, err = wallet.Balance(context.TODO(), client, workerAddress)
		fcpCollateralStub, err := fcp.NewCollateralStub(client)
		if err == nil {
			fcpCollateralInfo, err := fcpCollateralStub.CollateralInfo()
			if err == nil {
				fcpCollateralBalance = fcpCollateralInfo.AvailableBalance
				fcpEscrowBalance = fcpCollateralInfo.LockedCollateral
			}
		}

		ecpCollateral, err := ecp.NewCollateralStub(client, ecp.WithPublicKey(ownerAddress))
		if err == nil {
			cpCollateralInfo, err := ecpCollateral.CpInfo()
			if err == nil {
				ecpCollateralBalance = cpCollateralInfo.CollateralBalance
				ecpEscrowBalance = cpCollateralInfo.FrozenBalance
			}
		}

		var domain = conf.GetConfig().API.Domain
		if strings.HasPrefix(domain, ".") {
			domain = domain[1:]
		}
		var taskData [][]string

		taskData = append(taskData, []string{"   Network:", netWork})
		taskData = append(taskData, []string{fmt.Sprintf("   CP Account Address(%s):", version), contractAddress})
		taskData = append(taskData, []string{"   Name:", conf.GetConfig().API.NodeName})
		taskData = append(taskData, []string{"   Owner:", ownerAddress})
		taskData = append(taskData, []string{"   Node ID:", localNodeId})
		taskData = append(taskData, []string{"   Domain:", domain})
		taskData = append(taskData, []string{"   Multi-Address:", conf.GetConfig().API.MultiAddress})
		taskData = append(taskData, []string{"   Worker Address:", workerAddress})
		taskData = append(taskData, []string{"   Beneficiary Address:", beneficiaryAddress})
		taskData = append(taskData, []string{""})
		taskData = append(taskData, []string{"Capabilities:"})
		taskData = append(taskData, []string{"   Task Types:", taskTypes})
		taskData = append(taskData, []string{"   Applications:", strconv.Itoa(count)})
		taskData = append(taskData, []string{""})
		taskData = append(taskData, []string{"Owner Balance(sETH):", ownerBalance})
		taskData = append(taskData, []string{"Worker Balance(sETH):", workerBalance})
		taskData = append(taskData, []string{""})
		taskData = append(taskData, []string{"ECP Balance(SWANC):"})
		taskData = append(taskData, []string{"   Collateral:", ecpCollateralBalance})
		taskData = append(taskData, []string{"   Escrow:", ecpEscrowBalance})
		taskData = append(taskData, []string{"FCP Balance(SWANC):"})
		taskData = append(taskData, []string{"   Collateral:", fcpCollateralBalance})
		taskData = append(taskData, []string{"   Escrow:", fcpEscrowBalance})

		var rowColorList []RowColor
		if taskTypes != "" {
			rowColorList = append(rowColorList,
				RowColor{
					row:    0,
					column: []int{1},
					color:  []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgBlueColor}},
				},
				RowColor{
					row:    11,
					column: []int{1},
					color:  []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}},
				})
		}
		header := []string{"CP Account Info:"}
		NewVisualTable(header, taskData, rowColorList).Generate(false)
		if err != nil {
			return err
		}
		if localNodeId != chainNodeId {
			fmt.Printf("NodeId mismatch, local node id: %s, chain node id: %s.\n", localNodeId, chainNodeId)
		}
		return nil
	},
}

var stateCmd = &cli.Command{
	Name:  "state",
	Usage: "Print computing-provider info on the chain",
	Subcommands: []*cli.Command{
		stateInfoCmd,
		taskInfoCmd,
	},
	Before: func(c *cli.Context) error {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}
		return nil
	},
}

var stateInfoCmd = &cli.Command{
	Name:      "cp-info",
	Usage:     "Print computing-provider chain info",
	ArgsUsage: "[cp_account_contract_address]",
	Action: func(cctx *cli.Context) error {
		chainRpc, err := conf.GetRpcByNetWorkName()
		if err != nil {
			return err
		}
		client, err := ethclient.Dial(chainRpc)
		if err != nil {
			return err
		}
		defer client.Close()

		var fcpCollateralBalance = "0.0000"
		var fcpEscrowBalance = "0.0000"
		var ecpCollateralBalance = "0.0000"
		var ecpEscrowBalance = "0.0000"
		var ownerBalance = "0.0000"
		var workerBalance = "0.0000"
		var chainMultiAddress string
		var contractAddress, ownerAddress, workerAddress, beneficiaryAddress, taskTypes, chainNodeId, version string

		cpStub, err := account2.NewAccountStub(client, account2.WithContractAddress(cctx.Args().Get(0)))
		if err == nil {
			cpAccount, err := cpStub.GetCpAccountInfo()
			if err != nil {
				err = fmt.Errorf("get cpAccount failed, error: %v", err)
			}

			for _, taskType := range cpAccount.TaskTypes {
				taskTypes += models.TaskTypeStr(int(taskType)) + ","
			}
			if taskTypes != "" {
				taskTypes = taskTypes[:len(taskTypes)-1]
			}

			contractAddress = cpStub.ContractAddress
			ownerAddress = cpAccount.OwnerAddress
			workerAddress = cpAccount.WorkerAddress
			beneficiaryAddress = cpAccount.Beneficiary
			chainNodeId = cpAccount.NodeId
			chainMultiAddress = strings.Join(cpAccount.MultiAddresses, ",")
			version = cpAccount.Version
		}

		if strings.HasSuffix(chainMultiAddress, ",") {
			chainMultiAddress = chainMultiAddress[:len(chainMultiAddress)-1]
		}

		ownerBalance, err = wallet.Balance(context.TODO(), client, ownerAddress)
		workerBalance, err = wallet.Balance(context.TODO(), client, workerAddress)
		fcpCollateralStub, err := fcp.NewCollateralStub(client, fcp.WithCpAccountAddress(contractAddress))
		if err == nil {
			fcpCollateralInfo, err := fcpCollateralStub.CollateralInfo()
			if err == nil {
				fcpCollateralBalance = fcpCollateralInfo.AvailableBalance
				fcpEscrowBalance = fcpCollateralInfo.LockedCollateral
			}
		}

		ecpCollateral, err := ecp.NewCollateralStub(client, ecp.WithCpAccountAddress(contractAddress))
		if err == nil {
			cpCollateralInfo, err := ecpCollateral.CpInfo()
			if err == nil {
				ecpCollateralBalance = cpCollateralInfo.CollateralBalance
				ecpEscrowBalance = cpCollateralInfo.FrozenBalance
			}
		}

		var taskData [][]string
		taskData = append(taskData, []string{"Node ID:", chainNodeId})
		taskData = append(taskData, []string{"Multi-Address:", chainMultiAddress})
		taskData = append(taskData, []string{"Owner:", ownerAddress})
		taskData = append(taskData, []string{"Worker Address:", workerAddress})
		taskData = append(taskData, []string{"Beneficiary Address:", beneficiaryAddress})
		taskData = append(taskData, []string{"Task Types:", taskTypes})
		taskData = append(taskData, []string{""})
		taskData = append(taskData, []string{"Owner Balance(sETH):", ownerBalance})
		taskData = append(taskData, []string{"Worker Balance(sETH):", workerBalance})
		taskData = append(taskData, []string{""})
		taskData = append(taskData, []string{"ECP Balance(SWANC):"})
		taskData = append(taskData, []string{"   Collateral:", ecpCollateralBalance})
		taskData = append(taskData, []string{"   Escrow:", ecpEscrowBalance})
		taskData = append(taskData, []string{"FCP Balance(SWANC):"})
		taskData = append(taskData, []string{"   Collateral:", fcpCollateralBalance})
		taskData = append(taskData, []string{"   Escrow:", fcpEscrowBalance})

		var rowColorList []RowColor
		if taskTypes != "" {
			var rowColor []tablewriter.Colors
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			rowColorList = append(rowColorList, RowColor{
				row:    5,
				column: []int{1},
				color:  rowColor,
			})
		}
		header := []string{fmt.Sprintf("CP Account Address(%s):", version), contractAddress}
		NewVisualTable(header, taskData, rowColorList).Generate(false)
		return nil
	},
}

var taskInfoCmd = &cli.Command{
	Name:      "task-info",
	Usage:     "Print task info on the chain",
	ArgsUsage: "[task_contract_address]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:     "ecp",
			Usage:    "Check ECP task on the chain",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		taskContract := cctx.Args().Get(0)
		if strings.TrimSpace(taskContract) == "" {
			return fmt.Errorf("the task contract address is required")
		}

		chainRpc, err := conf.GetRpcByNetWorkName()
		if err != nil {
			return err
		}

		taskInfo, err := computing.GetTaskInfoOnChain(taskContract)
		if err != nil {
			return fmt.Errorf("get task info on the chain failed, error: %v", err)
		}

		var lockFundTx, unlockFundTx, rewardTx, challengeTx, slashTx, reward string
		if taskInfo.LockFundTx != "" {
			lockFundTx = taskInfo.LockFundTx
		} else {
			lockFundTx = "-"
		}
		if taskInfo.UnlockFundTx != "" {
			unlockFundTx = taskInfo.UnlockFundTx
		} else {
			unlockFundTx = "-"
		}
		if taskInfo.RewardTx != "" {
			rewardTx = taskInfo.RewardTx
		} else {
			rewardTx = "-"
		}
		if taskInfo.ChallengeTx != "" {
			challengeTx = taskInfo.ChallengeTx
		} else {
			challengeTx = "-"
		}
		if taskInfo.SlashTx != "" {
			slashTx = taskInfo.SlashTx
		} else {
			slashTx = "-"
		}

		if taskInfo.RewardTx != "" {
			client, err := ethclient.Dial(chainRpc)
			if err == nil {
				defer client.Close()
				receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(taskInfo.RewardTx))
				if err == nil {
					contractAbi, err := abi.JSON(strings.NewReader(token.TokenMetaData.ABI))
					if err == nil {
						for _, l := range receipt.Logs {
							event := struct {
								From  common.Address
								To    common.Address
								Value *big.Int
							}{}
							if err := contractAbi.UnpackIntoInterface(&event, "Transfer", l.Data); err != nil {
								continue
							}

							if len(l.Topics) == 3 && l.Topics[0] == contractAbi.Events["Transfer"].ID {
								balance := event.Value
								if balance.String() == "0" {
									reward = "0.0000"
								} else {
									fbalance := new(big.Float)
									fbalance.SetString(balance.String())
									etherQuotient := new(big.Float).Quo(fbalance, new(big.Float).SetInt(big.NewInt(1e18)))
									reward = etherQuotient.Text('f', 4)
								}
							}
						}
					}
				}
			}
		}

		var taskData [][]string
		taskData = append(taskData, []string{"ZK Type:", models.TaskTypeStr(int(taskInfo.TaskType.Int64()))})
		taskData = append(taskData, []string{"Resource Type:", models.GetSourceTypeStr(int(taskInfo.ResourceType.Int64()))})
		taskData = append(taskData, []string{"CP Account:", taskInfo.CpContractAddress.Hex()})
		taskData = append(taskData, []string{"Task Status:", taskInfo.Status})
		taskData = append(taskData, []string{"Deadline:", taskInfo.Deadline.String()})
		taskData = append(taskData, []string{"Reward(SWAN):", reward})
		taskData = append(taskData, []string{"LockFund TxHash:", lockFundTx})
		taskData = append(taskData, []string{"UnLockFund TxHash:", unlockFundTx})
		taskData = append(taskData, []string{"Reward TxHash:", rewardTx})
		taskData = append(taskData, []string{"Challenge TxHash:", challengeTx})
		taskData = append(taskData, []string{"Slash TxHash:", slashTx})

		header := []string{"Task Contract:", taskContract}
		NewVisualTable(header, taskData, []RowColor{}).Generate(false)
		return nil

	},
}

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new cp",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "multi-address",
			Usage: "The multiAddress for libp2p(example: /ip4/<PUBLIC_IP>/tcp/<PORT>)",
		},
		&cli.StringFlag{
			Name:  "node-name",
			Usage: "The name of cp",
		},
		&cli.IntFlag{
			Name:  "port",
			Usage: "The cp listens on port",
			Value: 9085,
		},
	},
	Action: func(cctx *cli.Context) error {
		multiAddr := cctx.String("multi-address")
		port := cctx.Int("port")
		if strings.TrimSpace(multiAddr) == "" {
			return fmt.Errorf("the multi-address field required")
		}
		nodeName := cctx.String("node-name")

		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		return conf.GenerateAndUpdateConfigFile(cpRepoPath, strings.TrimSpace(multiAddr), nodeName, port)
	},
}

var accountCmd = &cli.Command{
	Name:  "account",
	Usage: "Manage account info of CP",
	Subcommands: []*cli.Command{
		createAccountCmd,
		changeMultiAddressCmd,
		changeOwnerAddressCmd,
		changeWorkerAddressCmd,
		changeBeneficiaryAddressCmd,
		changeTaskTypesCmd,
	},
	Before: func(c *cli.Context) error {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}
		return nil
	},
}

var createAccountCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a cp account to chain",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ownerAddress",
			Usage: "Specify a OwnerAddress",
		},
		&cli.StringFlag{
			Name:  "workerAddress",
			Usage: "Specify a workerAddress",
		},
		&cli.StringFlag{
			Name:  "beneficiaryAddress",
			Usage: "Specify a beneficiaryAddress to receive rewards. If not specified, use the same address as ownerAddress",
		},
		&cli.StringFlag{
			Name:  "task-types",
			Usage: "Task types of CP (1:Fil-C2-512M, 2:Aleo, 3:AI, 4:Fil-C2-32G), separated by commas",
		},
	},
	Action: func(cctx *cli.Context) error {
		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		workerAddress := cctx.String("workerAddress")
		if strings.TrimSpace(workerAddress) == "" {
			return fmt.Errorf("workerAddress is not empty")
		}

		beneficiaryAddress := cctx.String("beneficiaryAddress")
		if strings.TrimSpace(beneficiaryAddress) == "" {
			beneficiaryAddress = ownerAddress
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the ownerAddress is invalid wallet address")
		}

		if err := checkWalletAddress(ownerAddress, "owner"); err != nil {
			return err
		}

		if !isValidWalletAddress(workerAddress) {
			return fmt.Errorf("the workerAddress is invalid wallet address")
		}

		if err := checkWalletAddress(ownerAddress, "worker"); err != nil {
			return err
		}

		if !isValidWalletAddress(beneficiaryAddress) {
			return fmt.Errorf("the beneficiaryAddress is invalid wallet address")
		}

		if err := checkWalletAddress(ownerAddress, "beneficiary"); err != nil {
			return err
		}

		taskTypes := strings.TrimSpace(cctx.String("task-types"))
		if strings.TrimSpace(taskTypes) == "" {
			return fmt.Errorf("taskTypes is not empty")
		}

		var taskTypesUint []uint8

		taskTypes = handleStr(taskTypes)
		if strings.Index(taskTypes, ",") > 0 {
			for _, taskT := range strings.Split(taskTypes, ",") {
				tt, err := strconv.ParseUint(taskT, 10, 64)
				if err != nil {
					return err
				}
				if tt < 0 {
					return fmt.Errorf("task-types must be int")
				}
				taskTypesUint = append(taskTypesUint, uint8(tt))
			}
		} else {
			tt, err := strconv.ParseUint(taskTypes, 10, 64)
			if err != nil {
				return err
			}
			if tt < 0 {
				return fmt.Errorf("task-types must be int")
			}
			taskTypesUint = append(taskTypesUint, uint8(tt))
		}

		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		return createAccount(cpRepoPath, ownerAddress, beneficiaryAddress, workerAddress, taskTypesUint)
	},
}

var changeMultiAddressCmd = &cli.Command{
	Name:      "changeMultiAddress",
	Usage:     "Update MultiAddress of CP (/ip4/<public_ip>/tcp/<port>)",
	ArgsUsage: "[multiAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is required")
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the ownerAddress is invalid wallet address")
		}

		multiAddr := cctx.Args().Get(0)
		if strings.TrimSpace(multiAddr) == "" {
			return fmt.Errorf("multiAddress is required")
		}

		cpRepoPath, _ := os.LookupEnv("CP_PATH")

		client, cpStub, err := getVerifyAccountClient(ownerAddress)
		if err != nil {
			return fmt.Errorf("get cp account client failed, error: %v", err)
		}
		defer client.Close()

		newMultiAddress := []string{strings.TrimSpace(multiAddr)}
		changeMultiAddressTx, err := cpStub.ChangeMultiAddress(newMultiAddress)
		if err != nil {
			return fmt.Errorf("changeMultiAddress tx failed, error: %v", err)
		}

		nodeId := computing.GetNodeId(cpRepoPath)
		if err = computing.NewCpInfoService().UpdateCpInfoByNodeId(&models.CpInfoEntity{NodeId: nodeId, MultiAddresses: newMultiAddress}); err != nil {
			return fmt.Errorf("update multi_addresses of cp to db failed, error: %v", err)
		}
		fmt.Printf("changeMultiAddress Transaction hash: %s\n", changeMultiAddressTx)
		fmt.Printf("Multi-Address is changed successfully! please manually update the `MultiAddress` in the config file \n")

		return nil
	},
}

var changeOwnerAddressCmd = &cli.Command{
	Name:      "changeOwnerAddress",
	Usage:     "Update OwnerAddress of CP",
	ArgsUsage: "[the target newOwnerAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the ownerAddress is invalid wallet address")
		}

		newOwnerAddr := cctx.Args().Get(0)
		if strings.TrimSpace(newOwnerAddr) == "" {
			return fmt.Errorf("the target newOwnerAddress is required")
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the target newOwnerAddress is invalid wallet address")
		}

		if err := checkWalletAddress(newOwnerAddr, "owner"); err != nil {
			return err
		}

		cpRepoPath, _ := os.LookupEnv("CP_PATH")

		client, cpStub, err := getVerifyAccountClient(ownerAddress)
		if err != nil {
			return fmt.Errorf("get cp account client failed, error: %v", err)
		}
		defer client.Close()

		changeOwnerAddressTx, err := cpStub.ChangeOwnerAddress(common.HexToAddress(newOwnerAddr))
		if err != nil {
			logs.GetLogger().Errorf("changeOwnerAddress tx failed, error: %v", err)
			return err
		}

		nodeId := computing.GetNodeId(cpRepoPath)
		if err = computing.NewCpInfoService().UpdateCpInfoByNodeId(&models.CpInfoEntity{NodeId: nodeId, OwnerAddress: newOwnerAddr}); err != nil {
			return fmt.Errorf("update owner_address of cp to db failed, error: %v", err)
		}

		fmt.Printf("changeOwnerAddress Transaction hash: %s\n", changeOwnerAddressTx)
		return nil
	},
}

var changeBeneficiaryAddressCmd = &cli.Command{
	Name:      "changeBeneficiaryAddress",
	Usage:     "Update beneficiaryAddress of CP",
	ArgsUsage: "[beneficiaryAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the ownerAddress is invalid wallet address")
		}

		beneficiaryAddress := cctx.Args().Get(0)
		if strings.TrimSpace(beneficiaryAddress) == "" {
			return fmt.Errorf("failed to parse target beneficiary address: %s", beneficiaryAddress)
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the target beneficiary address is invalid wallet address")
		}

		if err := checkWalletAddress(beneficiaryAddress, "beneficiary"); err != nil {
			return err
		}

		cpRepoPath, _ := os.LookupEnv("CP_PATH")

		client, cpStub, err := getVerifyAccountClient(ownerAddress)
		if err != nil {
			return fmt.Errorf("get cp account client failed, error: %v", err)
		}
		defer client.Close()

		changeBeneficiaryAddressTx, err := cpStub.ChangeBeneficiary(common.HexToAddress(beneficiaryAddress))
		if err != nil {
			logs.GetLogger().Errorf("changeBeneficiaryAddress tx failed, error: %v", err)
			return err
		}

		nodeId := computing.GetNodeId(cpRepoPath)
		if err = computing.NewCpInfoService().UpdateCpInfoByNodeId(&models.CpInfoEntity{NodeId: nodeId, Beneficiary: beneficiaryAddress}); err != nil {
			return fmt.Errorf("update beneficiary_address of cp to db failed, error: %v", err)
		}

		fmt.Printf("changeBeneficiaryAddress Transaction hash: %s \n", changeBeneficiaryAddressTx)
		return nil
	},
}

var changeWorkerAddressCmd = &cli.Command{
	Name:      "changeWorkerAddress",
	Usage:     "Update workerAddress of CP",
	ArgsUsage: "[workerAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the ownerAddress is invalid wallet address")
		}

		workerAddress := cctx.Args().Get(0)
		if strings.TrimSpace(workerAddress) == "" {
			return fmt.Errorf("failed to parse target worker address: %s", workerAddress)
		}

		if !isValidWalletAddress(workerAddress) {
			return fmt.Errorf("the target worker address is invalid wallet address")
		}

		if err := checkWalletAddress(workerAddress, "worker"); err != nil {
			return err
		}

		cpRepoPath, _ := os.LookupEnv("CP_PATH")

		client, cpStub, err := getVerifyAccountClient(ownerAddress)
		if err != nil {
			return fmt.Errorf("get cp account client failed, error: %v", err)
		}
		defer client.Close()

		changeBeneficiaryAddressTx, err := cpStub.ChangeWorkerAddress(common.HexToAddress(workerAddress))
		if err != nil {
			logs.GetLogger().Errorf("changeWorkerAddress tx failed, error: %v", err)
			return err
		}

		nodeId := computing.GetNodeId(cpRepoPath)
		if err = computing.NewCpInfoService().UpdateCpInfoByNodeId(&models.CpInfoEntity{NodeId: nodeId, WorkerAddress: workerAddress}); err != nil {
			return fmt.Errorf("update worker_address of cp to db failed, error: %v", err)
		}

		fmt.Printf("changeWorkerAddress Transaction hash: %s \n", changeBeneficiaryAddressTx)
		return nil
	},
}

var changeTaskTypesCmd = &cli.Command{
	Name:      "changeTaskTypes",
	Usage:     "Update taskTypes of CP (1:Fil-C2-512M, 2:Aleo, 3: AI, 4:Fil-C2-32G), separated by commas",
	ArgsUsage: "[TaskTypes]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if !isValidWalletAddress(ownerAddress) {
			return fmt.Errorf("the ownerAddress is invalid wallet address")
		}

		taskTypes := strings.TrimSpace(cctx.Args().Get(0))
		if strings.TrimSpace(taskTypes) == "" {
			return fmt.Errorf("taskTypes is required")
		}

		var taskTypesUint []uint8
		taskTypes = handleStr(taskTypes)
		if strings.Index(taskTypes, ",") > 0 {
			for _, taskT := range strings.Split(taskTypes, ",") {
				tt, err := strconv.ParseUint(taskT, 10, 64)
				if err != nil {
					return err
				}
				if tt < 0 {
					return fmt.Errorf("task-types must be int")
				}
				taskTypesUint = append(taskTypesUint, uint8(tt))
			}
		} else {
			tt, err := strconv.ParseUint(taskTypes, 10, 64)
			if err != nil {
				return err
			}
			if tt < 0 {
				return fmt.Errorf("task-types must be int")
			}
			taskTypesUint = append(taskTypesUint, uint8(tt))
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			logs.GetLogger().Fatal(err)
		}

		client, cpStub, err := getVerifyAccountClient(ownerAddress)
		if err != nil {
			return fmt.Errorf("get cp account client failed, error: %v", err)
		}
		defer client.Close()

		changeTaskTypesTx, err := cpStub.ChangeTaskTypes(taskTypesUint)
		if err != nil {
			logs.GetLogger().Errorf("changeTaskTypes tx failed, error: %v", err)
			return err
		}

		nodeId := computing.GetNodeId(cpRepoPath)
		if err = computing.NewCpInfoService().UpdateCpInfoByNodeId(&models.CpInfoEntity{NodeId: nodeId, TaskTypes: taskTypesUint}); err != nil {
			return fmt.Errorf("update task_types of cp to db failed, error: %v", err)
		}

		fmt.Printf("changeTaskTypes Transaction hash: %s \n", changeTaskTypesTx)
		return nil
	},
}

var contractCmd = &cli.Command{
	Name:  "contract",
	Usage: "Manage contract info of CP",
	Subcommands: []*cli.Command{
		{
			Name:  "default",
			Usage: "Print the contract info that the current network CP is using",
			Action: func(c *cli.Context) error {
				cpRepoPath, _ := os.LookupEnv("CP_PATH")
				if err := conf.InitConfig(cpRepoPath, true); err != nil {
					return fmt.Errorf("load config file failed, error: %+v", err)
				}

				chainRpc, err := conf.GetRpcByNetWorkName()
				if err != nil {
					return err
				}
				client, err := ethclient.Dial(chainRpc)
				if err != nil {
					return err
				}
				defer client.Close()

				var netWork = ""
				chainId, err := client.ChainID(context.Background())
				if err != nil {
					return err
				}
				if chainId.Int64() == 254 {
					netWork = fmt.Sprintf("Mainnet(%d)", chainId.Int64())
				} else {
					netWork = fmt.Sprintf("Testnet(%d)", chainId.Int64())
				}

				contract := conf.GetConfig().CONTRACT
				var taskData [][]string

				taskData = append(taskData, []string{"Network:", netWork})
				taskData = append(taskData, []string{"Swan Token:", contract.SwanToken})
				taskData = append(taskData, []string{"Orchestrator Collateral:", contract.Collateral})
				taskData = append(taskData, []string{"Register CP:", contract.Register})
				taskData = append(taskData, []string{"ZK Collateral:", contract.ZkCollateral})

				var rowColorList []RowColor
				rowColorList = append(rowColorList,
					RowColor{
						row:    0,
						column: []int{1},
						color:  []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgBlueColor}},
					})

				header := []string{"CP Contract Info:"}
				NewVisualTable(header, taskData, rowColorList).Generate(false)
				return nil

			},
		},
	},
}

func isValidWalletAddress(address string) bool {
	re := regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	return re.MatchString(address)
}

func checkWalletAddress(walletAddress string, msg string) error {
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return err
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return err
	}
	defer client.Close()

	checkOwnerAddress := common.HexToAddress(walletAddress)
	bytecode, err := client.CodeAt(context.Background(), checkOwnerAddress, nil)
	if err != nil {
		return fmt.Errorf("check owner address failed, error: %v", err)
	}

	if len(bytecode) > 0 {
		return fmt.Errorf("the %s must be a wallet address", msg)
	}
	return nil
}

func handleStr(str string) string {
	if strings.HasSuffix(str, ",") {
		return strings.TrimSuffix(str, ",")
	} else {
		return str
	}
}
