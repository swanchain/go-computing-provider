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
	"strings"
	"time"
)

type Sequencer struct {
	url string
}

const (
	token = "/token"
	task  = "/task"
)

var tokenCache string

func NewSequencer() *Sequencer {
	return &Sequencer{
		url: conf.GetConfig().UBI.SequencerUrl,
	}
}

func (s *Sequencer) getToken() error {
	var client = http.Client{
		Timeout: 30 * time.Second,
	}

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

	reqBody, err := json.Marshal(&data)
	if err != nil {
		return fmt.Errorf("failed to convet json, error: %v", err)
	}

	req, err := http.NewRequest("POST", s.url+token, strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	var returnResult ReturnResult
	if err = json.Unmarshal(body, &returnResult); err != nil {
		return err
	}
	if returnResult.Code == 0 {
		if dataMap, ok := returnResult.Data.(map[string]interface{}); ok {
			dataJSON, err := json.Marshal(dataMap)
			if err != nil {
				return err
			}
			var t Token
			err = json.Unmarshal(dataJSON, &t)
			if err != nil {
				return err
			}
			tokenCache = t.Token
		}
		return nil
	}
	return fmt.Errorf(returnResult.Msg)
}

func (s *Sequencer) SendTaskProof(data []byte) error {
	if tokenCache == "" {
		if err := s.getToken(); err != nil {
			return fmt.Errorf("failed to get token, error: %v", err)
		}
	}

	req, err := http.NewRequest("POST", s.url+task, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenCache)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	newToken := resp.Header.Get("new-token")
	if newToken != "" {
		tokenCache = newToken
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	logs.GetLogger().Infof("SendTaskProof: code: %d, result: %s", resp.StatusCode, string(body))

	var returnResult ReturnResult
	err = json.Unmarshal(body, &returnResult)
	if err != nil {
		return fmt.Errorf("failed to convert to json, error: %v", err)
	}

	if returnResult.Code != 0 {
		return fmt.Errorf(returnResult.Msg)
	}
	return nil
}

func (s *Sequencer) QueryTask(pageNo, pageSize int, taskId ...int) (TaskResp, error) {
	if tokenCache == "" {
		if err := s.getToken(); err != nil {
			return TaskResp{}, fmt.Errorf("failed to get token, error: %v", err)
		}
	}
	var url string
	if len(taskId) > 0 && taskId[0] != 0 {
		url = s.url + task + fmt.Sprintf("?id=%d&page_no=%d&page_size=%d", taskId[0], pageNo, pageSize)
	} else {
		url = s.url + task + fmt.Sprintf("?page_no=%d&page_size=%d", pageNo, pageSize)
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return TaskResp{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenCache)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TaskResp{}, fmt.Errorf("failed to sending request: %v", err)
	}
	defer resp.Body.Close()

	newToken := resp.Header.Get("new-token")
	if newToken != "" {
		tokenCache = newToken
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TaskResp{}, fmt.Errorf("error reading response: %v", err)
	}

	logs.GetLogger().Infof("QueryTask: code: %d, result: %s", resp.StatusCode, string(body))

	var returnResult ReturnResult
	err = json.Unmarshal(body, &returnResult)
	if err != nil {
		return TaskResp{}, fmt.Errorf("response convert to json failed: %v", err)
	}

	if returnResult.Code != 0 {
		return TaskResp{}, fmt.Errorf(returnResult.Msg)
	}
	return returnResult.Data.(TaskResp), nil
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

type ReturnResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

type TaskResp struct {
	Total int `json:"total"`
	List  []struct {
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
	} `json:"list"`
}
