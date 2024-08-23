package wallet

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/internal/contract/ecp"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/contract/token"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/wallet/tablewriter"
	"golang.org/x/xerrors"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	WalletRepo  = "keystore"
	KNamePrefix = "wallet-"
)

var (
	ErrKeyInfoNotFound = fmt.Errorf("key info not found")
)

var reAddress = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func SetupWallet(dir string) (*LocalWallet, error) {
	cpPath, exit := os.LookupEnv("CP_PATH")
	if !exit {
		return nil, fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
	}

	var walletRepo = filepath.Join(cpPath, dir)
	var resultErr error
	timeOutCh := time.After(10 * time.Second)
loop:
	for {
		select {
		case <-timeOutCh:
			return nil, fmt.Errorf("open %s wallet repo timeout, retry again", walletRepo)
		default:
			kstore, err := OpenOrInitKeystore(walletRepo)
			if err != nil {
				if strings.Contains(err.Error(), "permission denied") {
					resultErr = err
					break loop
				}
				time.Sleep(time.Second)
				continue
			}
			return NewWallet(kstore), nil
		}
	}
	return nil, resultErr
}

type LocalWallet struct {
	*DiskKeyStore
}

func NewWallet(keystore *DiskKeyStore) *LocalWallet {
	w := &LocalWallet{
		keystore,
	}
	return w
}

func (w *LocalWallet) WalletSign(ctx context.Context, addr string, msg []byte) (string, error) {
	ki, err := w.FindKey(addr)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("signing using private key '%s': %w", addr, ErrKeyInfoNotFound)
	}
	signByte, err := Sign(ki.PrivateKey, msg)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(signByte), nil
}

func (w *LocalWallet) WalletVerify(ctx context.Context, addr string, sigByte []byte, data string) (bool, error) {
	hash := crypto.Keccak256Hash([]byte(data))
	return Verify(addr, sigByte, hash.Bytes())
}

func (w *LocalWallet) FindKey(addr string) (*KeyInfo, error) {
	defer w.Close()
	ki, err := w.Get(KNamePrefix + addr)
	return &ki, err
}

func (w *LocalWallet) WalletExport(ctx context.Context, addr string) (*KeyInfo, error) {
	k, err := w.FindKey(addr)
	if err != nil {
		return nil, xerrors.Errorf("failed to find key to export: %w", err)
	}
	if k == nil {
		return nil, xerrors.Errorf("private key not found for %s", addr)
	}

	return k, nil
}

func (w *LocalWallet) WalletImport(ctx context.Context, ki *KeyInfo) (string, error) {
	defer w.Close()
	if ki == nil || len(strings.TrimSpace(ki.PrivateKey)) == 0 {
		return "", fmt.Errorf("not found private key")
	}

	_, publicKeyECDSA, err := ToPublic(ki.PrivateKey)
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	existAddress, err := w.Get(KNamePrefix + address)
	if err == nil && existAddress.PrivateKey != "" {
		return "", xerrors.Errorf("This wallet address already exists")
	}

	if err := w.Put(KNamePrefix+address, *ki); err != nil {
		return "", xerrors.Errorf("saving to keystore: %w", err)
	}
	return "", nil
}

func (w *LocalWallet) WalletList(ctx context.Context, contractFlag bool) error {
	addressList, err := w.addressList(ctx)
	if err != nil {
		return err
	}

	addressKey := "Address"
	balanceKey := "Balance"
	nonceKey := "Nonce"
	errorKey := "Error"

	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(chainRpc)
	if err != nil {
		return err
	}
	defer client.Close()

	var wallets []map[string]interface{}
	for _, addr := range addressList {
		var balance string
		if contractFlag {
			tokenStub, err := token.NewTokenStub(client, token.WithPublicKey(addr))
			if err == nil {
				balance, err = tokenStub.BalanceOf()
			}
		} else {
			balance, err = Balance(ctx, client, addr)
		}

		var errmsg string
		if err != nil {
			errmsg = err.Error()
		}

		nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(addr))
		if err != nil {
			errmsg = err.Error()
		}

		wallet := map[string]interface{}{
			addressKey: addr,
			balanceKey: balance,
			errorKey:   errmsg,
			nonceKey:   nonce,
		}
		wallets = append(wallets, wallet)
	}

	tw := tablewriter.New(
		tablewriter.Col(addressKey),
		tablewriter.Col(balanceKey),
		tablewriter.Col(nonceKey),
		tablewriter.NewLineCol(errorKey))

	for _, wallet := range wallets {
		tw.Write(wallet)
	}
	return tw.Flush(os.Stdout)
}

