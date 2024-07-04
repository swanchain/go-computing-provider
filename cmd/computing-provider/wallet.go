package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/swanchain/go-computing-provider/wallet"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var walletCmd = &cli.Command{
	Name:  "wallet",
	Usage: "Manage wallets",
	Subcommands: []*cli.Command{
		walletNew,
		walletList,
		walletExport,
		walletImport,
		walletDelete,
		walletSign,
		walletVerify,
		walletSend,
	},
}

var walletNew = &cli.Command{
	Name:  "new",
	Usage: "Generate a new key",
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		addr, err := localWallet.WalletNew(ctx)
		if err != nil {
			return err
		}
		fmt.Println(addr)

		return nil
	},
}

var walletList = &cli.Command{
	Name:  "list",
	Usage: "List wallet address",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "contract",
			Usage: "Specify the token contract",
			Value: false,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		contractFlag := cctx.Bool("contract")

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}

		return localWallet.WalletList(ctx, contractFlag)
	},
}

var walletExport = &cli.Command{
	Name:      "export",
	Usage:     "Export keys",
	ArgsUsage: "[address]",
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		if !cctx.Args().Present() {
			err := fmt.Errorf("must specify key to export")
			return err
		}

		addr := cctx.Args().First()
		if err != nil {
			return err
		}

		ki, err := localWallet.WalletExport(ctx, addr)
		if err != nil {
			return err
		}

		fmt.Println(ki.PrivateKey)
		return nil
	},
}

var walletImport = &cli.Command{
	Name:      "import",
	Usage:     "Import keys",
	ArgsUsage: "[<path> (optional, will read from stdin if omitted)]",
	Flags:     []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}

		var inpdata []byte
		if !cctx.Args().Present() || cctx.Args().First() == "-" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter private key: ")
			indata, err := reader.ReadBytes('\n')
			if err != nil {
				return err
			}
			inpdata = indata

		} else {
			fdata, err := os.ReadFile(cctx.Args().First())
			if err != nil {
				return err
			}
			inpdata = fdata
		}

		var ki wallet.KeyInfo

		ki.PrivateKey = strings.TrimSuffix(string(inpdata), "\n")

		addr, err := localWallet.WalletImport(ctx, &ki)
		if err != nil {
			return err
		}

		fmt.Printf("imported key %s successfully!\n", addr)
		return nil
	},
}

var walletDelete = &cli.Command{
	Name:      "delete",
	Usage:     "Delete an account from the wallet",
	ArgsUsage: "<address> ",
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}

		if !cctx.Args().Present() || cctx.NArg() != 1 {
			return fmt.Errorf("must specify address to delete")
		}

		addr := cctx.Args().First()
		return localWallet.WalletDelete(ctx, addr)
	},
}

var walletSign = &cli.Command{
	Name:      "sign",
	Usage:     "Sign a message",
	ArgsUsage: "<signing address> <Message>",
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}

		if !cctx.Args().Present() || cctx.NArg() != 2 {
			return fmt.Errorf("must specify signing address and message to sign")
		}

		addr := cctx.Args().First()
		if strings.TrimSpace(addr) == "" {
			return fmt.Errorf("failed to parse sign address")
		}

		msg := cctx.Args().Get(1)
		if strings.TrimSpace(msg) == "" {
			return fmt.Errorf("failed to parse message")
		}

		sig, err := localWallet.WalletSign(ctx, addr, []byte(msg))
		if err != nil {
			return err
		}
		fmt.Println(sig)
		return nil
	},
}

var walletVerify = &cli.Command{
	Name:      "verify",
	Usage:     "Verify the signature of a message",
	ArgsUsage: "<signing address>  <signature> <rawMessage>",
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)

		if cctx.NArg() != 3 {
			return fmt.Errorf("incorrect number of arguments, requires 3 parameters")
		}

		addr := cctx.Args().First()

		sigBytes, err := hexutil.Decode(cctx.Args().Get(1))
		if err != nil {
			return err
		}

		messageData := cctx.Args().Get(2)
		if strings.TrimSpace(messageData) == "" {
			return fmt.Errorf("failed to get raw message")
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}

		pass, err := localWallet.WalletVerify(ctx, addr, sigBytes, messageData)
		if err != nil {
			return err
		}
		fmt.Println(pass)
		return nil
	},
}

var walletSend = &cli.Command{
	Name:      "send",
	Usage:     "Send funds between accounts",
	ArgsUsage: "[targetAddress] [amount]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "from",
			Usage: "Optionally specify the account to send funds from",
		},
		&cli.Uint64Flag{
			Name:  "nonce",
			Usage: "optionally specify the nonce to use",
			Value: 0,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		if cctx.NArg() != 2 {
			return fmt.Errorf(" need two params: the target address and amount")
		}

		from := cctx.String("from")
		if strings.TrimSpace(from) == "" {
			return fmt.Errorf("failed to parse from address: %s", from)
		}

		to := cctx.Args().Get(0)
		if strings.TrimSpace(to) == "" {
			return fmt.Errorf("failed to parse target address: %s", to)
		}

		amount := cctx.Args().Get(1)
		if strings.TrimSpace(amount) == "" {
			return fmt.Errorf("failed to get amount: %s", amount)
		}
		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		txHash, err := localWallet.WalletSend(ctx, from, to, amount)
		if err != nil {
			return err
		}
		fmt.Println(txHash)
		return nil
	},
}

