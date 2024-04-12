package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/swanchain/go-computing-provider/account"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/initializer"
	"github.com/swanchain/go-computing-provider/internal/task/ubi"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/swanchain/go-computing-provider/wallet"
	"github.com/swanchain/go-computing-provider/wallet/contract/collateral"
	"github.com/urfave/cli/v2"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Start a cp process",
	Action: func(cctx *cli.Context) error {
		logs.GetLogger().Info("Start in computing provider mode.")

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
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
		ubi.RegisterUbiTaskHandles(ubi.NewFilC2Task(), ubi.NewAleoTask())

		v1 := r.Group("/api/v1")
		cpManager(v1.Group("/computing"))

		shutdownChan := make(chan struct{})
		httpStopper, err := util.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port))
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
	router.GET("/host/info", computing.GetServiceProviderInfo)
	router.POST("/lagrange/jobs", computing.ReceiveJob)
	router.POST("/lagrange/jobs/redeploy", computing.RedeployJob)
	router.DELETE("/lagrange/jobs", computing.CancelJob)
	router.GET("/lagrange/cp", computing.StatisticalSources)
	router.POST("/lagrange/jobs/renew", computing.ReNewJob)
	router.GET("/lagrange/spaces/log", computing.GetSpaceLog)
	router.POST("/lagrange/cp/proof", computing.DoProof)

	router.GET("/cp", computing.StatisticalSources)
	router.GET("/cp/info", computing.GetCpInfo)
	router.POST("/cp/ubi", ubi.DoUbiTask)
	router.POST("/cp/receive/ubi", ubi.ReceiveUbiProof)

}