func (w *LocalWallet) WalletNew(ctx context.Context) (string, error) {
	defer w.Close()

	privateK, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateK)
	privateKey := hexutil.Encode(privateKeyBytes)[2:]

	_, publicKeyECDSA, err := ToPublic(privateKey)
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	keyInfo := KeyInfo{PrivateKey: privateKey}
	if err := w.Put(KNamePrefix+address, keyInfo); err != nil {
		return "", xerrors.Errorf("saving to keystore: %w", err)
	}
	return address, nil
}

func (w *LocalWallet) WalletDelete(ctx context.Context, addr string) error {
	defer w.Close()
	if err := w.Delete(KNamePrefix + addr); err != nil {
		return xerrors.Errorf("failed to delete key %s: %w", addr, err)
	}

	fmt.Printf("%s has been deleted from the local success \n", addr)
	return nil
}

func (w *LocalWallet) WalletSend(ctx context.Context, from, to string, amount string) (string, error) {
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}
	ki, err := w.FindKey(from)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private %w,", from, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	sendAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	txHash, err := sendTransaction(client, ki.PrivateKey, to, sendAmount)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

func (w *LocalWallet) WalletCollateral(ctx context.Context, from string, amount string, cpAccountAddress string, collateralType string) (string, error) {
	sendAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}
	ki, err := w.FindKey(from)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", from, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if len(cpAccountAddress) > 0 {
		cpAccount := common.HexToAddress(cpAccountAddress)
		bytecode, err := client.CodeAt(context.Background(), cpAccount, nil)
		if err != nil {
			return "", fmt.Errorf("failed to check cp account contract address, error: %v", err)
		}

		if len(bytecode) <= 0 {
			return "", fmt.Errorf("the account parameter must be a cpAccount contract address")
		}

		cpStub, err := account.NewAccountStub(client, account.WithContractAddress(cpAccountAddress))
		if err != nil {
			return "", err
		}
		if _, err = cpStub.GetCpAccountInfo(); err != nil {
			return "", fmt.Errorf("cp account: %s does not exist on the chain", cpAccountAddress)
		}
	}

	if collateralType == "fcp" {
		tokenStub, err := token.NewTokenStub(client, token.WithCollateralContract(conf.GetConfig().CONTRACT.JobCollateral), token.WithPrivateKey(ki.PrivateKey))
		if err != nil {
			return "", err
		}
		swanTokenTxHash, err := tokenStub.Approve(sendAmount)
		if err != nil {
			return "", err
		}

		timeout := time.After(3 * time.Minute)
		ticker := time.Tick(3 * time.Second)
		for {
			select {
			case <-timeout:
				return "", fmt.Errorf("timeout waiting for transaction confirmation, tx: %s", swanTokenTxHash)
			case <-ticker:
				receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(swanTokenTxHash))
				if err != nil {
					if errors.Is(err, ethereum.NotFound) {
						continue
					}
					return "", fmt.Errorf("check swan token Approve tx, error: %+v", err)
				}
				if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
					fmt.Printf("swan token approve TX Hash: %s \n", swanTokenTxHash)
					collateralStub, err := fcp.NewCollateralStub(client, fcp.WithPrivateKey(ki.PrivateKey), fcp.WithCpAccountAddress(cpAccountAddress))
					if err != nil {
						return "", err
					}
					collateralTxHash, err := collateralStub.Deposit(sendAmount)
					if err != nil {
						return "", err
					}
					return collateralTxHash, nil

				} else if receipt != nil && receipt.Status == 0 {
					return "", fmt.Errorf("failed to check swan token approve transaction, tx: %s", swanTokenTxHash)
				}
			}
		}
	} else if collateralType == "ecp" {
		tokenStub, err := token.NewTokenStub(client, token.WithCollateralContract(conf.GetConfig().CONTRACT.ZkCollateral), token.WithPrivateKey(ki.PrivateKey))
		if err != nil {
			return "", err
		}

		swanTokenTxHash, err := tokenStub.Approve(sendAmount)
		if err != nil {
			return "", err
		}

		timeout := time.After(3 * time.Minute)
		ticker := time.Tick(3 * time.Second)
		for {
			select {
			case <-timeout:
				return "", fmt.Errorf("timeout waiting for transaction confirmation, tx: %s", swanTokenTxHash)
			case <-ticker:
				receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(swanTokenTxHash))
				if err != nil {
					if errors.Is(err, ethereum.NotFound) {
						continue
					}
					return "", fmt.Errorf("check swan token Approve tx, error: %+v", err)
				}

				if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
					fmt.Printf("swan token approve TX Hash: %s \n", swanTokenTxHash)
					cpStub, err := account.NewAccountStub(client, account.WithContractAddress(cpAccountAddress))
					if err != nil {
						return "", err
					}
					if _, err = cpStub.GetCpAccountInfo(); err != nil {
						return "", fmt.Errorf("cp account: %s does not exist on the chain", cpAccountAddress)
					}
					zkCollateral, err := ecp.NewCollateralStub(client, ecp.WithPrivateKey(ki.PrivateKey), ecp.WithCpAccountAddress(cpAccountAddress))
					if err != nil {
						return "", err
					}

					collateralTxHash, err := zkCollateral.Deposit(sendAmount)
					if err != nil {
						return "", err
					}
					return collateralTxHash, nil
				} else if receipt != nil && receipt.Status == 0 {
					return "", fmt.Errorf("failed to check swan token approve transaction, tx: %s", swanTokenTxHash)
				}
			}
		}
	}
	return "", nil
}

