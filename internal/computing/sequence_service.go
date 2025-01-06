package computing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/wallet"
	"golang.org/x/xerrors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Sequencer struct {
	url string
}

const (
	token = "/v1/token"
	task  = "/v1/tasks"
)

var tokenCache string

func NewSequencer() *Sequencer {
	return &Sequencer{
		url: conf.GetConfig().UBI.SequencerUrl,
	}
}

func GetToken() (string, error) {
	cpPath, _ := os.LookupEnv("CP_PATH")
	tokenFileName := filepath.Join(cpPath, "token")

	if tokenCache != "" {
		if err := os.WriteFile(tokenFileName, []byte(tokenCache), 0666); err != nil {
			return "", fmt.Errorf("failed to write sequencer token to file, error: %v", err)
		}
	} else {
		if err := NewSequencer().GetToken(); err != nil {
			return "", fmt.Errorf("failed to get token, error: %v", err)
		}
		if err := os.WriteFile(tokenFileName, []byte(tokenCache), 0666); err != nil {
			return "", fmt.Errorf("failed to write sequencer token to file, error: %v", err)
		}
		return tokenCache, nil
	}
	return tokenCache, nil
}

func (s *Sequencer) GetToken() error {
	accountInfo, err := account.GetAccountInfo()
	if err != nil {
		return err
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return fmt.Errorf("failed to get rpc url, error: %v", err)
	}
	client, err := contract.GetEthClient(chainUrl)
	if err != nil {
		return fmt.Errorf("failed to dial rpc connect, error: %v", err)
	}
	client.Close()

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get block_number, error: %v", err)
	}
	signMsg, err := signMessage(fmt.Sprintf("%s%d", accountInfo.Contract, blockNumber), accountInfo.WorkerAddress)
	if err != nil {
		return err
	}

	var data struct {
		CpAddr      string `json:"cp_addr"`
		BlockNumber uint64 `json:"block_number"`
		Sign        string `json:"sign"`
	}

	data.CpAddr = accountInfo.Contract
	data.BlockNumber = blockNumber
	data.Sign = signMsg

	var header = make(http.Header)
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", tokenCache)

	var tokenResp TokenResp
	err = NewHttpClient(s.url, header).PostJSON(token, data, &tokenResp)
	if err != nil {
		return fmt.Errorf("failed to get token, blockNumber: %d, error: %v", blockNumber, err)
	}

	if tokenResp.Code == 0 {
		tokenCache = tokenResp.Data.Token
	} else {
		return fmt.Errorf(tokenResp.Msg)
	}
	return nil
}

func (s *Sequencer) SendTaskProof(data []byte) (SendProofResp, error) {
	if tokenCache == "" {
		if err := s.GetToken(); err != nil {
			return SendProofResp{}, fmt.Errorf("failed to get token, error: %v", err)
		}
	}

	req, err := http.NewRequest("POST", s.url+task, bytes.NewBuffer(data))
	if err != nil {
		return SendProofResp{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenCache)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return SendProofResp{}, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	newToken := resp.Header.Get("new-token")
	if newToken != "" {
		tokenCache = newToken
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SendProofResp{}, fmt.Errorf("failed to read response: %v", err)
	}

	var spr SendProofResp
	if resp.StatusCode != http.StatusOK {
		if err := s.GetToken(); err != nil {
			return SendProofResp{}, fmt.Errorf("failed to get token, error: %v", err)
		}
		return spr, fmt.Errorf("response status: %d, %s", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &spr)
	if err != nil {
		return spr, fmt.Errorf("failed to response convert to json, error: %v", err)
	}
	if spr.Code == 0 {
		return spr, nil
	}
	return spr, fmt.Errorf(spr.Msg)
}

func (s *Sequencer) QueryTask(taskType int, taskIds ...int64) (TaskListResp, error) {
	if tokenCache == "" {
		if err := s.GetToken(); err != nil {
			return TaskListResp{}, fmt.Errorf("failed to get token, error: %v", err)
		}
	}

	reqData, err := json.Marshal(taskIds)
	if err != nil {
		return TaskListResp{}, err
	}

	var reqUrl string
	if taskType == 3 {
		reqUrl = s.url + task + fmt.Sprintf("?type=%d&uuids=%s", taskType, string(reqData))
	} else {
		reqUrl = s.url + task + fmt.Sprintf("?type=%d&ids=%s", taskType, string(reqData))
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return TaskListResp{}, fmt.Errorf("error creating request: %v", err)
	}

	logs.GetLogger().Infof("tokenCache: %s", tokenCache)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", tokenCache)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TaskListResp{}, fmt.Errorf("failed to sending request: %v", err)
	}
	defer resp.Body.Close()

	newToken := resp.Header.Get("new-token")
	if newToken != "" {
		tokenCache = newToken
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TaskListResp{}, fmt.Errorf("error reading response: %v", err)
	}

	var taskListResp TaskListResp
	if resp.StatusCode != http.StatusOK {
		if err := s.GetToken(); err != nil {
			return taskListResp, fmt.Errorf("failed to get token, error: %v", err)
		}
		return taskListResp, fmt.Errorf("response status: %d, %s", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &taskListResp)
	if err != nil {
		return TaskListResp{}, fmt.Errorf("response convert to json failed: %v", err)
	}

	if taskListResp.Code != 0 {
		return taskListResp, fmt.Errorf(taskListResp.Msg)
	} else {
		return taskListResp, nil
	}
}

func signMessage(msg string, ownerAddress string) (string, error) {
	localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
	if err != nil {
		return "", err
	}
	ki, err := localWallet.FindKey(ownerAddress)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", ownerAddress, wallet.ErrKeyInfoNotFound)
	}

	key, err := crypto.HexToECDSA(ki.PrivateKey)
	if err != nil {
		return "", err
	}

	hash := crypto.Keccak256Hash([]byte(msg))
	sig, err := crypto.Sign(hash.Bytes(), key)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(sig), nil
}

type TokenResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type SendProofResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		BlockHash string `json:"block_hash"`
		Sign      string `json:"sign"`
	} `json:"data"`
}

type TaskListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total int            `json:"total"`
		List  []SequenceTask `json:"list"`
	} `json:"data"`
}

type SequenceTask struct {
	Id                 int    `json:"id"`
	Type               int    `json:"type"`
	InputParam         string `json:"input_param"`
	VerifyParam        string `json:"verify_param"`
	ResourceType       int    `json:"resource_type"`
	Deadline           int    `json:"deadline"`
	Proof              string `json:"proof"`
	CheckCode          string `json:"check_code"`
	Reward             string `json:"reward"`
	Status             string `json:"status"`
	SequenceCid        string `json:"sequence_cid"`
	SettlementCid      string `json:"settlement_cid"`
	SequenceTaskAddr   string `json:"sequence_task_addr"`
	SettlementTaskAddr string `json:"settlement_task_addr"`
}
