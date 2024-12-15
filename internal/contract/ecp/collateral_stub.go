package ecp

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

type CollateralStub struct {
	client           *ethclient.Client
	collateral       *EcpCollateral
	privateK         string
	publicK          string
	cpAccountAddress string
	contract         string
}

type CollateralOption func(*CollateralStub)

func WithPrivateKey(pk string) CollateralOption {
	return func(obj *CollateralStub) {
		obj.privateK = pk
	}
}

func WithPublicKey(pk string) CollateralOption {
	return func(obj *CollateralStub) {
		obj.publicK = pk
	}
}

func WithCpAccountAddress(cpAccountAddress string) CollateralOption {
	return func(obj *CollateralStub) {
		obj.cpAccountAddress = cpAccountAddress
	}
}

func NewCollateralWithUbiZeroStub(client *ethclient.Client, options ...CollateralOption) (*CollateralStub, error) {
	stub := &CollateralStub{}
	for _, option := range options {
		option(stub)
	}

	contractAddr := conf.GetConfig().CONTRACT.ZkCollateralUbiZero
	collateralAddress := common.HexToAddress(contractAddr)
	collateralClient, err := NewEcpCollateral(collateralAddress, client)
	if err != nil {
		return nil, fmt.Errorf("ECP create collateral contract client, error: %+v", err)
	}

	stub.contract = contractAddr
	stub.collateral = collateralClient
	stub.client = client
	return stub, nil
}

func NewCollateralStub(client *ethclient.Client, options ...CollateralOption) (*CollateralStub, error) {
	stub := &CollateralStub{}
	for _, option := range options {
		option(stub)
	}

	contractAddr := conf.GetConfig().CONTRACT.ZkCollateral
	collateralAddress := common.HexToAddress(contractAddr)
	collateralClient, err := NewEcpCollateral(collateralAddress, client)
	if err != nil {
		return nil, fmt.Errorf("ECP create collateral contract client, error: %+v", err)
	}

	stub.contract = contractAddr
	stub.collateral = collateralClient
	stub.client = client
	return stub, nil
}

func (s *CollateralStub) Deposit(amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create tx opts, error: %+v", publicAddress, err)
	}

	transaction, err := s.collateral.Deposit(txOptions, common.HexToAddress(s.cpAccountAddress), amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client deposit tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CollateralStub) Withdraw(amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create transaction, error: %+v", publicAddress, err)
	}

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}

	transaction, err := s.collateral.Withdraw(txOptions, common.HexToAddress(s.cpAccountAddress), amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client withdraw tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CollateralStub) WithdrawRequest(amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create transaction, error: %+v", publicAddress, err)
	}

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("failed to get cp account contract address, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}

	transaction, err := s.collateral.RequestWithdraw(txOptions, common.HexToAddress(s.cpAccountAddress), amount)
	if err != nil {
		return "", fmt.Errorf("failed to request withdraw for ecp, cp account address: %s, error: %+v", s.cpAccountAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CollateralStub) WithdrawView() (models.WithdrawRequest, error) {
	var request models.WithdrawRequest

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return request, fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}

	withdrawRequest, err := s.collateral.ViewWithdrawRequest(&bind.CallOpts{}, common.HexToAddress(s.cpAccountAddress))
	if err != nil {
		return request, fmt.Errorf("failed to view withdraw request for ecp, cp account address: %s, error: %+v", s.cpAccountAddress, err)
	}

	request.RequestBlock = withdrawRequest.RequestBlock.Int64()
	request.Amount = contract.BalanceToStr(withdrawRequest.Amount)
	return request, nil
}

func (s *CollateralStub) WithdrawConfirm() (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("failed to create transaction opts, address: %s, error: %+v", publicAddress, err)
	}

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("failed to get cp account contract address, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}

	transaction, err := s.collateral.ConfirmWithdraw(txOptions, common.HexToAddress(s.cpAccountAddress))
	if err != nil {
		return "", fmt.Errorf("failed to confirm withdraw for ecp, cp account address: %s, error: %+v", s.cpAccountAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CollateralStub) ContractInfo() (models.CollateralContractInfoForECP, error) {
	var contractInfo models.CollateralContractInfoForECP

	collateral, err := s.collateral.GetECPCollateralInfo(&bind.CallOpts{})
	if err != nil {
		return contractInfo, fmt.Errorf("failed to get ecp collaternal contract info, contract address: %s, error: %+v", s.contract, err)
	}

	contractInfo.CollateralToken = collateral.CollateralToken.Hex()
	contractInfo.WithdrawDelay = collateral.WithdrawDelay.Int64()
	return contractInfo, nil
}

func (s *CollateralStub) CpInfo() (models.CpCollateralInfoForECP, error) {
	var cpInfo models.CpCollateralInfoForECP

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return models.CpCollateralInfoForECP{}, fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
		s.cpAccountAddress = cpAccountAddress
	}

	cpCollateralInfo, err := s.collateral.CpInfo(&bind.CallOpts{}, common.HexToAddress(s.cpAccountAddress))
	if err != nil {
		return cpInfo, fmt.Errorf("address: %s, collateral client cpInfo tx error: %+v", s.cpAccountAddress, err)
	}

	cpInfo.CpAddress = cpCollateralInfo.Cp.Hex()
	cpInfo.CollateralBalance = contract.BalanceToStr(cpCollateralInfo.Balance)
	cpInfo.FrozenBalance = contract.BalanceToStr(cpCollateralInfo.FrozenBalance)
	cpInfo.Status = cpCollateralInfo.Status
	return cpInfo, nil
}

func (s *CollateralStub) privateKeyToPublicKey() (common.Address, error) {
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

func (s *CollateralStub) createTransactOpts() (*bind.TransactOpts, error) {
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
