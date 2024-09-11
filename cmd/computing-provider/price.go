package main

import (
	"fmt"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/urfave/cli/v2"
	"os"
)

var priceCmd = &cli.Command{
	Name:  "price",
	Usage: "Manage cp hardware price",
	Subcommands: []*cli.Command{
		generateCmd,
		viewCmd,
	},
	Before: func(c *cli.Context) error {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}
		return nil
	},
}

var generateCmd = &cli.Command{
	Name:  "generate",
	Usage: "Generate resource price profile",
	Action: func(cctx *cli.Context) error {
		return computing.GeneratePriceConfig()
	},
}

var viewCmd = &cli.Command{
	Name:  "view",
	Usage: "View resource price configuration information",
	Action: func(cctx *cli.Context) error {
		config, err := computing.ReadPriceConfig()
		if err != nil {
			return err
		}
		var taskData [][]string
		for key, value := range config.Resources {
			taskData = append(taskData, []string{fmt.Sprintf("%s:", key), value})
		}
		header := []string{"CP Hardware Price Info:"}
		NewVisualTable(header, taskData, []RowColor{}).SetAutoWrapText(false).Generate(false)
		return nil
	},
}
