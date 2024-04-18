package main

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal"
	"github.com/swanchain/go-computing-provider/internal/pkg"
	"github.com/swanchain/go-computing-provider/internal/v1/service"
	"github.com/urfave/cli/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"os"
	"strings"
	"time"
)

var taskCmd = &cli.Command{
	Name:  "task",
	Usage: "Manage tasks",
	Subcommands: []*cli.Command{
		taskList,
		taskDetail,
		taskDelete,
	},
}

var taskList = &cli.Command{
	Name:  "list",
	Usage: "List task",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Usage:   "--verbose",
			Aliases: []string{"v"},
		},
	},
	Action: func(cctx *cli.Context) error {

		fullFlag := cctx.Bool("verbose")

		cpPath, exit := os.LookupEnv("CP_PATH")
		if !exit {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=xxx")
		}
		if err := conf.InitConfig(cpPath); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		conn := pkg.GetRedisClient()
		prefix := internal.REDIS_SPACE_PREFIX + "*"
		keys, err := redis.Strings(conn.Do("KEYS", prefix))
		if err != nil {
			return fmt.Errorf("failed get redis %s prefix, error: %+v", prefix, err)
		}

		var taskData [][]string
		var rowColorList []pkg.RowColor
		var number int
		for _, key := range keys {
			jobDetail, err := service.RetrieveJobMetadata(key)
			if err != nil {
				return fmt.Errorf("failed get job detail: %s, error: %+v", key, err)
			}

			k8sService := pkg.NewK8sService()
			status, err := k8sService.GetDeploymentStatus(jobDetail.WalletAddress, jobDetail.SpaceUuid)
			if err != nil {
				return fmt.Errorf("failed get job status: %s, error: %+v", jobDetail.JobUuid, err)
			}

			var fullSpaceUuid string
			if len(jobDetail.DeployName) > 0 {
				fullSpaceUuid = jobDetail.DeployName[7:]
			}

			if fullFlag {
				taskData = append(taskData,
					[]string{jobDetail.TaskUuid, jobDetail.TaskType, jobDetail.WalletAddress, fullSpaceUuid, jobDetail.SpaceName, status})
			} else {
				var walletAddress string
				if len(jobDetail.WalletAddress) > 0 {
					walletAddress = jobDetail.WalletAddress[:5] + "..." + jobDetail.WalletAddress[37:]
				}

				var taskUuid string
				if len(jobDetail.TaskUuid) > 0 {
					taskUuid = "..." + jobDetail.TaskUuid[26:]
				}

				var spaceUuid string
				if len(jobDetail.SpaceUuid) > 0 {
					spaceUuid = "..." + jobDetail.SpaceUuid[26:]
				}

				taskData = append(taskData,
					[]string{taskUuid, jobDetail.TaskType, walletAddress, spaceUuid, jobDetail.SpaceName, status})
			}

			var rowColor []tablewriter.Colors
			if status == "Pending" {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
			} else if status == "Running" {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			} else {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
			}

			rowColorList = append(rowColorList, pkg.RowColor{
				Row:    number,
				Column: []int{5},
				Color:  rowColor,
			})

			number++
		}

		header := []string{"TASK UUID", "TASK TYPE", "WALLET ADDRESS", "SPACE UUID", "SPACE NAME", "STATUS"}
		pkg.NewVisualTable(header, taskData, rowColorList).Generate(true)

		return nil

	},
}

var taskDetail = &cli.Command{
	Name:      "get",
	Usage:     "Get task detail info",
	ArgsUsage: "[space_uuid]",
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf("incorrect number of arguments, got %d, missing args: space_uuid", cctx.NArg())
		}

		cpPath, exit := os.LookupEnv("CP_PATH")
		if !exit {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=xxx")
		}
		if err := conf.InitConfig(cpPath); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}
		pkg.GetRedisClient()

		spaceUuid := internal.REDIS_SPACE_PREFIX + cctx.Args().First()
		jobDetail, err := service.RetrieveJobMetadata(spaceUuid)
		if err != nil {
			return fmt.Errorf("failed get job detail: %s, error: %+v", spaceUuid, err)
		}

		k8sService := pkg.NewK8sService()
		status, err := k8sService.GetDeploymentStatus(jobDetail.WalletAddress, jobDetail.SpaceUuid)
		if err != nil {
			return fmt.Errorf("failed get job status: %s, error: %+v", jobDetail.JobUuid, err)
		}

		var taskData [][]string
		taskData = append(taskData, []string{"TASK TYPE:", jobDetail.TaskType})
		taskData = append(taskData, []string{"WALLET ADDRESS:", jobDetail.WalletAddress})
		taskData = append(taskData, []string{"SPACE NAME:", jobDetail.SpaceName})
		taskData = append(taskData, []string{"SPACE URL:", jobDetail.Url})
		taskData = append(taskData, []string{"HARDWARE:", jobDetail.Hardware})
		taskData = append(taskData, []string{"STATUS:", status})

		var rowColor []tablewriter.Colors
		if status == "Pending" {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}, {tablewriter.Bold, tablewriter.FgWhiteColor}}
		} else if status == "Running" {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}, {tablewriter.Bold, tablewriter.FgWhiteColor}}
		} else {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}, {tablewriter.Bold, tablewriter.FgWhiteColor}}
		}

		header := []string{"TASK UUID:", jobDetail.TaskUuid}

		var rowColorList []pkg.RowColor
		rowColorList = append(rowColorList, pkg.RowColor{
			Row:    6,
			Column: []int{1},
			Color:  rowColor,
		})
		pkg.NewVisualTable(header, taskData, rowColorList).Generate(false)
		return nil
	},
}

var taskDelete = &cli.Command{
	Name:      "delete",
	Usage:     "Delete an task from the k8s",
	ArgsUsage: "[space_uuid]",
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf("incorrect number of arguments, got %d, missing args: space_uuid", cctx.NArg())
		}

		cpPath, exit := os.LookupEnv("CP_PATH")
		if !exit {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=xxx")
		}
		if err := conf.InitConfig(cpPath); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}
		pkg.GetRedisClient()

		spaceUuid := strings.ToLower(cctx.Args().First())
		jobDetail, err := service.RetrieveJobMetadata(internal.REDIS_SPACE_PREFIX + spaceUuid)
		if err != nil {
			return fmt.Errorf("failed get job detail: %s, error: %+v", spaceUuid, err)
		}

		deployName := internal.K8S_DEPLOY_NAME_PREFIX + spaceUuid
		namespace := internal.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(jobDetail.WalletAddress)
		k8sService := pkg.NewK8sService()
		if err := k8sService.DeleteDeployment(context.TODO(), namespace, deployName); err != nil && !errors.IsNotFound(err) {
			return err
		}
		time.Sleep(6 * time.Second)

		if err := k8sService.DeleteDeployRs(context.TODO(), namespace, spaceUuid); err != nil && !errors.IsNotFound(err) {
			return err
		}

		conn := pkg.GetRedisClient()
		conn.Do("DEL", redis.Args{}.AddFlat(internal.REDIS_SPACE_PREFIX+spaceUuid)...)

		return nil
	},
}
