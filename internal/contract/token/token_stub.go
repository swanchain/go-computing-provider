package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/conf"
	"math/big"
	"strings"
)

type Stub struct {
	client             *ethclient.Client
	token              *Token
	privateK           string
	publicK            string
	collateralContract string
}

type Option func(*Stub)

func WithPrivateKey(pk string) Option {
	return func(obj *Stub) {
		obj.privateK = pk
	}
}

func WithPublicKey(pk string) Option {
	return func(obj *Stub) {
		obj.publicK = pk
	}
}

func WithCollateralContract(collateralContract string) Option {
	return func(obj *Stub) {
		obj.collateralContract = collateralContract
	}
}

func NewTokenStub(client *ethclient.Client, options ...Option) (*Stub, error) {
	stub := &Stub{}
	for _, option := range options {
		option(stub)
	}

	tokenAddress := common.HexToAddress(conf.GetConfig().CONTRACT.SwanToken)
	tokenClient, err := NewToken(tokenAddress, client)
	if err != nil {
		return nil, fmt.Errorf("create collateral contract client, error: %+v", err)
	}

	stub.token = tokenClient
	stub.client = client
	return stub, nil
}

func (s *Stub) BalanceOf() (string, error) {
	if len(strings.TrimSpace(s.publicK)) == 0 {
		return "", fmt.Errorf("wallet address must be not empty")
	}

	publicAddress := common.HexToAddress(s.publicK)

	balance, err := s.token.BalanceOf(&bind.CallOpts{}, publicAddress)
	if err != nil {
		return "", fmt.Errorf("address: %s, read token contract balance, error: %+v", publicAddress, err)
	}
	var ethValue string
	if balance.String() == "0" {
		ethValue = "0.0000"
	} else {
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		etherQuotient := new(big.Float).Quo(fbalance, new(big.Float).SetInt(big.NewInt(1e18)))
		ethValue = etherQuotient.Text('f', 4)
	}
	return ethValue, nil
}

func (s *Stub) Approve(amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, collateral client create transaction, error: %+v", publicAddress, err)
	}

	if s.collateralContract == "" {
		return "", fmt.Errorf("must be set a collateral contract address")
	}

	transaction, err := s.token.Approve(txOptions, common.HexToAddress(s.collateralContract), amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, token contract approve, error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *Stub) Transfer(to string, amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts()
	if err != nil {
		return "", fmt.Errorf("address: %s, collateral client create transaction, error: %+v", publicAddress, err)
	}

	toAddress := common.HexToAddress(to)

	transaction, err := s.token.Transfer(txOptions, toAddress, amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, token contract transfer, error: %+v", publicAddress, err)
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
