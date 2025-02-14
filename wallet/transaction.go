package wallet

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func Balance(ctx context.Context, client *ethclient.Client, addr string) (string, error) {
	account := common.HexToAddress(addr)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", err
	}

	var ethValue string
	if balance.String() == "0" {
		ethValue = "0.0000"
	} else {
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		etherQuotient := new(big.Float).Quo(fbalance, new(big.Float).SetInt(big.NewInt(1e18)))
		ethValue = etherQuotient.Text('f', 4)
	}
	return ethValue, nil
}

func BalanceNumber(client *ethclient.Client, addr string) (float64, error) {
	account := common.HexToAddress(addr)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return 0, err
	}

	balanceEther := new(big.Float).SetInt(balance)
	weiPerEther := new(big.Float).SetFloat64(1e18) // 1 ether = 1e18 wei
	balanceEther.Quo(balanceEther, weiPerEther)
	balanceFloat64, _ := balanceEther.Float64()
	return balanceFloat64, nil
}

func sendTransaction(client *ethclient.Client, privateK string, to string, amount *big.Int, nonce uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateK)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if nonce == 0 {
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(to)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx.Hash().String(), nil
}
