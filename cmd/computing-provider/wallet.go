package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/wallet"
	"github.com/urfave/cli/v2"
	"math/big"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
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
	Before: func(c *cli.Context) error {
		if c.Args().Present() {
			if strings.EqualFold(c.Args().First(), walletList.Name) || strings.EqualFold(c.Args().First(), walletSend.Name) {
				cpRepoPath, _ := os.LookupEnv("CP_PATH")
				if err := conf.InitConfig(cpRepoPath, true); err != nil {
					return err
				}
			}
		}
		return nil
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
			Name:  "swan",
			Usage: "Specify the swan token contract",
			Value: false,
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		swanContractFlag := cctx.Bool("swan")

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}

		return localWallet.WalletList(ctx, swanContractFlag)
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
		fmt.Println("Signature:")
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
		collateralWithDrawRequestCmd,
		collateralWithDrawConfirmCmd,
		collateralWithDrawViewCmd,
	},
	Before: func(c *cli.Context) error {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return err
		}
		return nil
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

		if !fcpCollateral && !ecpCollateral {
			return fmt.Errorf("must specify one of fcp or ecp or sequencer")
		}
		var collateralType string
		if fcpCollateral {
			collateralType = "fcp"
		}
		if ecpCollateral {
			collateralType = "ecp"
		}
		if collateralType == "" {
			return fmt.Errorf("not suport collateral type")
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

		fmt.Printf("collateral TX Hash: %s \n", txHash)
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

		if !fcpCollateral && !ecpCollateral {
			return fmt.Errorf("must specify one of fcp or ecp")
		}
		var collateralType string
		if fcpCollateral {
			collateralType = "fcp"
		}
		if ecpCollateral {
			collateralType = "ecp"
		}

		if collateralType == "" {
			return fmt.Errorf("not suport withdraw type")
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

var collateralWithDrawRequestCmd = &cli.Command{
	Name:  "withdraw-request",
	Usage: "Send a request to withdraw tokens from the collateral escrow account from the collateral contract",
	Flags: []cli.Flag{
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
		txHash, err := localWallet.CollateralWithdrawRequest(ctx, ownerAddress, amount, cpAccountAddress)
		if err != nil {
			return err
		}
		fmt.Println(txHash)
		return nil
	},
}

var collateralWithDrawConfirmCmd = &cli.Command{
	Name:  "withdraw-confirm",
	Usage: "Confirm a request to withdraw tokens from the collateral escrow account from the collateral contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "owner",
			Usage: "Specify the owner address",
		},
		&cli.StringFlag{
			Name:  "account",
			Usage: "Specify the cp account address, if not specified, cp account is the content of the account file under the CP_PATH variable",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)

		ownerAddress := cctx.String("owner")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("the owner address is required")
		}

		cpAccountAddress := cctx.String("account")

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		txHash, err := localWallet.CollateralWithdrawConfirm(ctx, ownerAddress, cpAccountAddress)
		if err != nil {
			return err
		}
		fmt.Println(txHash)
		return nil
	},
}

var collateralWithDrawViewCmd = &cli.Command{
	Name:  "withdraw-view",
	Usage: "View a request to withdraw tokens from the collateral escrow account from the collateral contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "account",
			Usage: "Specify the cp account address, if not specified, cp account is the content of the account file under the CP_PATH variable",
		},
	},
	Action: func(cctx *cli.Context) error {
		ctx := reqContext(cctx)
		cpAccountAddress := cctx.String("account")

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			return err
		}
		withdrawView, err := localWallet.CollateralWithdrawView(ctx, cpAccountAddress)
		if err != nil {
			return err
		}

		chainRpc, err := conf.GetRpcByNetWorkName()
		if err != nil {
			return err
		}
		client, err := contract.GetEthClient(chainRpc)
		if err != nil {
			return err
		}
		defer client.Close()
		latestBlockNumber, _ := client.BlockNumber(ctx)

		var amount = "0.0000"
		var requestBlock int64
		var confirmableBlock int64
		var confirmableBlockStr string
		if withdrawView.RequestBlock != 0 {
			amount = withdrawView.Amount
			requestBlock = withdrawView.RequestBlock
			confirmableBlock = withdrawView.RequestBlock + withdrawView.WithdrawDelay

			var secondFlag int64 = 2
			chainId, err := client.ChainID(context.Background())
			if err != nil {
				return fmt.Errorf("failed to get chain id, error: %v", err)
			}
			if chainId.Int64() == 254 {
				secondFlag = 5
			}

			block, err := client.HeaderByNumber(context.Background(), big.NewInt(withdrawView.RequestBlock))
			if err != nil {
				return fmt.Errorf("failed to get block time, block: %d, error: %v", withdrawView.RequestBlock, err)
			}

			currentTime := time.Unix(int64(block.Time)+withdrawView.WithdrawDelay*secondFlag, 0)
			timeStr := currentTime.Format("2006-01-02T15:04:05")
			timeZone, _ := currentTime.Zone()
			confirmableBlockStr = fmt.Sprintf("%d\u200B(%s+\u200B%s)", confirmableBlock, timeStr, timeZone)
		}
		var taskData [][]string
		taskData = append(taskData, []string{"Amount(SWAN):", amount})
		taskData = append(taskData, []string{"Latest Block:", strconv.FormatUint(latestBlockNumber, 10)})
		taskData = append(taskData, []string{"Request Block:", strconv.Itoa(int(requestBlock))})
		taskData = append(taskData, []string{"Confirmable Block:", confirmableBlockStr})

		header := []string{"Withdraw View:"}
		NewVisualTable(header, taskData, []RowColor{}).SetAutoWrapText(false).Generate(false)

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

var sequencerCmd = &cli.Command{
	Name:      "sequencer",
	Usage:     "Manage the sequencer account",
	ArgsUsage: "[fromAddress] [amount]",
	Subcommands: []*cli.Command{
		sequencerTokenCmd,
		sequencerAddCmd,
		sequencerWithdrawCmd,
	},
	Before: func(c *cli.Context) error {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return err
		}
		return nil
	},
}

var sequencerTokenCmd = &cli.Command{
	Name:  "token",
	Usage: "Print the token of sequencer service",
	Action: func(cctx *cli.Context) error {
		token, err := computing.GetToken()
		if err != nil {
			return err
		}
		fmt.Println(token)
		return nil
	},
}

var sequencerAddCmd = &cli.Command{
	Name:  "add",
	Usage: "Send the amount to the sequencer contract address",
	Flags: []cli.Flag{
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
		txHash, err := localWallet.SequencerDeposit(ctx, fromAddress, amount, cpAccountAddress)
		if err != nil {
			return err
		}
		fmt.Printf("Transfer to Sequencer Account Tx Hash: %s \n", txHash)
		return nil
	},
}

var sequencerWithdrawCmd = &cli.Command{
	Name:  "withdraw",
	Usage: "Withdraw funds from the sequencer contract",
	Flags: []cli.Flag{
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
		txHash, err := localWallet.SequencerWithdraw(ctx, ownerAddress, amount, cpAccountAddress)
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