var infoCmd = &cli.Command{
	Name:  "info",
	Usage: "Print computing-provider info",
	Action: func(cctx *cli.Context) error {

		cpPath, exit := os.LookupEnv("CP_PATH")
		if !exit {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=xxx")
		}
		if err := conf.InitConfig(cpPath); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		localNodeId := computing.GetNodeId(cpPath)

		k8sService := computing.NewK8sService()
		count, err := k8sService.GetDeploymentActiveCount()
		if err != nil {
			return err
		}

		chainRpc, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			return err
		}
		client, err := ethclient.Dial(chainRpc)
		if err != nil {
			return err
		}
		defer client.Close()

		var balance, collateralBalance string
		var contractAddress, ownerAddress, beneficiaryAddress, ubiFlag, chainNodeId string

		cpStub, err := account.NewAccountStub(client)
		if err == nil {
			cpAccount, err := cpStub.GetCpAccountInfo()
			if err != nil {
				return fmt.Errorf("get cpAccount faile, error: %v", err)
			}
			if cpAccount.UbiFlag == 1 {
				ubiFlag = "Accept"
			} else {
				ubiFlag = "Reject"
			}
			contractAddress = cpStub.ContractAddress
			ownerAddress = cpAccount.OwnerAddress
			beneficiaryAddress = cpAccount.Beneficiary.BeneficiaryAddress
			chainNodeId = cpAccount.NodeId
		}

		balance, err = wallet.Balance(context.TODO(), client, conf.GetConfig().HUB.WalletAddress)
		collateralStub, err := collateral.NewCollateralStub(client, collateral.WithPublicKey(conf.GetConfig().HUB.WalletAddress))
		if err == nil {
			collateralBalance, err = collateralStub.Balances()
		}

		var domain = conf.GetConfig().API.Domain
		if strings.HasPrefix(domain, ".") {
			domain = domain[1:]
		}
		var taskData [][]string

		taskData = append(taskData, []string{"Contract Address:", contractAddress})
		taskData = append(taskData, []string{"Multi-Address:", conf.GetConfig().API.MultiAddress})
		taskData = append(taskData, []string{"Name:", conf.GetConfig().API.NodeName})
		taskData = append(taskData, []string{"Node ID:", localNodeId})
		taskData = append(taskData, []string{"Domain:", domain})
		taskData = append(taskData, []string{"Running deployments:", strconv.Itoa(count)})

		taskData = append(taskData, []string{"Available(SWAN-ETH):", balance})
		taskData = append(taskData, []string{"Collateral(SWAN-ETH):", collateralBalance})

		taskData = append(taskData, []string{"UBI FLAG:", ubiFlag})
		taskData = append(taskData, []string{"Beneficiary Address:", beneficiaryAddress})

		header := []string{"Owner:", ownerAddress}
		NewVisualTable(header, taskData, []RowColor{}).Generate(false)
		if localNodeId != chainNodeId {
			fmt.Printf("NodeId mismatch, local node id: %s, chain node id: %s.\n", localNodeId, chainNodeId)
		}
		return nil
	},
}

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new cp",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ownerAddress",
			Usage: "Specify a OwnerAddress",
		},
		&cli.StringFlag{
			Name:  "beneficiaryAddress",
			Usage: "Specify a beneficiaryAddress to receive rewards. If not specified, use the same address as ownerAddress",
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		beneficiaryAddress := cctx.String("beneficiaryAddress")
		if strings.TrimSpace(beneficiaryAddress) == "" {
			beneficiaryAddress = ownerAddress
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			return fmt.Errorf("get rpc url failed, error: %v", err)
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return fmt.Errorf("setup wallet failed, error: %v", err)
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			return fmt.Errorf("the address: %s, private key %v", ownerAddress, wallet.ErrKeyInfoNotFound)
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			return fmt.Errorf("dial rpc connect failed, error: %v", err)
		}
		defer client.Close()

		privateKey, err := crypto.HexToECDSA(ki.PrivateKey)
		if err != nil {
			return err
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("failed to cast public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
		if err != nil {
			return err
		}

		suggestGasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}

		chainId, _ := client.ChainID(context.Background())
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			return err
		}

		auth.Nonce = big.NewInt(int64(nonce))
		suggestGasPrice = suggestGasPrice.Mul(suggestGasPrice, big.NewInt(3))
		suggestGasPrice = suggestGasPrice.Div(suggestGasPrice, big.NewInt(2))
		auth.GasFeeCap = suggestGasPrice
		auth.Context = context.Background()

		nodeID := computing.GetNodeId(cpRepoPath)
		multiAddresses := conf.GetConfig().API.MultiAddress
		var ubiTaskFlag uint8
		if conf.GetConfig().UBI.UbiTask {
			ubiTaskFlag = 1
		}

		contractAddress, tx, _, err := account.DeployAccount(auth, client, nodeID, []string{multiAddresses}, ubiTaskFlag, common.HexToAddress(beneficiaryAddress))
		if err != nil {
			return fmt.Errorf("deploy cp account contract failed, error: %v", err)
		}

		err = os.WriteFile(filepath.Join(cpRepoPath, "account"), []byte(contractAddress.Hex()), 0666)
		if err != nil {
			return fmt.Errorf("write cp account contract address failed, error: %v", err)
		}

		fmt.Printf("Contract deployed! Address: %s\n", contractAddress.Hex())
		fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())

		var blockNumber string
		timeout := time.After(3 * time.Minute)
		ticker := time.Tick(3 * time.Second)
		for {
			select {
			case <-timeout:
				return fmt.Errorf("timeout waiting for transaction confirmation, tx: %s", tx.Hash().Hex())
			case <-ticker:
				receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					if errors.Is(err, ethereum.NotFound) {
						continue
					}
					return err
				}

				if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
					blockNumber = receipt.BlockNumber.String()
					err := DoSend(contractAddress.Hex(), blockNumber)
					if err != nil {
						return err
					}
					fmt.Println("cp successfully initialized, you can now start it with 'computing-provider run'")
					return nil
				} else if receipt != nil && receipt.Status == 0 {
					return err
				}
			}
		}
	},
}

