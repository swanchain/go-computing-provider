package main

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/computing"
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
		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		var taskData [][]string
		var rowColorList []RowColor

		list, err := computing.NewJobService().GetJobList()
		if err != nil {
			return fmt.Errorf("get jobs failed, error: %+v", err)
		}
		for i, job := range list {
			k8sService := computing.NewK8sService()
			status, err := k8sService.GetDeploymentStatus(job.WalletAddress, job.SpaceUuid)
			if err != nil {
				return fmt.Errorf("failed get job status: %s, error: %+v", job.JobUuid, err)
			}
			var fullSpaceUuid string
			if len(job.K8sDeployName) > 0 {
				fullSpaceUuid = job.K8sDeployName[7:]
			}

			expireTime := time.Unix(job.ExpireTime, 0).Format("2006-01-02 15:04:05")

			if fullFlag {
				taskData = append(taskData,
					[]string{job.TaskUuid, job.ResourceType, job.WalletAddress, fullSpaceUuid, job.Name, status, expireTime})
			} else {
				var walletAddress string
				if len(job.WalletAddress) > 0 {
					walletAddress = job.WalletAddress[:5] + "..." + job.WalletAddress[37:]
				}

				var taskUuid string
				if len(job.TaskUuid) > 0 {
					taskUuid = "..." + job.TaskUuid[26:]
				}

				var spaceUuid string
				if len(job.SpaceUuid) > 0 {
					spaceUuid = "..." + job.SpaceUuid[26:]
				}

				taskData = append(taskData,
					[]string{taskUuid, job.ResourceType, walletAddress, spaceUuid, job.Name, status, expireTime})
			}

			var rowColor []tablewriter.Colors
			if status == "Pending" {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
			} else if status == "Running" {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			} else {
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
			}

			rowColorList = append(rowColorList, RowColor{
				row:    i,
				column: []int{5},
				color:  rowColor,
			})
		}

		header := []string{"TASK UUID", "TASK TYPE", "WALLET ADDRESS", "SPACE UUID", "SPACE NAME", "STATUS", "EXPIRE TIME"}
		NewVisualTable(header, taskData, rowColorList).Generate(true)
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

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		spaceUuid := cctx.Args().First()
		job, err := computing.NewJobService().GetJobEntityBySpaceUuid(spaceUuid)
		if err != nil {
			return fmt.Errorf("failed get job detail: %s, error: %+v", spaceUuid, err)
		}

		k8sService := computing.NewK8sService()
		status, err := k8sService.GetDeploymentStatus(job.WalletAddress, job.SpaceUuid)
		if err != nil {
			return fmt.Errorf("failed get job status: %s, error: %+v", job.JobUuid, err)
		}

		var taskData [][]string
		taskData = append(taskData, []string{"TASK TYPE:", job.ResourceType})
		taskData = append(taskData, []string{"WALLET ADDRESS:", job.WalletAddress})
		taskData = append(taskData, []string{"SPACE NAME:", job.Name})
		taskData = append(taskData, []string{"SPACE URL:", job.RealUrl})
		taskData = append(taskData, []string{"HARDWARE:", job.Hardware})
		taskData = append(taskData, []string{"STATUS:", status})

		var rowColor []tablewriter.Colors
		if status == "Pending" {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}, {tablewriter.Bold, tablewriter.FgYellowColor}}
		} else if status == "Running" {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}, {tablewriter.Bold, tablewriter.FgGreenColor}}
		} else {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}, {tablewriter.Bold, tablewriter.FgCyanColor}}
		}

		header := []string{"TASK UUID:", job.TaskUuid}

		var rowColorList []RowColor
		rowColorList = append(rowColorList, RowColor{
			row:    6,
			column: []int{1},
			color:  rowColor,
		})
		NewVisualTable(header, taskData, rowColorList).Generate(false)
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

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		spaceUuid := strings.ToLower(cctx.Args().First())
		job, err := computing.NewJobService().GetJobEntityBySpaceUuid(spaceUuid)
		if err != nil {
			return fmt.Errorf("failed get job detail: %s, error: %+v", spaceUuid, err)
		}

		deployName := constants.K8S_DEPLOY_NAME_PREFIX + spaceUuid
		namespace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(job.WalletAddress)
		k8sService := computing.NewK8sService()
		if err := k8sService.DeleteDeployment(context.TODO(), namespace, deployName); err != nil && !errors.IsNotFound(err) {
			return err
		}
		time.Sleep(6 * time.Second)

		if err := k8sService.DeleteDeployRs(context.TODO(), namespace, spaceUuid); err != nil && !errors.IsNotFound(err) {
			return err
		}

		computing.NewJobService().DeleteJobEntityBySpaceUuId(spaceUuid)
		return nil
	},
}