func (w *LocalWallet) CollateralWithdraw(ctx context.Context, address string, amount string, cpAccountAddress string, collateralType string) (string, error) {
	withDrawAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}

	ki, err := w.FindKey(address)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", address, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if len(cpAccountAddress) > 0 {
		cpAccount := common.HexToAddress(cpAccountAddress)
		bytecode, err := client.CodeAt(context.Background(), cpAccount, nil)
		if err != nil {
			return "", fmt.Errorf("check cp account contract address failed, error: %v", err)
		}

		if len(bytecode) <= 0 {
			return "", fmt.Errorf("the account parameter must be a cpAccount contract address")
		}
	}

	if collateralType == "fcp" {
		collateralStub, err := fcp.NewCollateralStub(client, fcp.WithPrivateKey(ki.PrivateKey), fcp.WithCpAccountAddress(cpAccountAddress))
		if err != nil {
			return "", err
		}
		return collateralStub.Withdraw(withDrawAmount)
	} else if collateralType == "ecp" {
		zkCollateral, err := ecp.NewCollateralStub(client, ecp.WithPrivateKey(ki.PrivateKey), ecp.WithCpAccountAddress(cpAccountAddress))
		if err != nil {
			return "", err
		}
		return zkCollateral.Withdraw(withDrawAmount)
	} else {
		sequencerStub, err := ecp.NewSequencerStub(client, ecp.WithSequencerPrivateKey(ki.PrivateKey), ecp.WithSequencerCpAccountAddress(cpAccountAddress))
		if err != nil {
			return "", err
		}
		return sequencerStub.Withdraw(withDrawAmount)
	}
}

func (w *LocalWallet) SequencerDeposit(ctx context.Context, from string, amount string, cpAccountAddress string) (string, error) {
	sendAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}
	ki, err := w.FindKey(from)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", from, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if len(cpAccountAddress) > 0 {
		cpAccount := common.HexToAddress(cpAccountAddress)
		bytecode, err := client.CodeAt(context.Background(), cpAccount, nil)
		if err != nil {
			return "", fmt.Errorf("failed to check cp account contract address, error: %v", err)
		}

		if len(bytecode) <= 0 {
			return "", fmt.Errorf("the account parameter must be a cpAccount contract address")
		}

		cpStub, err := account.NewAccountStub(client, account.WithContractAddress(cpAccountAddress))
		if err != nil {
			return "", err
		}
		if _, err = cpStub.GetCpAccountInfo(); err != nil {
			return "", fmt.Errorf("cp account: %s does not exist on the chain", cpAccountAddress)
		}
	}

	sequencerStub, err := ecp.NewSequencerStub(client, ecp.WithSequencerPrivateKey(ki.PrivateKey), ecp.WithSequencerCpAccountAddress(cpAccountAddress))
	if err != nil {
		return "", err
	}
	return sequencerStub.Deposit(sendAmount)
}

func (w *LocalWallet) SequencerWithdraw(ctx context.Context, address string, amount string, cpAccountAddress string) (string, error) {
	withDrawAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}

	ki, err := w.FindKey(address)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", address, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if len(cpAccountAddress) > 0 {
		cpAccount := common.HexToAddress(cpAccountAddress)
		bytecode, err := client.CodeAt(context.Background(), cpAccount, nil)
		if err != nil {
			return "", fmt.Errorf("check cp account contract address failed, error: %v", err)
		}

		if len(bytecode) <= 0 {
			return "", fmt.Errorf("the account parameter must be a cpAccount contract address")
		}
	}

	sequencerStub, err := ecp.NewSequencerStub(client, ecp.WithSequencerPrivateKey(ki.PrivateKey), ecp.WithSequencerCpAccountAddress(cpAccountAddress))
	if err != nil {
		return "", err
	}
	return sequencerStub.Withdraw(withDrawAmount)
}

