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
			Name:  "fcp",
			Usage: "Specify the fcp task",
		},
		&cli.BoolFlag{
			Name:  "ecp",
			Usage: "Specify the ecp task",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Usage:   "--verbose",
			Aliases: []string{"v"},
		},
		&cli.IntFlag{
			Name:  "tail",
			Usage: "Show the last number of lines. If not specified, all are displayed by default",
		},
	},
	Action: func(cctx *cli.Context) error {
		showCompleted := cctx.Bool("show-completed")
		fullFlag := cctx.Bool("verbose")
		fcpFlag := cctx.Bool("fcp")
		ecpFlag := cctx.Bool("ecp")
		if !fcpFlag && !ecpFlag {
			return fmt.Errorf("must specify one of fcp or ecp")
		}
		var taskType string
		if fcpFlag {
			taskType = "fcp"
		}
		if ecpFlag {
			taskType = "ecp"
		}
		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		tailNum := cctx.Int("tail")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		switch taskType {
		case "fcp":
			return fcpTaskList(showCompleted, fullFlag, tailNum)
		case "ecp":
			return ecpTaskList(tailNum)
		default:
			return fmt.Errorf("only support fcp and edge types")
		}
	},
}

var taskDetail = &cli.Command{
	Name:      "get",
	Usage:     "Get job detail info",
	ArgsUsage: "[job_uuid]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "fcp",
			Usage: "Specify the fcp task",
		},
		&cli.BoolFlag{
			Name:  "ecp",
			Usage: "Specify the ecp task",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf("incorrect number of arguments, got %d, missing args: job_uuid", cctx.NArg())
		}
		fcpFlag := cctx.Bool("fcp")
		ecpFlag := cctx.Bool("ecp")
		if !fcpFlag && !ecpFlag {
			return fmt.Errorf("must specify one of fcp or ecp")
		}
		var taskType string
		if fcpFlag {
			taskType = "fcp"
		}
		if ecpFlag {
			taskType = "ecp"
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		jobUuid := cctx.Args().First()

		if taskType == "fcp" {
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
			taskData = append(taskData, []string{"RESULT URL:", job.ResultUrl})
			taskData = append(taskData, []string{"CREATE TIME:", time.Unix(job.CreateTime, 0).Format("2006-01-02 15:04:05")})
			taskData = append(taskData, []string{"EXPIRE TIME:", time.Unix(job.ExpireTime, 0).Format("2006-01-02 15:04:05")})

			rowColor := getColor(job.Status)
			header := []string{"JOB UUID:", job.JobUuid}

			var rowColorList []RowColor
			rowColorList = append(rowColorList, RowColor{
				row:    6,
				column: []int{1},
				color:  rowColor,
			})
			NewVisualTable(header, taskData, rowColorList).SetAutoWrapText(false).Generate(false)
		} else {
			job, err := computing.NewEcpJobService().GetEcpJobByUuid(jobUuid)
			if err != nil {
				return fmt.Errorf("failed to get job, job_uuid: %s, error: %v", jobUuid, err)
			}

			var jobType = "Mining"
			if job.JobType == models.InferenceJobType {
				jobType = "Inference"
			}

			var taskData [][]string
			taskData = append(taskData, []string{"TASK NAME:", job.Name})
			taskData = append(taskData, []string{"TASK TYPE:", jobType})
			taskData = append(taskData, []string{"CONTAINER NAME:", job.ContainerName})
			taskData = append(taskData, []string{"GPU NAME:", job.GpuName})
			taskData = append(taskData, []string{"GPU INDEX:", job.GpuIndex})
			taskData = append(taskData, []string{"SERVICE URL:", job.ServiceUrl})
			taskData = append(taskData, []string{"PORTS:", job.PortMap})
			taskData = append(taskData, []string{"STATUS:", job.Status})
			taskData = append(taskData, []string{"CREATE TIME:", time.Unix(job.CreateTime, 0).Format("2006-01-02 15:04:05")})

			header := []string{"TASK UUID:", job.Uuid}
			NewVisualTable(header, taskData, []RowColor{}).SetAutoWrapText(false).Generate(false)
		}

		return nil
	},
}

var taskDelete = &cli.Command{
	Name:      "delete",
	Usage:     "Delete an task from the k8s",
	ArgsUsage: "[job_uuid]",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "fcp",
			Usage: "Specify the fcp task",
		},
		&cli.BoolFlag{
			Name:  "ecp",
			Usage: "Specify the ecp task",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf("incorrect number of arguments, got %d, missing args: task_uuid", cctx.NArg())
		}
		jobUuid := strings.ToLower(cctx.Args().First())

		fcpFlag := cctx.Bool("fcp")
		ecpFlag := cctx.Bool("ecp")
		if !fcpFlag && !ecpFlag {
			return fmt.Errorf("must specify one of fcp or ecp")
		}
		var taskType string
		if fcpFlag {
			taskType = "fcp"
		}
		if ecpFlag {
			taskType = "ecp"
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("failed to load config file, error: %+v", err)
		}

		if taskType == "fcp" {
			job, err := computing.NewJobService().GetJobEntityByJobUuid(jobUuid)
			if err != nil {
				return fmt.Errorf("failed to get job detail, job_uuid: %s, error: %+v", jobUuid, err)
			}

			if job.JobUuid == "" {
				return fmt.Errorf("not found the job_uuid='%s' job detail", jobUuid)
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
		} else {
			ecpJobEntity, err := computing.NewEcpJobService().GetEcpJobByUuid(jobUuid)
			if err != nil {
				return fmt.Errorf("failed to get job, job_uuid: %s, error: %v", jobUuid, err)
			}
			containerName := ecpJobEntity.ContainerName
			if err = computing.NewDockerService().RemoveContainerByName(containerName); err != nil {
				return fmt.Errorf("failed to remove container, job_uuid: %s, error: %v", jobUuid, err)
			}
			computing.NewEcpJobService().DeleteContainerByUuid(jobUuid)
			fmt.Printf("job_uuid: %s serivce successfully deleted \n", jobUuid)
		}
		return nil
	},
}