var accountCmd = &cli.Command{
	Name:  "account",
	Usage: "Manage account info of CP",
	Subcommands: []*cli.Command{
		changeMultiAddressCmd,
		changeOwnerAddressCmd,
		changeBeneficiaryAddressCmd,
		changeUbiFlagCmd,
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
		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a multiAddress")
		}

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		multiAddr := cctx.Args().Get(0)
		if strings.TrimSpace(multiAddr) == "" {
			return fmt.Errorf("failed to parse : %s", multiAddr)
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %v,", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		submitUBIProofTx, err := cpStub.ChangeMultiAddress([]string{multiAddr})
		if err != nil {
			logs.GetLogger().Errorf("change multi-addr tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeMultiAddress: %s", submitUBIProofTx)

		return nil
	},
}

var changeOwnerAddressCmd = &cli.Command{
	Name:      "changeOwnerAddress",
	Usage:     "Update OwnerAddress of CP",
	ArgsUsage: "[newOwnerAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("oldOwnerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a ownerAddress")
		}

		newOwnerAddr := cctx.Args().Get(0)
		if strings.TrimSpace(newOwnerAddr) == "" {
			return fmt.Errorf("failed to parse : %s", newOwnerAddr)
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the old owner address: %s, private key %v,", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		submitUBIProofTx, err := cpStub.ChangeOwnerAddress(common.HexToAddress(newOwnerAddr))
		if err != nil {
			logs.GetLogger().Errorf("change owner address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeOwnerAddress: %s", submitUBIProofTx)

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

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a beneficiaryAddress")
		}

		beneficiaryAddress := cctx.Args().Get(0)
		if strings.TrimSpace(beneficiaryAddress) == "" {
			return fmt.Errorf("failed to parse target address: %s", beneficiaryAddress)
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %v. Please import the address into the wallet", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}
		newQuota := big.NewInt(int64(0))
		newExpiration := big.NewInt(int64(0))
		changeBeneficiaryAddressTx, err := cpStub.ChangeBeneficiary(common.HexToAddress(beneficiaryAddress), newQuota, newExpiration)
		if err != nil {
			logs.GetLogger().Errorf("change owner address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("changeBeneficiaryAddress Transaction hash: %s", changeBeneficiaryAddressTx)
		return nil
	},
}

var changeUbiFlagCmd = &cli.Command{
	Name:      "changeUbiFlag",
	Usage:     "Update ubiFlag of CP (0:Reject, 1:Accept)",
	ArgsUsage: "[ubiFlag]",
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

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a beneficiaryAddress")
		}

		ubiFlag := cctx.Args().Get(0)
		if strings.TrimSpace(ubiFlag) == "" {
			return fmt.Errorf("ubiFlag is not empty")
		}

		if strings.TrimSpace(ubiFlag) != "0" && strings.TrimSpace(ubiFlag) != "1" {
			return fmt.Errorf("ubiFlag must be 0 or 1")
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %v. Please import the address into the wallet", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		newUbiFlag, _ := strconv.ParseUint(strings.TrimSpace(ubiFlag), 10, 64)

		changeBeneficiaryAddressTx, err := cpStub.ChangeUbiFlag(uint8(newUbiFlag))
		if err != nil {
			logs.GetLogger().Errorf("change ubi flag tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeUbiFlag Transaction hash: %s", changeBeneficiaryAddressTx)
		return nil
	},
}

func DoSend(contractAddr, height string) error {
	type ContractReq struct {
		Addr   string `json:"addr"`
		Height int    `json:"height"`
	}
	h, _ := strconv.ParseInt(height, 10, 64)
	var contractReq ContractReq
	contractReq.Addr = contractAddr
	contractReq.Height = int(h)

	jsonData, err := json.Marshal(contractReq)
	if err != nil {
		logs.GetLogger().Errorf("JSON encoding failed: %v", err)
		return err
	}

	resp, err := http.Post(conf.GetConfig().UBI.UbiUrl+"/contracts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logs.GetLogger().Errorf("POST request failed: %v", err)
		return err
	}
	defer resp.Body.Close()

	var resultResp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data,omitempty"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.GetLogger().Errorf("read response failed: %v", err)
		return err
	}
	err = json.Unmarshal(body, &resultResp)
	if err != nil {
		logs.GetLogger().Errorf("response convert to json failed: %v", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("register cp to ubi hub failed, error: %s", resultResp.Msg)
	}
	return nil
}
