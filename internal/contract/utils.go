package contract

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

func BalanceToStr(balance *big.Int) string {
	var ethValue string
	if balance.String() == "0" {
		ethValue = "0.000"
	} else {
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		etherQuotient := new(big.Float).Quo(fbalance, new(big.Float).SetInt(big.NewInt(1e18)))
		ethValue = etherQuotient.Text('f', 3)
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
		return "", fmt.Errorf("please use the account create command to initialize the account of CP")
	}

	accountAddress, err := os.ReadFile(filepath.Join(cpPath, "account"))
	if err != nil {
		return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
	}

	return string(accountAddress), err
}
