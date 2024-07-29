package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/urfave/cli/v2"
	"k8s.io/apimachinery/pkg/api/errors"
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
			Name:  "show-completed",
			Usage: "Display completed jobs",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Usage:   "--verbose",
			Aliases: []string{"v"},
		},
	},
	Action: func(cctx *cli.Context) error {
		showCompleted := cctx.Bool("show-completed")
		fullFlag := cctx.Bool("verbose")
		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		var taskData [][]string
		var rowColorList []RowColor

		var jobStatus int
		if showCompleted {
			jobStatus = models.All_FLAG
		} else {
			jobStatus = models.UN_DELETEED_FLAG
		}

		list, err := computing.NewJobService().GetJobList(jobStatus)
		if err != nil {
			return fmt.Errorf("get jobs failed, error: %+v", err)
		}
		for i, job := range list {
			var fullSpaceUuid string
			if len(job.K8sDeployName) > 0 {
				fullSpaceUuid = job.K8sDeployName[7:]
			}

			expireTime := time.Unix(job.ExpireTime, 0).Format("2006-01-02 15:04:05")

			if fullFlag {
				taskData = append(taskData,
					[]string{job.TaskUuid, job.ResourceType, job.WalletAddress, fullSpaceUuid, job.Name, models.GetJobStatus(job.Status), expireTime})
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
					[]string{taskUuid, job.ResourceType, walletAddress, spaceUuid, job.Name, models.GetJobStatus(job.Status), job.Reward, expireTime})
			}

			var rowColor []tablewriter.Colors
			switch job.Status {
			case models.JOB_DEPLOY_STATUS:
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
			case models.JOB_RUNNING_STATUS:
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			case models.JOB_TERMINATED_STATUS:
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
			case models.JOB_COMPLETED_STATUS:
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiMagentaColor}}
			}

			rowColorList = append(rowColorList, RowColor{
				row:    i,
				column: []int{5},
				color:  rowColor,
			})
		}

		header := []string{"TASK UUID", "TASK TYPE", "WALLET ADDRESS", "SPACE UUID", "SPACE NAME", "STATUS", "REWARD", "EXPIRE TIME"}
		NewVisualTable(header, taskData, rowColorList).Generate(true)
		return nil
	},
}

var taskDetail = &cli.Command{
	Name:      "get",
	Usage:     "Get task detail info",
	ArgsUsage: "[task_uuid]",
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf("incorrect number of arguments, got %d, missing args: task_uuid", cctx.NArg())
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		taskUuid := cctx.Args().First()
		job, err := computing.NewJobService().GetJobEntityByTaskUuid(taskUuid)
		if err != nil {
			return fmt.Errorf("task_uuid: %s, get job detail failed, error: %+v", taskUuid, err)
		}

		var taskData [][]string
		taskData = append(taskData, []string{"TASK TYPE:", job.ResourceType})
		taskData = append(taskData, []string{"WALLET ADDRESS:", job.WalletAddress})
		taskData = append(taskData, []string{"SPACE NAME:", job.Name})
		taskData = append(taskData, []string{"SPACE URL:", job.RealUrl})
		taskData = append(taskData, []string{"HARDWARE:", job.Hardware})
		taskData = append(taskData, []string{"STATUS:", models.GetJobStatus(job.Status)})

		var rowColor []tablewriter.Colors
		switch job.Status {
		case models.JOB_DEPLOY_STATUS:
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
		case models.JOB_RUNNING_STATUS:
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
		case models.JOB_TERMINATED_STATUS:
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
		case models.JOB_COMPLETED_STATUS:
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiMagentaColor}}
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
	ArgsUsage: "[task_uuid]",
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf("incorrect number of arguments, got %d, missing args: task_uuid", cctx.NArg())
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		taskUuid := strings.ToLower(cctx.Args().First())
		job, err := computing.NewJobService().GetJobEntityByTaskUuid(taskUuid)
		if err != nil {
			return fmt.Errorf("failed get job detail: %s, error: %+v", taskUuid, err)
		}

		deployName := constants.K8S_DEPLOY_NAME_PREFIX + job.SpaceUuid
		namespace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(job.WalletAddress)
		k8sService := computing.NewK8sService()
		if err := k8sService.DeleteDeployment(context.TODO(), namespace, deployName); err != nil && !errors.IsNotFound(err) {
			return err
		}
		time.Sleep(3 * time.Second)

		if err := k8sService.DeleteDeployRs(context.TODO(), namespace, job.SpaceUuid); err != nil && !errors.IsNotFound(err) {
			return err
		}
		computing.NewJobService().DeleteJobEntityBySpaceUuId(job.SpaceUuid, models.JOB_TERMINATED_STATUS)
		fmt.Printf("space_uuid: %s space serivce successfully deleted \n", job.SpaceUuid)
		return nil
	},
}
