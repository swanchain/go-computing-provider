package computing

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/swanchain/go-computing-provider/internal/models"
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

    // Verify that required fields are not empty before sending them
    if nodeID == "" || conf.GetConfig().HUB.WalletAddress == "" || status == "" {
        logs.GetLogger().Error("Required fields are missing: ensure NodeID, WalletAddress, and Status are provided")
        return
    }

    provider := models.ComputingProvider{
        PublicAddress: conf.GetConfig().HUB.WalletAddress,
        Name:          cpName,
        NodeId:        nodeID,
        MultiAddress:  conf.GetConfig().API.MultiAddress,
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

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + conf.GetConfig().HUB.AccessToken)

    resp, err := client.Do(req)
    if err != nil {
        logs.GetLogger().Errorf("Error sending request to update provider info: %v", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        bodyBytes, _ := io.ReadAll(resp.Body)
        responseBody := string(bodyBytes)
        logs.GetLogger().Errorf("Failed to update provider info, status code: %d, response: %s", resp.StatusCode, responseBody)
    } else {
        bodyBytes, _ := io.ReadAll(resp.Body)
        responseBody := string(bodyBytes)
        logs.GetLogger().Infof("Provider info updated successfully, response: %s", responseBody)
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
