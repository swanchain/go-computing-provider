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
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/account"
	"github.com/swanchain/go-computing-provider/internal/pkg"
	"github.com/swanchain/go-computing-provider/internal/v1/routers"
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
		mlog.Info("Start in computing provider mode.")

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		ProjectInit(cpRepoPath)

		r := gin.New()
		r.Use(pkg.LoggerMiddleware(), cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			ValidateHeaders: false,
		}), gin.Recovery())
		pprof.Register(r)

		v1 := r.Group("/api/v1")
		routers.CPApiV1(v1.Group("/computing"))

		shutdownChan := make(chan struct{})
		httpStopper, err := pkg.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port))
		if err != nil {
			mlog.Fatal("failed to start cp-api endpoint: %s", err)
		}

		finishCh := pkg.MonitorShutdown(shutdownChan,
			pkg.ShutdownHandler{Component: "cp-api", StopFunc: httpStopper},
		)
		<-finishCh

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
			mlog.Fatal(err)
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

		nodeID := pkg.GetNodeId(cpRepoPath)
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

		localNodeId := pkg.GetNodeId(cpPath)

		k8sService := pkg.NewK8sService()
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
		pkg.NewVisualTable(header, taskData, []pkg.RowColor{}).Generate(false)
		if localNodeId != chainNodeId {
			fmt.Printf("NodeId mismatch, local node id: %s, chain node id: %s.\n", localNodeId, chainNodeId)
		}
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
		mlog.Errorf("JSON encoding failed: %v", err)
		return err
	}

	resp, err := http.Post(conf.GetConfig().UBI.UbiUrl+"/contracts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		mlog.Errorf("POST request failed: %v", err)
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
		mlog.Errorf("read response failed: %v", err)
		return err
	}
	err = json.Unmarshal(body, &resultResp)
	if err != nil {
		mlog.Errorf("response convert to json failed: %v", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("register cp to ubi hub failed, error: %s", resultResp.Msg)
	}
	return nil
}
