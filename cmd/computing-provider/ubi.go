package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/build"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/urfave/cli/v2"
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
			taskList, err = computing.NewTaskService().GetTaskList(tailNum)
			if err != nil {
				return fmt.Errorf("failed get ubi task, error: %+v", err)
			}
		} else {
			taskList, err = computing.NewTaskService().GetTaskList(tailNum, []int{
				models.TASK_RECEIVED_STATUS,
				models.TASK_RUNNING_STATUS,
				models.TASK_SUBMITTED_STATUS,
				models.TASK_VERIFIED_STATUS,
				models.TASK_REWARDED_STATUS,
			}...)
			if err != nil {
				return fmt.Errorf("failed get ubi task, error: %+v", err)
			}
		}

		if fullFlag {
			for i, task := range taskList {
				createTime := time.Unix(task.CreateTime, 0).Format("2006-01-02 15:04:05")
				var sequencerStr string
				var contract string
				if task.Sequencer == 1 {
					sequencerStr = "YES"
					if task.SequenceTaskAddr != "" {
						contract = task.SequenceTaskAddr
					}
				} else if task.Sequencer == 0 {
					sequencerStr = "NO"
					contract = task.Contract
				} else {
					sequencerStr = ""
				}

				var taskId string
				if task.Type == models.Mining {
					taskId = task.Uuid
				} else {
					taskId = strconv.Itoa(int(task.Id))
				}

				taskData = append(taskData,
					[]string{taskId, contract, models.GetResourceTypeStr(task.ResourceType), models.UbiTaskTypeStr(task.Type),
						task.CheckCode, task.Sign, models.TaskStatusStr(task.Status), sequencerStr, createTime})

				rowColorList = append(rowColorList, RowColor{
					row:    i,
					column: []int{6},
					color:  getStatusColor(task.Status),
				})
			}
			header := []string{"TASK ID", "TASK CONTRACT", "TASK TYPE", "ZK TYPE", "CHECK CODE", "SIGNATURE", "STATUS", "SEQUENCER", "CREATE TIME"}
			NewVisualTable(header, taskData, rowColorList).Generate(false)

		} else {
			for i, task := range taskList {
				createTime := time.Unix(task.CreateTime, 0).Format("2006-01-02 15:04:05")
				var sequencerStr string
				var contract string
				if task.Sequencer == 1 {
					sequencerStr = "YES"
					if task.SequenceTaskAddr != "" {
						contract = task.SequenceTaskAddr
					}
				} else if task.Sequencer == 0 {
					sequencerStr = "NO"
					contract = task.Contract
				} else {
					sequencerStr = ""
				}

				var taskId string
				if task.Type == models.Mining {
					taskId = task.Uuid
				} else {
					taskId = strconv.Itoa(int(task.Id))
				}

				taskData = append(taskData,
					[]string{taskId, contract, models.GetResourceTypeStr(task.ResourceType), models.UbiTaskTypeStr(task.Type),
						models.TaskStatusStr(task.Status), sequencerStr, createTime})

				rowColorList = append(rowColorList, RowColor{
					row:    i,
					column: []int{4},
					color:  getStatusColor(task.Status),
				})
			}

			header := []string{"TASK ID", "TASK CONTRACT", "TASK TYPE", "ZK TYPE", "STATUS", "SEQUENCER", "CREATE TIME"}
			NewVisualTable(header, taskData, rowColorList).Generate(false)
		}
		return nil

	},
}

var daemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a cp process",

	Action: func(cctx *cli.Context) error {
		logs.GetLogger().Info("Starting a computing-provider client.")
		cpRepoPath, _ := os.LookupEnv("CP_PATH")

		computing.NewDockerService().CleanResourceForDocker(true)

		resourceExporterContainerName := "resource-exporter"
		rsExist, version, err := computing.NewDockerService().CheckRunningContainer(resourceExporterContainerName)
		if err != nil {
			return fmt.Errorf("check %s container failed, error: %v", resourceExporterContainerName, err)
		}

		if version != "" {
			if errMsg := util.CheckVersion(build.ResourceExporterVersion, version); errMsg != nil {
				logs.GetLogger().Fatalf("resource-exporter %s", errMsg)
			}
		}

		if !rsExist {
			if err = computing.RestartResourceExporter(); err != nil {
				logs.GetLogger().Errorf("restartResourceExporter failed, error: %v", err)
			}
		}

		traefikServiceContainerName := "traefik-service"
		tsExist, _, err := computing.NewDockerService().CheckRunningContainer(traefikServiceContainerName)
		if err != nil {
			return fmt.Errorf("check %s container failed, error: %v", traefikServiceContainerName, err)
		}
		if !tsExist {
			if err = computing.RestartTraefikService(); err != nil {
				logs.GetLogger().Errorf("restartTraefikService failed, error: %v", err)
			}
		}

		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			logs.GetLogger().Fatal(err)
		}
		logs.GetLogger().Info("Your config file is:", filepath.Join(cpRepoPath, "config.toml"))

		computing.SyncCpAccountInfo()
		computing.CronTaskForEcp()

		gin.SetMode(gin.ReleaseMode)
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

		router := r.Group("/api/v1/computing")
		router.GET("/cp", computing.GetCpResource)
		router.GET("/cp/metrics", computing.GetUbiResourceExporterMetrics)
		router.POST("/cp/ubi", computing.DoUbiTaskForDocker)
		router.POST("/cp/docker/receive/ubi", computing.ReceiveUbiProof)

		ecpImageService := computing.NewImageJobService()
		router.POST("/cp/deploy/check", ecpImageService.CheckJobCondition)
		router.GET("/cp/price", computing.GetPrice)
		router.POST("/cp/deploy", ecpImageService.DeployJob)
		router.GET("/cp/job/status", ecpImageService.GetJobStatus)
		router.GET("/cp/job/log", ecpImageService.DockerLogsHandler)
		router.DELETE("/cp/job/:job_uuid", ecpImageService.DeleteJob)
		router.POST("/cp/zk_task", computing.DoZkTask)

		shutdownChan := make(chan struct{})
		httpStopper, err := util.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port), false)
		if err != nil {
			logs.GetLogger().Fatal("failed to start cp-api endpoint: %s", err)
		}
		logs.GetLogger().Infof("CP service started successfully, listening on port: %d", conf.GetConfig().API.Port)

		finishCh := util.MonitorShutdown(shutdownChan,
			util.ShutdownHandler{Component: "cp-api", StopFunc: httpStopper},
		)
		<-finishCh

		return nil
	},
}

func getStatusColor(taskStatus int) []tablewriter.Colors {
	var rowColor []tablewriter.Colors
	switch taskStatus {
	case models.TASK_REJECTED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case models.TASK_RECEIVED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgYellowColor}}
	case models.TASK_RUNNING_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgCyanColor}}
	case models.TASK_SUBMITTED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgBlueColor}}
	case models.TASK_FAILED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case models.TASK_VERIFIED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgBlueColor}}
	case models.TASK_REWARDED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
	case models.TASK_INVALID_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case models.TASK_TIMEOUT_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case models.TASK_VERIFYFAILED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case models.TASK_REPEATED_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
	case models.TASK_NSC_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
	case models.TASK_UNKNOWN_STATUS:
		rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgBlackColor}}
	}
	return rowColor
}
