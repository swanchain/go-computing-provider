package main

import (
	"context"
	"fmt"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/internal/contract/ecp"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/wallet"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var ubiZeroCmd = &cli.Command{
	Name:  "ubi-0",
	Usage: "Manage ubi-0 account collateral info",
	Subcommands: []*cli.Command{
		collateralInfoCmd,
		withdrawFromCollateralCmd,
	},
}

var collateralInfoCmd = &cli.Command{
	Name:  "info",
	Usage: "Print computing-provider info",
	Action: func(cctx *cli.Context) error {
		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		localNodeId := computing.GetNodeId(cpRepoPath)
		chainRpc, err := conf.GetRpcByNetWorkName()
		if err != nil {
			return err
		}
		client, err := contract.GetEthClient(chainRpc)
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
		var contractAddress, ownerAddress, workerAddress, beneficiaryAddress, chainNodeId, version string
		var cpAccount models.Account

		cpStub, err := account.NewAccountStub(client)
		if err == nil {
			cpAccount, err = cpStub.GetCpAccountInfo()
			if err != nil {
				err = fmt.Errorf("get cpAccount info on the chain failed, error: %v", err)
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
		fcpCollateralStub, err := fcp.NewCollateralWithUbiZeroStub(client)
		if err == nil {
			fcpCollateralInfo, err := fcpCollateralStub.CollateralInfo()
			if err == nil {
				fcpCollateralBalance = fcpCollateralInfo.AvailableBalance
				fcpEscrowBalance = fcpCollateralInfo.LockedCollateral
			}
		}

		ecpCollateral, err := ecp.NewCollateralWithUbiZeroStub(client, ecp.WithPublicKey(ownerAddress))
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
		taskData = append(taskData, []string{"   Owner Balance(ETH):", ownerBalance})
		taskData = append(taskData, []string{"   Worker Balance(ETH):", workerBalance})
		taskData = append(taskData, []string{""})
		taskData = append(taskData, []string{"ECP Balance(SWAN):"})
		taskData = append(taskData, []string{"   Collateral:", ecpCollateralBalance})
		taskData = append(taskData, []string{"   Escrow:", ecpEscrowBalance})
		taskData = append(taskData, []string{"FCP Balance(SWAN):"})
		taskData = append(taskData, []string{"   Collateral:", fcpCollateralBalance})
		taskData = append(taskData, []string{"   Escrow:", fcpEscrowBalance})

		header := []string{"CP Account Info:"}
		NewVisualTable(header, taskData, []RowColor{}).SetAutoWrapText(false).Generate(false)
		if err != nil {
			return err
		}

		if contractAddress == "" {
			fmt.Printf("Error: CP Account does not exist, please run 'computing-provider account create'.\n")
			return nil
		}

		if localNodeId != chainNodeId {
			fmt.Printf("NodeId mismatch, local node id: %s, chain node id: %s.\n", localNodeId, chainNodeId)
		}
		return nil
	},
}

var withdrawFromCollateralCmd = &cli.Command{
	Name:  "withdraw",
	Usage: "Withdraw funds from the collateral contract",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "fcp",
			Usage: "Specify the fcp collateral",
		},
		&cli.BoolFlag{
			Name:  "ecp",
			Usage: "Specify the ecp collateral",
		},
		&cli.StringFlag{
			Name:  "owner",
			Usage: "Specify the owner address",
		},
		&cli.StringFlag{
			Name:  "account",
			Usage: "Specify the cp account address, if not specified, cp account is the content of the account file under the CP_PATH variable",
		},
	},
	ArgsUsage: "[amount]",
	Action: func(cctx *cli.Context) error {
		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		ctx := reqContext(cctx)
		fcpWithdraw := cctx.Bool("fcp")
		ecpWithDraw := cctx.Bool("ecp")

		if !fcpWithdraw && !ecpWithDraw {
			return fmt.Errorf("must specify one of fcp or ecp")
		}
		var withdrawType string
		if fcpWithdraw {
			withdrawType = "fcp"
		}
		if ecpWithDraw {
			withdrawType = "ecp"
		}
		if withdrawType == "" {
			return fmt.Errorf("not suport withdraw type")
		}

		ownerAddress := cctx.String("owner")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("the owner address is required")
		}

		cpAccountAddress := cctx.String("account")
		amount := cctx.Args().Get(0)
		if strings.TrimSpace(amount) == "" {
			return fmt.Errorf("the amount param is required")
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		txHash, err := localWallet.CollateralWithdrawWithUbiZero(ctx, ownerAddress, amount, cpAccountAddress, withdrawType)
		if err != nil {
			return err
		}
		fmt.Println(txHash)
		return nil
	},
}