func getColor(status int) []tablewriter.Colors {
	var rowColor []tablewriter.Colors
	switch status {
	case models.JOB_DEPLOY_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
	case models.JOB_RUNNING_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
	case models.JOB_TERMINATED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case models.JOB_COMPLETED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiMagentaColor}}
	case models.JOB_RECEIVED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiBlueColor}}
	default:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiCyanColor}}
	}
	return rowColor
}

func fcpTaskList(showCompleted, fullFlag bool, tailNum int) error {
	var taskData [][]string
	var rowColorList []RowColor

	var jobStatus int
	if showCompleted {
		jobStatus = models.All_FLAG
	} else {
		jobStatus = models.UN_DELETEED_FLAG
	}

	list, err := computing.NewJobService().GetJobList(jobStatus, tailNum)
	if err != nil {
		return fmt.Errorf("get jobs failed, error: %+v", err)
	}
	for i, job := range list {

		expireTime := time.Unix(job.ExpireTime, 0).Format("2006-01-02 15:04:05")
		createTime := time.Unix(job.CreateTime, 0).Format("2006-01-02 15:04:05")

		var reward = "0.00"
		if len(strings.TrimSpace(job.Reward)) > 0 {
			reward = job.Reward
		}

		rowColor := getColor(job.Status)

		if fullFlag {
			taskData = append(taskData,
				[]string{job.JobUuid, job.TaskUuid, job.ResourceType, job.WalletAddress, job.SpaceUuid, job.Name, models.GetJobStatus(job.Status), reward, createTime, expireTime})

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
			if len(job.JobUuid) > 0 {
				jobUuidLen := len(job.JobUuid) - 1
				jobUuid = "..." + job.JobUuid[jobUuidLen-5:]
			}

			var spaceUuid string
			if len(job.SpaceUuid) > 0 {
				spaceUuidLen := len(job.SpaceUuid) - 1
				spaceUuid = "..." + job.SpaceUuid[spaceUuidLen-5:]
			}

			taskData = append(taskData,
				[]string{jobUuid, job.ResourceType, walletAddress, spaceUuid, job.Name, models.GetJobStatus(job.Status), createTime, expireTime})

			rowColorList = append(rowColorList, RowColor{
				row:    i,
				column: []int{5},
				color:  rowColor,
			})
		}
	}

	if fullFlag {
		header := []string{"JOB UUID", "TASK UUID", "TASK TYPE", "WALLET ADDRESS", "SPACE UUID", "SPACE NAME", "STATUS", "REWARD", "CREATE TIME", "EXPIRE TIME"}
		NewVisualTable(header, taskData, rowColorList).Generate(true)
	} else {
		header := []string{"JOB UUID", "TASK TYPE", "WALLET ADDRESS", "SPACE UUID", "SPACE NAME", "STATUS", "CREATE TIME", "EXPIRE TIME"}
		NewVisualTable(header, taskData, rowColorList).Generate(true)
	}
	return nil
}

func ecpTaskList(tailNum int) error {
	var taskData [][]string
	var rowColorList []RowColor

	ecpJobs, err := computing.NewEcpJobService().GetEcpJobsByLimit(tailNum)
	if err != nil {
		return err
	}

	containerStatus, err := computing.NewDockerService().GetContainerStatus()
	if err != nil {
		return err
	}

	for i, entity := range ecpJobs {
		createTime := time.Unix(entity.CreateTime, 0).Format("2006-01-02 15:04:05")
		statusStr := "terminated"
		if status, ok := containerStatus[entity.ContainerName]; ok {
			if entity.Status != "terminated" {
				computing.NewEcpJobService().UpdateEcpJobEntity(entity.Uuid, status)
			}
			statusStr = status
		}
		taskData = append(taskData, []string{entity.Uuid, entity.Name, entity.Image, entity.ContainerName, statusStr, fmt.Sprintf("%.4f", entity.Reward), createTime})
		rowColorList = append(rowColorList, RowColor{
			row:    i,
			column: []int{4},
			color:  getContainerStatusColor(statusStr),
		})
	}
	header := []string{"TASK UUID", "TASK NAME", "IMAGE NAME", "CONTAINER NAME", "CONTAINER STATUS", "REWARD", "CREATE TIME"}
	NewVisualTable(header, taskData, rowColorList).Generate(true)
	return nil
}

func getContainerStatusColor(status string) []tablewriter.Colors {
	var rowColor []tablewriter.Colors
	switch status {
	case "created":
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
	case "running":
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
	case "removing":
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case "paused":
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiMagentaColor}}
	case "exited":
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiBlueColor}}
	case "terminated":
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiBlueColor}}
	default:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgHiCyanColor}}
	}
	return rowColor
}
