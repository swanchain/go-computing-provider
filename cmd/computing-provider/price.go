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
		hardwarePrice, err := computing.ReadPriceConfig()
		if err != nil {
			return err
		}

		hardwareFields, err := computing.GetStructByTag(hardwarePrice)
		if err != nil {
			return err
		}

		for k, v := range hardwarePrice.GpusPrice {
			hardwareFields = append(hardwareFields, computing.HardwareField{
				Name:  k,
				Value: v,
			})
		}

		var taskData [][]string
		for _, field := range hardwareFields {
			var valStr string
			switch field.Name {
			case "TARGET_CPU":
				valStr = field.Value + " SWAN/thread-hour"
				break
			case "TARGET_GPU_DEFAULT":
				valStr = field.Value + " SWAN/Default GPU unit a hour"
				break
			case "TARGET_MEMORY":
				fallthrough
			case "TARGET_HD_EPHEMERAL":
				valStr = field.Value + " SWAN/GB-hour"
				break
			default:
				valStr = field.Value + " SWAN/GPU unit a hour"
			}
			taskData = append(taskData, []string{fmt.Sprintf("%s:", field.Name), valStr})
		}
		header := []string{"CP Hardware Price Info:"}
		NewVisualTable(header, taskData, []RowColor{}).SetAutoWrapText(false).Generate(false)
		return nil
	},
}
