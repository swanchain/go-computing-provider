package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/account"
	"github.com/swanchain/go-computing-provider/wallet"
	"github.com/urfave/cli/v2"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var accountCmd = &cli.Command{
	Name:  "account",
	Usage: "Manage account info of CP",
	Subcommands: []*cli.Command{
		changeMultiAddressCmd,
		changeOwnerAddressCmd,
		changeBeneficiaryAddressCmd,
		changeUbiFlagCmd,
	},
}

var changeMultiAddressCmd = &cli.Command{
	Name:      "changeMultiAddress",
	Usage:     "Update MultiAddress of CP (/ip4/<public_ip>/tcp/<port>)",
	ArgsUsage: "[multiAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a multiAddress")
		}

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		multiAddr := cctx.Args().Get(0)
		if strings.TrimSpace(multiAddr) == "" {
			return fmt.Errorf("failed to parse : %s", multiAddr)
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			mlog.Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			mlog.Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			mlog.Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			mlog.Errorf("the address: %s, private key %v,", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			mlog.Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			mlog.Errorf("create ubi task client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		submitUBIProofTx, err := cpStub.ChangeMultiAddress([]string{multiAddr})
		if err != nil {
			mlog.Errorf("change multi-addr tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeMultiAddress: %s", submitUBIProofTx)

		return nil
	},
}

var changeOwnerAddressCmd = &cli.Command{
	Name:      "changeOwnerAddress",
	Usage:     "Update OwnerAddress of CP",
	ArgsUsage: "[newOwnerAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("oldOwnerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a ownerAddress")
		}

		newOwnerAddr := cctx.Args().Get(0)
		if strings.TrimSpace(newOwnerAddr) == "" {
			return fmt.Errorf("failed to parse : %s", newOwnerAddr)
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			mlog.Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			mlog.Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			mlog.Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			mlog.Errorf("the old owner address: %s, private key %v,", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			mlog.Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			mlog.Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		submitUBIProofTx, err := cpStub.ChangeOwnerAddress(common.HexToAddress(newOwnerAddr))
		if err != nil {
			mlog.Errorf("change owner address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeOwnerAddress: %s", submitUBIProofTx)

		return nil
	},
}

var changeBeneficiaryAddressCmd = &cli.Command{
	Name:      "changeBeneficiaryAddress",
	Usage:     "Update beneficiaryAddress of CP",
	ArgsUsage: "[beneficiaryAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a beneficiaryAddress")
		}

		beneficiaryAddress := cctx.Args().Get(0)
		if strings.TrimSpace(beneficiaryAddress) == "" {
			return fmt.Errorf("failed to parse target address: %s", beneficiaryAddress)
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			mlog.Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			mlog.Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			mlog.Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			mlog.Errorf("the address: %s, private key %v. Please import the address into the wallet", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			mlog.Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			mlog.Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}
		newQuota := big.NewInt(int64(0))
		newExpiration := big.NewInt(int64(0))
		changeBeneficiaryAddressTx, err := cpStub.ChangeBeneficiary(common.HexToAddress(beneficiaryAddress), newQuota, newExpiration)
		if err != nil {
			mlog.Errorf("change owner address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("changeBeneficiaryAddress Transaction hash: %s", changeBeneficiaryAddressTx)
		return nil
	},
}

var changeUbiFlagCmd = &cli.Command{
	Name:      "changeUbiFlag",
	Usage:     "Update ubiFlag of CP (0:Reject, 1:Accept)",
	ArgsUsage: "[ubiFlag]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a beneficiaryAddress")
		}

		ubiFlag := cctx.Args().Get(0)
		if strings.TrimSpace(ubiFlag) == "" {
			return fmt.Errorf("ubiFlag is not empty")
		}

		if strings.TrimSpace(ubiFlag) != "0" && strings.TrimSpace(ubiFlag) != "1" {
			return fmt.Errorf("ubiFlag must be 0 or 1")
		}

		cpRepoPath := cctx.String(FlagCpRepo)
		os.Setenv("CP_PATH", cpRepoPath)
		if err := conf.InitConfig(cpRepoPath); err != nil {
			mlog.Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			mlog.Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			mlog.Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			mlog.Errorf("the address: %s, private key %v. Please import the address into the wallet", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			mlog.Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			mlog.Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		newUbiFlag, _ := strconv.ParseUint(strings.TrimSpace(ubiFlag), 10, 64)

		changeBeneficiaryAddressTx, err := cpStub.ChangeUbiFlag(uint8(newUbiFlag))
		if err != nil {
			mlog.Errorf("change ubi flag tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeUbiFlag Transaction hash: %s", changeBeneficiaryAddressTx)
		return nil
	},
}
