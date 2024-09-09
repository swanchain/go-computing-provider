package main

import (
	"context"
	"fmt"
	"github.com/swanchain/go-computing-provider/constants"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/urfave/cli/v2"
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

			expireTime := time.Unix(job.ExpireTime, 0).Format("2006-01-02 15:04:05")

			var reward = "0.00"
			if len(strings.TrimSpace(job.Reward)) > 0 {
				reward = job.Reward
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
			default:
				rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiCyanColor}}
			}

			if fullFlag {
				taskData = append(taskData,
					[]string{job.JobUuid, job.TaskUuid, job.ResourceType, job.WalletAddress, job.SpaceUuid, job.Name, models.GetJobStatus(job.Status), reward, expireTime})

				rowColorList = append(rowColorList, RowColor{
					row:    i,
					column: []int{6},
					color:  rowColor,
				})

			} else {
				var walletAddress string
				if len(job.WalletAddress) > 0 {
					walletAddress = job.WalletAddress[:5] + "..." + job.WalletAddress[37:]
				}

				var jobUuid string
				if len(job.TaskUuid) > 0 {
					jobUuid = "..." + job.TaskUuid[26:]
				}

				var spaceUuid string
				if len(job.SpaceUuid) > 0 {
					spaceUuid = "..." + job.SpaceUuid[26:]
				}

				taskData = append(taskData,
					[]string{jobUuid, job.ResourceType, walletAddress, spaceUuid, job.Name, models.GetJobStatus(job.Status), expireTime})

				rowColorList = append(rowColorList, RowColor{
					row:    i,
					column: []int{5},
					color:  rowColor,
				})
			}
		}

		if fullFlag {
			header := []string{"JOB UUID", "TASK UUID", "TASK TYPE", "WALLET ADDRESS", "SPACE UUID", "SPACE NAME", "STATUS", "REWARD", "EXPIRE TIME"}
			NewVisualTable(header, taskData, rowColorList).Generate(true)
		} else {
			header := []string{"JOB UUID", "TASK TYPE", "WALLET ADDRESS", "SPACE UUID", "SPACE NAME", "STATUS", "EXPIRE TIME"}
			NewVisualTable(header, taskData, rowColorList).Generate(true)
		}

		return nil
	},
}

var taskDetail = &cli.Command{
	Name:      "get",
	Usage:     "Get task detail info",
	ArgsUsage: "[job_uuid]",
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

		jobUuid := cctx.Args().First()
		job, err := computing.NewJobService().GetJobEntityByJobUuid(jobUuid)
		if err != nil {
			return fmt.Errorf("job_uuid: %s, get job detail failed, error: %+v", jobUuid, err)
		}

		var taskData [][]string
		taskData = append(taskData, []string{"TASK UUID:", job.TaskUuid})
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
		default:
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiCyanColor}}
		}

		header := []string{"JOB UUID:", job.JobUuid}

		var rowColorList []RowColor
		rowColorList = append(rowColorList, RowColor{
			row:    7,
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
	ArgsUsage: "[job_uuid]",
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf("incorrect number of arguments, got %d, missing args: task_uuid", cctx.NArg())
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("failed to load config file, error: %+v", err)
		}

		jobUuid := strings.ToLower(cctx.Args().First())
		job, err := computing.NewJobService().GetJobEntityByJobUuid(jobUuid)
		if err != nil {
			return fmt.Errorf("failed to get job detail, job_uuid: %s, error: %+v", jobUuid, err)
		}

		serviceName := constants.K8S_SERVICE_NAME_PREFIX + jobUuid
		ingressName := constants.K8S_INGRESS_NAME_PREFIX + jobUuid

		k8sService := computing.NewK8sService()
		k8sService.DeleteIngress(context.TODO(), job.NameSpace, ingressName)
		k8sService.DeleteService(context.TODO(), job.NameSpace, serviceName)
		k8sService.DeleteDeployment(context.TODO(), job.NameSpace, job.K8sDeployName)
		time.Sleep(3 * time.Second)
		k8sService.DeleteDeployRs(context.TODO(), job.NameSpace, job.JobUuid)

		computing.NewJobService().DeleteJobEntityByJobUuId(job.JobUuid, models.JOB_TERMINATED_STATUS)
		fmt.Printf("job_uuid: %s space serivce successfully deleted \n", jobUuid)
		return nil
	},
}
