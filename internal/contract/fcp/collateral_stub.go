package fcp

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

type Stub struct {
	client           *ethclient.Client
	collateral       *FcpCollateral
	privateK         string
	publicK          string
	cpAccountAddress string
}

type Option func(*Stub)

func WithPrivateKey(pk string) Option {
	return func(obj *Stub) {
		obj.privateK = pk
	}
}

func WithCpAccountAddress(cpAccountAddress string) Option {
	return func(obj *Stub) {
		obj.cpAccountAddress = cpAccountAddress
	}
}

func NewCollateralStub(client *ethclient.Client, options ...Option) (*Stub, error) {
	stub := &Stub{}
	for _, option := range options {
		option(stub)
	}

	collateralAddress := common.HexToAddress(conf.GetConfig().CONTRACT.Collateral)
	collateralClient, err := NewFcpCollateral(collateralAddress, client)
	if err != nil {
		return nil, fmt.Errorf("create fcp collateral contract client, error: %+v", err)
	}

	stub.collateral = collateralClient
	stub.client = client
	return stub, nil
}

func (s *Stub) Deposit(amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, FCP collateral client create transaction, error: %+v", publicAddress, err)
	}

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}
	transaction, err := s.collateral.Deposit(txOptions, common.HexToAddress(s.cpAccountAddress), amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, FCP collateral client create deposit tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *Stub) CollateralInfo() (models.FcpCollateralInfo, error) {
	var cpInfo models.FcpCollateralInfo

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return cpInfo, fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}

	collateralInfo, err := s.collateral.CpInfo(&bind.CallOpts{}, common.HexToAddress(s.cpAccountAddress))
	if err != nil {
		return cpInfo, fmt.Errorf("address: %s, get FCP collateral info error: %+v", s.cpAccountAddress, err)
	}

	cpInfo.CpAddress = collateralInfo.CpAccount.Hex()
	cpInfo.AvailableBalance = contract.BalanceToStr(collateralInfo.AvailableBalance)
	cpInfo.LockedCollateral = contract.BalanceToStr(collateralInfo.LockedCollateral)
	cpInfo.Status = collateralInfo.Status
	return cpInfo, nil
}

func (s *Stub) Withdraw(amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, FCP collateral client create transaction, error: %+v", publicAddress, err)
	}

	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
	}

	transaction, err := s.collateral.Withdraw(txOptions, common.HexToAddress(cpAccountAddress), amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, FCP collateral withdraw tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *Stub) privateKeyToPublicKey() (common.Address, error) {
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

func (s *Stub) createTransactOpts() (*bind.TransactOpts, error) {
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
