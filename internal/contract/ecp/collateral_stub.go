package ecp

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/token"
	"github.com/swanchain/go-computing-provider/internal/models"
	"math/big"
	"strings"
	"time"
)

type CollateralStub struct {
	client           *ethclient.Client
	collateral       *EcpCollateral
	privateK         string
	publicK          string
	cpAccountAddress string
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

func NewCollateralStub(client *ethclient.Client, options ...CollateralOption) (*CollateralStub, error) {
	stub := &CollateralStub{}
	for _, option := range options {
		option(stub)
	}

	collateralAddress := common.HexToAddress(conf.GetConfig().CONTRACT.ZkCollateral)
	collateralClient, err := NewEcpCollateral(collateralAddress, client)
	if err != nil {
		return nil, fmt.Errorf("ECP create collateral contract client, error: %+v", err)
	}

	stub.collateral = collateralClient
	stub.client = client
	return stub, nil
}

func (s *CollateralStub) Deposit(cpAccountAddress string, amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	if cpAccountAddress == "" || len(strings.TrimSpace(cpAccountAddress)) == 0 {
		cpAccountAddress, err = contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create transaction, error: %+v", publicAddress, err)
	}

	tokenClient, err := token.NewToken(common.HexToAddress(conf.GetConfig().CONTRACT.SwanToken), s.client)
	if err != nil {
		return "", fmt.Errorf("create token client failed, error: %v", err)
	}

	approveTx, err := tokenClient.Approve(txOptions, common.HexToAddress(conf.GetConfig().CONTRACT.Collateral), amount)
	if err != nil {
		return "", fmt.Errorf("token approve failed, error: %v", err)
	}

	timeout := time.After(3 * time.Minute)
	ticker := time.Tick(3 * time.Second)
	for {
		select {
		case <-timeout:
			return "", fmt.Errorf("timeout waiting for token approve transaction confirmation, tx: %s", approveTx.Hash().String())
		case <-ticker:
			receipt, err := s.client.TransactionReceipt(context.Background(), approveTx.Hash())
			if err != nil {
				if errors.Is(err, ethereum.NotFound) {
					continue
				}
			}

			if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
				logs.GetLogger().Infof("token approve transaction hash: %s", approveTx.Hash().String())
				depositTxOptions, err := s.createTransactOpts()
				if err != nil {
					return "", fmt.Errorf("address: %s, ECP collateral client create transaction, error: %+v", publicAddress, err)
				}

				transaction, err := s.collateral.Deposit(depositTxOptions, common.HexToAddress(cpAccountAddress), amount)
				if err != nil {
					return "", fmt.Errorf("address: %s, ECP collateral client create deposit tx error: %+v", publicAddress, err)
				}
				return transaction.Hash().String(), nil

			} else if receipt != nil && receipt.Status == 0 {
				return "", fmt.Errorf("swan token approve transaction execution failed, tx: %s", approveTx.Hash().String())
			}
		}
	}
}

func (s *CollateralStub) Withdraw(cpAccountAddress string, amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create transaction, error: %+v", publicAddress, err)
	}

	if cpAccountAddress == "" || len(strings.TrimSpace(cpAccountAddress)) == 0 {
		cpAccountAddress, err = contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
	}

	transaction, err := s.collateral.Withdraw(txOptions, common.HexToAddress(cpAccountAddress), amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create withdraw tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *CollateralStub) CpInfo() (models.EcpCollateralInfo, error) {
	var cpInfo models.EcpCollateralInfo

	if s.cpAccountAddress == "" || len(strings.TrimSpace(s.cpAccountAddress)) == 0 {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			return models.EcpCollateralInfo{}, fmt.Errorf("get cp account contract address failed, error: %v", err)
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
