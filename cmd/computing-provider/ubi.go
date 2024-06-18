package main

import (
	_ "embed"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"time"
)

var ubiTaskCmd = &cli.Command{
	Name:  "ubi",
	Usage: "Manage ubi tasks",
	Subcommands: []*cli.Command{
		listCmd,
		daemonCmd,
	},
}

var listCmd = &cli.Command{
	Name:  "list",
	Usage: "List ubi task",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "show-failed",
			Usage: "show failed/failing ubi tasks",
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
		fullFlag := cctx.Bool("verbose")
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		showFailed := cctx.Bool("show-failed")
		tailNum := cctx.Int("tail")
		var taskData [][]string
		var rowColorList []RowColor
		var taskList []*models.TaskEntity
		var err error
		if showFailed {
			taskList, err = computing.NewTaskService().GetAllTask(tailNum)
			if err != nil {
				return fmt.Errorf("failed get ubi task, error: %+v", err)
			}
		} else {
			taskList, err = computing.NewTaskService().GetTaskList(models.TASK_SUCCESS_STATUS, tailNum)
			if err != nil {
				return fmt.Errorf("failed get ubi task, error: %+v", err)
			}
		}

		if fullFlag {
			for i, task := range taskList {
				var errorMsg string
				if showFailed {
					errorMsg = task.Error
				}
				createTime := time.Unix(task.CreateTime, 0).Format("2006-01-02 15:04:05")
				taskData = append(taskData,
					[]string{strconv.Itoa(int(task.Id)), task.Contract, models.GetResourceTypeStr(task.ResourceType), models.UbiTaskTypeStr(task.Type), task.TxHash, models.TaskStatusStr(task.Status),
						fmt.Sprintf("%s", task.Reward), createTime, errorMsg})

				var rowColor []tablewriter.Colors
				if task.Status == models.TASK_RECEIVED_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
				} else if task.Status == models.TASK_RUNNING_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgCyanColor}}
				} else if task.Status == models.TASK_SUCCESS_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
				} else if task.Status == models.TASK_FAILED_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
				}

				rowColorList = append(rowColorList, RowColor{
					row:    i,
					column: []int{5},
					color:  rowColor,
				})
			}

		} else {
			for i, task := range taskList {
				createTime := time.Unix(task.CreateTime, 0).Format("2006-01-02 15:04:05")
				contract := shortenAddress(task.Contract)
				proofHash := shortenAddress(task.TxHash)

				var errorMsg string
				if showFailed {
					errorMsg = task.Error
				}
				taskData = append(taskData,
					[]string{strconv.Itoa(int(task.Id)), contract, models.GetResourceTypeStr(task.ResourceType), models.UbiTaskTypeStr(task.Type), proofHash, models.TaskStatusStr(task.Status),
						fmt.Sprintf("%s", task.Reward), createTime, errorMsg})

				var rowColor []tablewriter.Colors
				if task.Status == models.TASK_RECEIVED_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
				} else if task.Status == models.TASK_RUNNING_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgCyanColor}}
				} else if task.Status == models.TASK_SUCCESS_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
				} else if task.Status == models.TASK_FAILED_STATUS {
					rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
				}

				rowColorList = append(rowColorList, RowColor{
					row:    i,
					column: []int{5},
					color:  rowColor,
				})
			}

		}

		header := []string{"TASK ID", "Task Contract", "TASK TYPE", "ZK TYPE", "PROOF HASH", "STATUS", "REWARD", "CREATE TIME", "ERROR"}
		NewVisualTable(header, taskData, rowColorList).Generate(false)

		return nil

	},
}

var daemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a cp process",

	Action: func(cctx *cli.Context) error {
		logs.GetLogger().Info("Start a computing-provider client.")
		cpRepoPath, _ := os.LookupEnv("CP_PATH")

		resourceExporterContainerName := "resource-exporter"
		rsExist, err := computing.NewDockerService().CheckRunningContainer(resourceExporterContainerName)
		if err != nil {
			return fmt.Errorf("check %s container failed, error: %v", resourceExporterContainerName, err)
		}

		if !rsExist {
			if err = computing.RestartResourceExporter(); err != nil {
				logs.GetLogger().Errorf("restartResourceExporter failed, error: %v", err)
			}
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			logs.GetLogger().Fatal(err)
		}

		computing.SyncCpAccountInfo()
		computing.CronTaskForEcp()

		r := gin.Default()
		r.Use(cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			ValidateHeaders: false,
		}))
		pprof.Register(r)

		v1 := r.Group("/api/v1")
		router := v1.Group("/computing")

		router.GET("/cp", computing.GetCpResource)
		router.POST("/cp/ubi", computing.DoUbiTaskForDocker)
		router.POST("/cp/docker/receive/ubi", computing.ReceiveUbiProofForDocker)

		shutdownChan := make(chan struct{})
		httpStopper, err := util.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port), false)
		if err != nil {
			logs.GetLogger().Fatal("failed to start cp-api endpoint: %s", err)
		}

		finishCh := util.MonitorShutdown(shutdownChan,
			util.ShutdownHandler{Component: "cp-api", StopFunc: httpStopper},
		)
		<-finishCh

		return nil
	},
}

func shortenAddress(address string) string {
	if len(address) <= 12 {
		return address
	}
	return address[:7] + "...." + address[len(address)-7:]
}
