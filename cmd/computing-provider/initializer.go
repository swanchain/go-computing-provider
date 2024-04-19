package main

import (
	"fmt"
	"github.com/swanchain/go-computing-provider/internal"
	"github.com/swanchain/go-computing-provider/internal/pkg"
	computing2 "github.com/swanchain/go-computing-provider/internal/v1/service"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/swanchain/go-computing-provider/conf"
)

func sendHeartbeat(nodeId string) {
	// Replace the following URL with your Flask application's heartbeat endpoint URL
	heartbeatURL := conf.GetConfig().HUB.ServerUrl + "/cp/heartbeat"
	payload := strings.NewReader(fmt.Sprintf(`{
	"public_address": "%s",
    "node_id": "%s",
    "status": "Active"
}`, conf.GetConfig().HUB.WalletAddress, nodeId))

	client := &http.Client{}
	req, err := http.NewRequest("POST", heartbeatURL, payload)
	if err != nil {
		mlog.Errorf("Error creating request: %v", err)
		return
	}
	// Set the API token in the request header (replace "your_api_token" with the actual token)
	req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		mlog.Errorf("Error sending heartbeat, retrying to connect to the Swan Hub server: %v", err)
		pkg.Reconnect(nodeId)
	} else {
		_, err := ioutil.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			mlog.Warnf("resp status: %d, retrying to connect to the Swan Hub server", resp.StatusCode)
			pkg.Reconnect(nodeId)
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func sendHeartbeats(nodeId string) {
	ticker := time.NewTicker(15 * time.Second)
	for range ticker.C {
		sendHeartbeat(nodeId)
	}
}
func ProjectInit(cpRepoPath string) {
	if err := conf.InitConfig(cpRepoPath); err != nil {
		mlog.Fatal(err)
	}
	nodeID := pkg.InitComputingProvider(cpRepoPath)
	// Start sending heartbeats
	go sendHeartbeats(nodeID)

	go computing2.NewScheduleTask().Run()

	computing2.NewCronTask().RunTask()
	computing2.RunSyncTask(nodeID)
	celeryService := pkg.NewCeleryService()
	celeryService.RegisterTask(internal.TASK_DEPLOY, computing2.DeploySpaceTask)
	celeryService.Start()

}
