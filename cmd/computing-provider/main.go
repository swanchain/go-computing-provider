package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/swanchain/go-computing-provider/build"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	FlagCpRepo = "repo"
)

var FlagRepo = &cli.StringFlag{
	Name:    FlagCpRepo,
	Usage:   "repo directory for computing-provider client",
	Value:   "~/.swan/computing",
	EnvVars: []string{"CP_PATH"},
}

func main() {
	app := &cli.App{
		Name:                 "computing-provider",
		Usage:                "Swanchain decentralized computing network client",
		EnableBashCompletion: true,
		Version:              build.UserVersion(),
		Flags: []cli.Flag{
			FlagRepo,
		},
		Commands: []*cli.Command{
			initCmd,
			runCmd,
			infoCmd,
			stateCmd,
			accountCmd,
			taskCmd,
			walletCmd,
			collateralCmd,
			ubiTaskCmd,
		},
		Before: func(c *cli.Context) error {
			cpRepoPath, err := homedir.Expand(c.String(FlagRepo.Name))
			if err != nil {
				return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
			}
			if _, err = os.Stat(cpRepoPath); os.IsNotExist(err) {
				return fmt.Errorf("CP_PATH: %s, no such directory", cpRepoPath)
			}
			os.Setenv("CP_PATH", cpRepoPath)
			db.InitDb(cpRepoPath)

			return nil
		},
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
	}
}
