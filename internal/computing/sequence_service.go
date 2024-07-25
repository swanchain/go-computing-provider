package computing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/wallet"
	"golang.org/x/xerrors"
	"io"
	"net/http"
	"time"
)

type Sequencer struct {
	url string
}

const (
	token = "/token"
	task  = "/tasks"
)

var tokenCache string

func NewSequencer() *Sequencer {
	return &Sequencer{
		url: conf.GetConfig().UBI.SequencerUrl,
	}
}

func (s *Sequencer) getToken() error {
	accountInfo, err := account.GetAccountInfo()
	if err != nil {
		return err
	}
	var timestamp = time.Now().Unix()
	signMsg, err := signMessage(fmt.Sprintf("%s%d", accountInfo.Contract, timestamp), accountInfo.OwnerAddress)
	if err != nil {
		return err
	}

	var data struct {
		CpAddr    string `json:"cp_addr"`
		Timestamp int64  `json:"timestamp"`
		Sign      string `json:"sign"`
	}

	data.CpAddr = accountInfo.Contract
	data.Timestamp = timestamp
	data.Sign = signMsg

	var header http.Header
	header.Set("Content-Type", "application/json")
	header.Set("Authorization", tokenCache)

	var tokenResp TokenResp
	err = NewHttpClient(s.url, header).PostJSON(token, data, &tokenResp)
	if err != nil {
		return fmt.Errorf("failed to get token, error: %v", err)
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
		if err := s.getToken(); err != nil {
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
		logs.GetLogger().Infof("SendTaskProof: code: %d, result: %s", resp.StatusCode, string(body))
		return spr, fmt.Errorf("response status: %d", resp.StatusCode)
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

func (s *Sequencer) QueryTask(taskIds ...int64) (TaskListResp, error) {
	if tokenCache == "" {
		if err := s.getToken(); err != nil {
			return TaskListResp{}, fmt.Errorf("failed to get token, error: %v", err)
		}
	}

	reqData, err := json.Marshal(taskIds)
	if err != nil {
		return TaskListResp{}, err
	}

	var reqUrl = s.url + task + fmt.Sprintf("?ids=%s", string(reqData))
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return TaskListResp{}, fmt.Errorf("error creating request: %v", err)
	}

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

	logs.GetLogger().Infof("QueryTask: code: %d, result: %s", resp.StatusCode, string(body))

	var taskListResp TaskListResp
	if resp.StatusCode != http.StatusOK {
		logs.GetLogger().Infof("QueryTask: code: %d, result: %s", resp.StatusCode, string(body))
		return taskListResp, fmt.Errorf("response status: %d", resp.StatusCode)
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
	Id            int    `json:"id"`
	Type          int    `json:"type"`
	InputParam    string `json:"input_param"`
	VerifyParam   string `json:"verify_param"`
	ResourceType  int    `json:"resource_type"`
	Deadline      int    `json:"deadline"`
	Proof         string `json:"proof"`
	CheckCode     string `json:"check_code"`
	Reward        string `json:"reward"`
	Status        string `json:"status"`
	SequenceCid   string `json:"sequence_cid"`
	SettlementCid string `json:"settlement_cid"`
}
