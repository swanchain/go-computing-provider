package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/models"
	"math/big"
	"strings"
)

type CpStub struct {
	client          *ethclient.Client
	account         *Account
	privateK        string
	publicK         string
	ContractAddress string
}

type CpOption func(*CpStub)

func WithCpPrivateKey(pk string) CpOption {
	return func(obj *CpStub) {
		obj.privateK = pk
	}
}

func WithContractAddress(contractAddress string) CpOption {
	return func(obj *CpStub) {
		obj.ContractAddress = contractAddress
	}
}

func NewAccountStub(client *ethclient.Client, options ...CpOption) (*CpStub, error) {
	stub := &CpStub{}
	for _, option := range options {
		option(stub)
	}

	if stub.ContractAddress == "" || len(strings.TrimSpace(stub.ContractAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return nil, fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
		stub.ContractAddress = cpAccountAddress
	}

	cpAccountAddress := common.HexToAddress(stub.ContractAddress)
	accountClient, err := NewAccount(cpAccountAddress, client)
	if err != nil {
		return nil, fmt.Errorf("create cp account contract client, error: %+v", err)
	}

	stub.account = accountClient
	stub.client = client
	return stub, nil
}

func (s *CpStub) ChangeMultiAddress(newMultiAddress []string) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeMultiaddrs(txOptions, newMultiAddress)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client ChangeMultiaddrs tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) ChangeOwnerAddress(newOwner common.Address) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeOwnerAddress(txOptions, newOwner)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create ChangeOwnerAddress tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) ChangeBeneficiary(newBeneficiary common.Address) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeBeneficiary(txOptions, newBeneficiary)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create ChangeBeneficiary tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) ChangeTaskTypes(newTaskTypes []uint8) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeTaskTypes(txOptions, newTaskTypes)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create ChangeTaskTypes tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) ChangeWorkerAddress(newWorkerAddress common.Address) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create transaction, error: %+v", publicAddress, err)
	}

	transaction, err := s.account.ChangeWorker(txOptions, newWorkerAddress)
	if err != nil {
		return "", fmt.Errorf("address: %s, cpAccount client create ChangeWorkerAddress tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CpStub) GetCpAccountInfo() (models.Account, error) {
	cpAccount, err := s.account.GetAccount(&bind.CallOpts{})
	if err != nil {
		return models.Account{}, fmt.Errorf("cpAccount client GetCpAccountInfo error: %+v", err)
	}
	var account models.Account
	account.OwnerAddress = cpAccount.Owner.Hex()
	account.NodeId = cpAccount.NodeId
	account.MultiAddresses = cpAccount.MultiAddresses
	account.TaskTypes = cpAccount.TaskTypes
	account.Beneficiary = cpAccount.Beneficiary.Hex()
	account.WorkerAddress = cpAccount.Worker.Hex()
	account.Version = cpAccount.Version
	account.Contract = s.ContractAddress
	return account, nil
}

func (s *CpStub) privateKeyToPublicKey() (common.Address, error) {
	if len(strings.TrimSpace(s.privateK)) == 0 {
		return common.Address{}, fmt.Errorf("wallet address private key must be not empty")
	}

	privateKey, err := crypto.HexToECDSA(s.privateK)
	if err != nil {
		return common.Address{}, fmt.Errorf("parses private key error: %+v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func (s *CpStub) createTransactOpts() (*bind.TransactOpts, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return nil, err
	}

	nonce, err := s.client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client get nonce error: %+v", publicAddress, err)
	}

	suggestGasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client retrieves the currently suggested gas price, error: %+v", publicAddress, err)
	}

	chainId, err := s.client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client get networkId, error: %+v", publicAddress, err)
	}

	privateKey, err := crypto.HexToECDSA(s.privateK)
	if err != nil {
		return nil, fmt.Errorf("parses private key error: %+v", err)
	}

	txOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("address: %s, collateral client create transaction, error: %+v", publicAddress, err)
	}
	txOptions.Nonce = big.NewInt(int64(nonce))
	suggestGasPrice = suggestGasPrice.Mul(suggestGasPrice, big.NewInt(3))
	suggestGasPrice = suggestGasPrice.Div(suggestGasPrice, big.NewInt(2))
	txOptions.GasFeeCap = suggestGasPrice
	txOptions.Context = context.Background()
	return txOptions, nil
}

func GetAccountInfo() (models.Account, error) {
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return models.Account{}, fmt.Errorf("failed to get rpc url, error: %v", err)
	}

	client, err := contract.GetEthClient(chainUrl)
	if err != nil {
		return models.Account{}, fmt.Errorf("failed to dial rpc connect, error: %v", err)
	}
	defer client.Close()

	cpStub, err := NewAccountStub(client)
	if err != nil {
		return models.Account{}, err
	}

	return cpStub.GetCpAccountInfo()
}