func (w *LocalWallet) CollateralWithdrawRequest(ctx context.Context, address string, amount string, cpAccountAddress string) (string, error) {
	withDrawAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}

	ki, err := w.FindKey(address)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", address, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if len(cpAccountAddress) > 0 {
		cpAccount := common.HexToAddress(cpAccountAddress)
		bytecode, err := client.CodeAt(context.Background(), cpAccount, nil)
		if err != nil {
			return "", fmt.Errorf("check cp account contract address failed, error: %v", err)
		}

		if len(bytecode) <= 0 {
			return "", fmt.Errorf("the account parameter must be a cpAccount contract address")
		}
	}

	zkCollateral, err := ecp.NewCollateralStub(client, ecp.WithPrivateKey(ki.PrivateKey), ecp.WithCpAccountAddress(cpAccountAddress))
	if err != nil {
		return "", err
	}
	return zkCollateral.WithdrawRequest(withDrawAmount)
}

func (w *LocalWallet) CollateralWithdrawConfirm(ctx context.Context, address string, cpAccountAddress string) (string, error) {
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}

	ki, err := w.FindKey(address)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", address, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if len(cpAccountAddress) > 0 {
		cpAccount := common.HexToAddress(cpAccountAddress)
		bytecode, err := client.CodeAt(context.Background(), cpAccount, nil)
		if err != nil {
			return "", fmt.Errorf("failed to check cp account contract address, error: %v", err)
		}

		if len(bytecode) <= 0 {
			return "", fmt.Errorf("the account parameter must be a cpAccount contract address")
		}
	}

	zkCollateral, err := ecp.NewCollateralStub(client, ecp.WithPrivateKey(ki.PrivateKey), ecp.WithCpAccountAddress(cpAccountAddress))
	if err != nil {
		return "", err
	}
	return zkCollateral.WithdrawConfirm()
}

func (w *LocalWallet) CollateralWithdrawView(ctx context.Context, cpAccountAddress string) (models.WithdrawRequest, error) {
	var withdrawRequest models.WithdrawRequest
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return withdrawRequest, err
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return withdrawRequest, err
	}
	defer client.Close()

	if len(cpAccountAddress) > 0 {
		cpAccount := common.HexToAddress(cpAccountAddress)
		bytecode, err := client.CodeAt(context.Background(), cpAccount, nil)
		if err != nil {
			return withdrawRequest, fmt.Errorf("check cp account contract address failed, error: %v", err)
		}

		if len(bytecode) <= 0 {
			return withdrawRequest, fmt.Errorf("the account parameter must be a cpAccount contract address")
		}
	}

	zkCollateral, err := ecp.NewCollateralStub(client, ecp.WithCpAccountAddress(cpAccountAddress))
	if err != nil {
		return withdrawRequest, err
	}

	contractInfo, err := zkCollateral.ContractInfo()
	if err != nil {
		return withdrawRequest, err
	}

	withdrawRequest, err = zkCollateral.WithdrawView()
	if err != nil {
		return withdrawRequest, err
	}
	withdrawRequest.WithdrawDelay = contractInfo.WithdrawDelay
	return withdrawRequest, nil
}

func (w *LocalWallet) CollateralSend(ctx context.Context, from, to string, amount string) (string, error) {
	withDrawAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}

	ki, err := w.FindKey(from)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", to, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	collateralStub, err := token.NewTokenStub(client, token.WithPrivateKey(ki.PrivateKey))
	if err != nil {
		return "", err
	}
	withdrawHash, err := collateralStub.Transfer(to, withDrawAmount)
	if err != nil {
		return "", err
	}

	return withdrawHash, nil
}

func (w *LocalWallet) addressList(ctx context.Context) ([]string, error) {
	defer w.Close()
	all, err := w.List()
	if err != nil {
		return nil, xerrors.Errorf("listing keystore: %w", err)
	}

	addressList := make([]string, 0, len(all))
	for _, a := range all {
		if strings.HasPrefix(a, KNamePrefix) {
			addr := strings.TrimPrefix(a, KNamePrefix)
			addressList = append(addressList, addr)
		}
	}
	return addressList, nil
}

func convertToWei(ethValue string) (*big.Int, error) {
	ethFloat, ok := new(big.Float).SetString(ethValue)
	if !ok {
		return nil, fmt.Errorf("conversion to float failed")
	}
	weiConversion := new(big.Float).SetFloat64(1e18)
	weiFloat := new(big.Float).Mul(ethFloat, weiConversion)
	weiInt, acc := new(big.Int).SetString(weiFloat.Text('f', 0), 10)
	if !acc {
		return nil, fmt.Errorf("conversion to Wei failed")
	}
	return weiInt, nil
}
