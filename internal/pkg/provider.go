package pkg

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/swanchain/go-computing-provider/internal/v1/models"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
)

func Reconnect(nodeID string) string {
	updateProviderInfo(nodeID, "", "", models.ActiveStatus)
	return nodeID
}

func updateProviderInfo(nodeID, peerID, address string, status string) {
	updateURL := conf.GetConfig().HUB.ServerUrl + "/cp"

	var cpName string
	if conf.GetConfig().API.NodeName != "" {
		cpName = conf.GetConfig().API.NodeName
	} else {
		cpName, _ = os.Hostname()
	}

	provider := models.ComputingProvider{
		PublicAddress: conf.GetConfig().HUB.WalletAddress,
		Name:          cpName,
		NodeId:        nodeID,
		MultiAddress:  conf.GetConfig().API.MultiAddress,
		Autobid:       1,
		Status:        status,
	}

	jsonData, err := json.Marshal(provider)
	if err != nil {
		logs.GetLogger().Errorf("Error marshaling provider data: %v", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", updateURL, bytes.NewBuffer(jsonData))
	if err != nil {
		logs.GetLogger().Errorf("Error creating request: %v", err)
		return
	}

	// Set the content type and API token in the request header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		logs.GetLogger().Errorf("Error updating provider info: %v", err)
	} else {
		if resp.StatusCode == 400 {
			respData, _ := io.ReadAll(resp.Body)
			logs.GetLogger().Info(string(respData))
		}

		err := resp.Body.Close()
		if err != nil {
			logs.GetLogger().Errorf(err.Error())
			return
		}
	}
}

func InitComputingProvider(cpRepoPath string) string {
	nodeID, peerID, address := GenerateNodeID(cpRepoPath)

	logs.GetLogger().Infof("Node ID :%s Peer ID:%s address:%s",
		nodeID,
		peerID, address)
	updateProviderInfo(nodeID, peerID, address, models.ActiveStatus)
	return nodeID
}

func GetNodeId(cpRepoPath string) string {
	nodeID, _, _ := GenerateNodeID(cpRepoPath)
	return nodeID
}

func GenerateNodeID(cpRepoPath string) (string, string, string) {
	privateKeyPath := filepath.Join(cpRepoPath, "private_key")
	var privateKeyBytes []byte

	if _, err := os.Stat(privateKeyPath); err == nil {
		privateKeyBytes, err = os.ReadFile(privateKeyPath)
		if err != nil {
			log.Fatalf("Error reading private key: %v", err)
		}
	} else {
		privateKeyBytes = make([]byte, 32)
		_, err := rand.Read(privateKeyBytes)
		if err != nil {
			log.Fatalf("Error generating random key: %v", err)
		}

		err = os.MkdirAll(filepath.Dir(privateKeyPath), os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating directory for private key: %v", err)
		}

		err = os.WriteFile(privateKeyPath, privateKeyBytes, 0644)
		if err != nil {
			log.Fatalf("Error writing private key: %v", err)
		}
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Error converting private key bytes: %v", err)
	}
	nodeID := hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey))
	peerID := hashPublicKey(&privateKey.PublicKey)
	address := crypto.PubkeyToAddress(privateKey.PublicKey).String()
	return nodeID, peerID, address
}

func hashPublicKey(publicKey *ecdsa.PublicKey) string {
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	hash := sha256.Sum256(publicKeyBytes)
	return hex.EncodeToString(hash[:])
}