var collateralCmd = &cli.Command{
	Name:      "collateral",
	Usage:     "Manage the collateral amount",
	ArgsUsage: "[fromAddress] [amount]",
	Subcommands: []*cli.Command{
		collateralAddCmd,
		collateralSendCmd,
		collateralWithdrawCmd,
	},
}

var collateralAddCmd = &cli.Command{
	Name:  "add",
	Usage: "Send the collateral amount to the collateral contract address",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "fcp",
			Usage: "Specify the fcp collateral",
		},
		&cli.BoolFlag{
			Name:  "ecp",
			Usage: "Specify the ecp collateral",
		},
		&cli.BoolFlag{
			Name:  "sequencer",
			Usage: "Specify the sequencer collateral",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "Specify the wallet address, if the fcp is true, --form must specify the owner wallet address",
		},
		&cli.StringFlag{
			Name:  "account",
			Usage: "Specify the cp account address, if not specified, cp account is the content of the account file under the CP_PATH variable",
		},
	},
	ArgsUsage: "[amount]",
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)

		fcpCollateral := cctx.Bool("fcp")
		ecpCollateral := cctx.Bool("ecp")
		sequencerCollateral := cctx.Bool("sequencer")

		if !fcpCollateral && !ecpCollateral && !sequencerCollateral {
			return fmt.Errorf("must specify one of fcp or ecp or sequencer")
		}
		var collateralType string
		if fcpCollateral {
			collateralType = "fcp"
		}
		if ecpCollateral {
			collateralType = "ecp"
		}
		if sequencerCollateral {
			collateralType = "sequencer"
		}

		fromAddress := cctx.String("from")
		if strings.TrimSpace(fromAddress) == "" {
			return fmt.Errorf("the wallet address is required")
		}

		cpAccountAddress := cctx.String("account")

		amount := cctx.Args().Get(0)
		if strings.TrimSpace(amount) == "" {
			return fmt.Errorf("failed to get amount: %s", amount)
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		txHash, err := localWallet.WalletCollateral(ctx, fromAddress, amount, cpAccountAddress, collateralType)
		if err != nil {
			return err
		}
		fmt.Println(txHash)
		return nil
	},
}

var collateralWithdrawCmd = &cli.Command{
	Name:  "withdraw",
	Usage: "Withdraw funds from the collateral contract",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "fcp",
			Usage: "Specify the fcp collateral",
		},
		&cli.BoolFlag{
			Name:  "ecp",
			Usage: "Specify the ecp collateral",
		},
		&cli.BoolFlag{
			Name:  "sequencer",
			Usage: "Specify the sequencer collateral",
		},
		&cli.StringFlag{
			Name:  "owner",
			Usage: "Specify the owner address",
		},
		&cli.StringFlag{
			Name:  "account",
			Usage: "Specify the cp account address, if not specified, cp account is the content of the account file under the CP_PATH variable",
		},
	},
	ArgsUsage: "[amount]",
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)

		fcpCollateral := cctx.Bool("fcp")
		ecpCollateral := cctx.Bool("ecp")
		sequencerCollateral := cctx.Bool("sequencer")

		if !fcpCollateral && !ecpCollateral && !sequencerCollateral {
			return fmt.Errorf("must specify one of fcp or ecp or sequencer")
		}
		var collateralType string
		if fcpCollateral {
			collateralType = "fcp"
		}
		if ecpCollateral {
			collateralType = "ecp"
		}
		if sequencerCollateral {
			collateralType = "sequencer"
		}

		ownerAddress := cctx.String("owner")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("the owner address is required")
		}

		cpAccountAddress := cctx.String("account")
		amount := cctx.Args().Get(0)
		if strings.TrimSpace(amount) == "" {
			return fmt.Errorf("the amount param is required")
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		txHash, err := localWallet.CollateralWithdraw(ctx, ownerAddress, amount, cpAccountAddress, collateralType)
		if err != nil {
			return err
		}
		fmt.Println(txHash)
		return nil
	},
}

var collateralSendCmd = &cli.Command{
	Name:      "send",
	Usage:     "Send funds between accounts",
	ArgsUsage: "[targetAddress] [amount]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "from",
			Usage:    "Optionally specify the account to send funds from",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		if cctx.NArg() != 2 {
			return fmt.Errorf(" need two params: the target address and amount")
		}

		from := cctx.String("from")
		if strings.TrimSpace(from) == "" {
			return fmt.Errorf("the from address param is required")
		}

		to := cctx.Args().Get(0)
		if strings.TrimSpace(to) == "" {
			return fmt.Errorf("the to address param is required")
		}

		amount := cctx.Args().Get(1)
		if strings.TrimSpace(amount) == "" {
			return fmt.Errorf("the amount param is required")
		}
		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		txHash, err := localWallet.CollateralSend(ctx, from, to, amount)
		if err != nil {
			return err
		}
		fmt.Println(txHash)
		return nil
	},
}

func reqContext(cctx *cli.Context) context.Context {
	ctx, done := context.WithCancel(cctx.Context)
	sigChan := make(chan os.Signal, 2)
	go func() {
		<-sigChan
		done()
	}()
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	return ctx
}
