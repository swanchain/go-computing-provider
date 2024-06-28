package ecp

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/token"
	"math/big"
	"strings"
	"sync"
	"time"
)

type TaskStub struct {
	client          *ethclient.Client
	task            *Task
	privateK        string
	publicK         string
	ContractAddress string
	nonceX          uint64
	taskL           sync.Mutex
}

type TaskOption func(*TaskStub)

func WithTaskPrivateKey(pk string) TaskOption {
	return func(obj *TaskStub) {
		obj.privateK = pk
	}
}

func WithTaskContractAddress(contractAddress string) TaskOption {
	return func(obj *TaskStub) {
		obj.ContractAddress = contractAddress
	}
}

func NewTaskStub(client *ethclient.Client, options ...TaskOption) (*TaskStub, error) {
	stub := &TaskStub{}
	for _, option := range options {
		option(stub)
	}

	if stub.ContractAddress == "" {
		return nil, errors.New("missing task contract address")
	}

	cpAccountAddress := common.HexToAddress(stub.ContractAddress)
	taskClient, err := NewTask(cpAccountAddress, client)
	if err != nil {
		return nil, fmt.Errorf("create task contract client, error: %+v", err)
	}

	stub.task = taskClient
	stub.client = client
	return stub, nil
}

func (s *TaskStub) SubmitUBIProof(taskId, proof string, timeOut int64) (string, error) {
	var err error
	var submitProofTxHash string
	var flag bool

	timeOutCh := time.After(time.Second * time.Duration(timeOut))
outerLoop:
	for {
		select {
		case <-timeOutCh:
			err = fmt.Errorf("Proof submission timed out")
			break outerLoop
		default:
			time.Sleep(3 * time.Second)
			if !flag {
				err = s.getNonce()
				if err != nil {
					logs.GetLogger().Warnf("taskId: %s, get nonce: %s, retrying", taskId, ParseError(err))
					continue
				}
			}

			txOptions, err := s.createTransactOpts(int64(s.nonceX))
			if err != nil {
				logs.GetLogger().Warnf("taskId: %s, create transaction opts failed, error: %s", taskId, ParseError(err))
				continue
			}
			transaction, err := s.task.SubmitProof(txOptions, proof)
			if err != nil {
				if err.Error() == "replacement transaction underpriced" {
					s.IncrementNonce()
					flag = true
				} else if strings.Contains(err.Error(), "next nonce") {
					err = s.getNonce()
					if err != nil {
						logs.GetLogger().Warnf("taskId: %s, get nonce: %s, retrying", taskId, ParseError(err))
						flag = false
						continue
					}
				} else {
					logs.GetLogger().Warnf("taskId: %s SubmitUBIProof failed, error: %s", taskId, ParseError(err))
					continue
				}
			}
			if transaction != nil {
				submitProofTxHash = transaction.Hash().String()
				break outerLoop
			} else {
				logs.GetLogger().Warnf("taskId: %s submitProofTxHash is nil, retrying", taskId)
			}
		}
	}
	return submitProofTxHash, err
}

func (s *TaskStub) GetTaskInfo() (ECPTaskTaskInfo, error) {
	return s.task.GetTaskInfo(&bind.CallOpts{})
}

func (s *TaskStub) privateKeyToPublicKey() (common.Address, error) {
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

func (s *TaskStub) createTransactOpts(nonce int64) (*bind.TransactOpts, error) {
	suggestGasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("task client retrieves the currently suggested gas price, error: %+v", err)
	}

	chainId, err := s.client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("task client get networkId, error: %+v", err)
	}

	privateKey, err := crypto.HexToECDSA(s.privateK)
	if err != nil {
		return nil, fmt.Errorf("parses private key error: %+v", err)
	}

	txOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("collateral client create transaction, error: %+v", err)
	}
	txOptions.Nonce = big.NewInt(nonce)
	suggestGasPrice = suggestGasPrice.Mul(suggestGasPrice, big.NewInt(3))
	suggestGasPrice = suggestGasPrice.Div(suggestGasPrice, big.NewInt(2))
	txOptions.GasFeeCap = suggestGasPrice
	txOptions.Context = context.Background()
	return txOptions, nil
}

func (s *TaskStub) getNonce() error {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return err
	}
	nonce, err := s.client.NonceAt(context.Background(), publicAddress, nil)
	if err != nil {
		return fmt.Errorf("address: %s, collateral client get nonce error: %s", publicAddress, ParseError(err))
	}
	s.taskL.Lock()
	defer s.taskL.Unlock()
	s.nonceX = nonce
	return nil
}

func (s *TaskStub) IncrementNonce() {
	s.taskL.Lock()
	defer s.taskL.Unlock()
	s.nonceX++
}

// GetReward  status: 1: Challenged  2: Slashed  3: rewarded
func (s *TaskStub) GetReward() (status int, rewardTx string, challengeTx string, slashTx string, reward string, err error) {
	reward = "0.0"
	taskInfo, err := s.GetTaskInfo()
	if err != nil {
		return 0, "", "", "", reward, err
	}

	if taskInfo.ChallengeTx != "" {
		return 1, "", taskInfo.ChallengeTx, "", reward, nil
	}

	if taskInfo.SlashTx != "" {
		return 2, "", "", taskInfo.SlashTx, reward, nil
	}

	if taskInfo.RewardTx != "" {
		receipt, err := s.client.TransactionReceipt(context.Background(), common.HexToHash(taskInfo.RewardTx))
		if err != nil {
			return 0, taskInfo.RewardTx, "", "", reward, err
		}
		contractAbi, err := abi.JSON(strings.NewReader(token.MainMetaData.ABI))
		if err != nil {
			return 0, taskInfo.RewardTx, "", "", reward, err
		}

		for _, l := range receipt.Logs {
			event := struct {
				From  common.Address
				To    common.Address
				Value *big.Int
			}{}
			if err := contractAbi.UnpackIntoInterface(&event, "Transfer", l.Data); err != nil {
				continue
			}

			if len(l.Topics) == 3 && l.Topics[0] == contractAbi.Events["Transfer"].ID {
				reward = contract.BalanceToStr(event.Value)
			}
		}
		return 3, taskInfo.RewardTx, "", "", reward, nil
	}
	return 0, "", "", "", reward, nil
}

func ParseError(err error) string {
	if strings.Contains(err.Error(), "503") {
		return "503 Service Temporarily Unavailable"
	}
	return err.Error()
}
