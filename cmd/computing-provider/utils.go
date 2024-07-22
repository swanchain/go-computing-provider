package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/wallet"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func createAccount(cpRepoPath, ownerAddress, beneficiaryAddress string, workerAddress string, taskTypes []uint8) error {
	chainUrl, err := conf.GetRpcByNetWorkName()
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
	auth.GasFeeCap = suggestGasPrice
	auth.Context = context.Background()

	nodeID := computing.GetNodeId(cpRepoPath)
	multiAddresses := conf.GetConfig().API.MultiAddress

	if strings.Contains(conf.GetConfig().API.MultiAddress, "<") || strings.Contains(conf.GetConfig().API.MultiAddress, "PUBLIC") {
		return fmt.Errorf("the multi-address field needs to be configured, by modify config file or computing-provider init")
	}

	contractAddress, tx, _, err := account.DeployAccount(auth, client, nodeID, []string{multiAddresses}, common.HexToAddress(beneficiaryAddress),
		common.HexToAddress(workerAddress), common.HexToAddress(conf.GetConfig().CONTRACT.CpAccountRegister), taskTypes)
	if err != nil {
		return fmt.Errorf("deploy cp account contract failed, error: %v", err)
	}
	cpAccountAddress := contractAddress.Hex()

	err = os.WriteFile(filepath.Join(cpRepoPath, "account"), []byte(cpAccountAddress), 0666)
	if err != nil {
		return fmt.Errorf("write cp account contract address to fie failed, error: %v", err)
	}

	var cpInfo = new(models.CpInfoEntity)
	cpInfo.NodeId = nodeID
	cpInfo.OwnerAddress = ownerAddress
	cpInfo.Beneficiary = beneficiaryAddress
	cpInfo.WorkerAddress = workerAddress
	cpInfo.ContractAddress = cpAccountAddress
	cpInfo.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	cpInfo.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	cpInfo.MultiAddresses = []string{multiAddresses}
	cpInfo.TaskTypes = taskTypes
	if err = computing.NewCpInfoService().SaveCpInfoEntity(cpInfo); err != nil {
		return fmt.Errorf("save cp info to db failed, error: %v", err)
	}

	fmt.Printf("Contract deployed! Address: %s\n", cpAccountAddress)
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
	fmt.Println("computing-provider account is created successfully! You can now start it with 'computing-provider run' or 'computing-provider ubi daemon'")
	return nil
}

func getVerifyAccountClient(ownerAddress string) (*ethclient.Client, *account.CpStub, error) {
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return nil, nil, fmt.Errorf("get rpc url failed, error: %v", err)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("dial rpc connect failed, error: %v", err)
	}

	cpStub, err := account.NewAccountStub(client)
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	cpAccount, err := cpStub.GetCpAccountInfo()
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("get cpAccount failed, error: %v", err)
	}
	if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
		client.Close()
		return nil, nil, fmt.Errorf("Only the owner can change CP account owner address, the CP account is: %s, the owner should be %s", cpAccount.Contract, cpAccount.OwnerAddress)
	}
	return client, cpStub, nil
}
