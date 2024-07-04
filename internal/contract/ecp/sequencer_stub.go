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
	"math/big"
	"strings"
)

type SequencerStub struct {
	client           *ethclient.Client
	sequencer        *EcpSequencer
	privateK         string
	publicK          string
	cpAccountAddress string
}

type SequencerOption func(*SequencerStub)

func WithSequencerPrivateKey(pk string) SequencerOption {
	return func(obj *SequencerStub) {
		obj.privateK = pk
	}
}

func WithSequencerCpAccountAddress(cpAccountAddress string) SequencerOption {
	return func(obj *SequencerStub) {
		obj.cpAccountAddress = cpAccountAddress
	}
}

func NewSequencerStub(client *ethclient.Client, options ...SequencerOption) (*SequencerStub, error) {
	stub := &SequencerStub{}
	for _, option := range options {
		option(stub)
	}

	sequencerAddress := common.HexToAddress(conf.GetConfig().CONTRACT.Sequencer)
	sequencerClient, err := NewEcpSequencer(sequencerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("ECP create sequencer contract client, error: %+v", err)
	}

	stub.sequencer = sequencerClient
	stub.client = client
	return stub, nil
}

func (s *SequencerStub) Deposit(cpAccountAddress string, amount *big.Int) (string, error) {
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

	txOptions, err := s.createTransactOpts(amount, true)
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP sequencer client create transaction, error: %+v", publicAddress, err)
	}
	transaction, err := s.sequencer.Deposit(txOptions, common.HexToAddress(cpAccountAddress))
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP sequencer client create deposit tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *SequencerStub) Withdraw(cpAccountAddress string, amount *big.Int) (string, error) {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return "", err
	}

	txOptions, err := s.createTransactOpts(nil, false)
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create transaction, error: %+v", publicAddress, err)
	}

	if cpAccountAddress == "" || len(strings.TrimSpace(cpAccountAddress)) == 0 {
		cpAccountAddress, err = contract.GetCpAccountAddress()
		if err != nil {
			return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
		}
	}

	transaction, err := s.sequencer.Withdraw(txOptions, common.HexToAddress(cpAccountAddress), amount)
	if err != nil {
		return "", fmt.Errorf("address: %s, ECP collateral client create withdraw tx error: %+v", publicAddress, err)
	}
	return transaction.Hash().String(), nil
}

func (s *SequencerStub) privateKeyToPublicKey() (common.Address, error) {
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

func (s *SequencerStub) createTransactOpts(amount *big.Int, isDeposit bool) (*bind.TransactOpts, error) {
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
	if isDeposit {
		txOptions.Value = amount
	}

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
