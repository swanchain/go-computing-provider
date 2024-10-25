package contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func BalanceToStr(balance *big.Int) string {
	var ethValue string
	if balance.String() == "0" {
		ethValue = "0.0000"
	} else {
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		etherQuotient := new(big.Float).Quo(fbalance, new(big.Float).SetInt(big.NewInt(1e18)))
		ethValue = etherQuotient.Text('f', 4)
	}
	return ethValue
}

func BalanceToStr2(balance *big.Int) string {
	var ethValue string
	if balance.String() == "0" {
		ethValue = "0.000000"
	} else {
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		etherQuotient := new(big.Float).Quo(fbalance, new(big.Float).SetInt(big.NewInt(1e18)))
		ethValue = etherQuotient.Text('f', 6)
	}
	return ethValue
}

func GetCpAccountAddress() (string, error) {
	cpPath, exit := os.LookupEnv("CP_PATH")
	if !exit {
		return "", fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
	}

	accountFileName := filepath.Join(cpPath, "account")
	if _, err := os.Stat(accountFileName); err != nil {
		return "", fmt.Errorf("CP Account is empty. Please create an account first")
	}

	accountAddress, err := os.ReadFile(accountFileName)
	if err != nil {
		return "", fmt.Errorf("failed to get cp account contract address, error: %v", err)
	}

	return string(accountAddress), err
}

func GetEthClient(rpcUrl string) (*ethclient.Client, error) {
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	rpcClient, err := rpc.DialOptions(context.TODO(), rpcUrl, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the rpc client: %v", err)
	}
	return ethclient.NewClient(rpcClient), nil

}
