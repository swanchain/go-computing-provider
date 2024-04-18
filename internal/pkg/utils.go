package pkg

import (
	"context"
	"encoding/json"
	logging "github.com/ipfs/go-log/v2"
	"io"
	"k8s.io/client-go/rest"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var ulog = logging.Logger("utils")

func GenerateString(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	source := characters + numbers
	result := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result[i] = source[rand.Intn(len(source))]
	}
	return string(result)
}

func GetLocation() (string, error) {
	publicIpAddress, err := getLocalIPAddress()
	if err != nil {
		return "", err
	}
	ulog.Infof("the cp publicIpAddress: %s", publicIpAddress)

	resp, err := http.Get("http://ip-api.com/json/" + publicIpAddress)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipInfo struct {
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		City        string `json:"city"`
		Region      string `json:"region"`
		RegionName  string `json:"regionName"`
	}
	if err = json.Unmarshal(body, &ipInfo); err != nil {
		return "", err
	}

	return strings.TrimSpace(ipInfo.RegionName) + "-" + ipInfo.CountryCode, nil
}

func getLocalIPAddress() (string, error) {
	req, err := http.NewRequest("GET", "https://ipapi.co/ip", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(ipBytes)), nil
}

func readLog(req *rest.Request) (*strings.Builder, error) {
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return nil, err
	}
	defer podLogs.Close()
	buf := new(strings.Builder)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func GenerateLabel(name string) map[string]string {
	if name != "" {
		return map[string]string{
			name: "true",
		}
	} else {
		return map[string]string{}
	}
}
