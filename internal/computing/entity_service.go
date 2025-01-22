package computing

import (
	"fmt"
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/gorm"
	"time"
)

type TaskService struct {
	*gorm.DB
}

func (taskServ TaskService) GetTaskList(tailNum int, taskStatus ...int) (list []*models.TaskEntity, err error) {
	if tailNum > 0 {
		if len(taskStatus) > 0 {
			err = taskServ.Model(&models.TaskEntity{}).Where("status in ?", taskStatus).Order("create_time desc").Limit(tailNum).Find(&list).Error
		} else {
			err = taskServ.Model(&models.TaskEntity{}).Order("create_time desc").Limit(tailNum).Find(&list).Error
		}
	} else {
		if len(taskStatus) > 0 {
			err = taskServ.Model(&models.TaskEntity{}).Where("status in ?", taskStatus).Order("create_time desc").Find(&list).Error
		} else {
			err = taskServ.Model(&models.TaskEntity{}).Order("create_time desc").Find(&list).Error
		}
	}
	if err != nil {
		return nil, err
	}

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return
}

func (taskServ TaskService) SaveTaskEntity(task *models.TaskEntity) (err error) {
	if task.Status == models.TASK_FAILED_STATUS || task.Status == models.TASK_SUBMITTED_STATUS {
		task.EndTime = time.Now().Unix()
	}
	return taskServ.Save(task).Error
}

func (taskServ TaskService) UpdateTaskStatusById(taskId int, status int) (err error) {
	return taskServ.Model(&models.TaskEntity{}).Where("id=?", taskId).Update("status", status).Error
}

func (taskServ TaskService) UpdateTaskStatusByUuid(uuid string, status int) (err error) {
	return taskServ.Model(&models.TaskEntity{}).Where("uuid=?", uuid).Update("status", status).Error
}

func (taskServ TaskService) GetTaskByUuid(uuid string) (*models.TaskEntity, error) {
	var taskEntity models.TaskEntity
	err := taskServ.Model(&models.TaskEntity{}).Where("uuid=?", uuid).Limit(1).Find(&taskEntity).Error
	return &taskEntity, err
}

func (taskServ TaskService) UpdateTaskEntityByTaskId(task *models.TaskEntity) (err error) {
	return taskServ.Model(&models.TaskEntity{}).Where("id=?", task.Id).Updates(task).Error
}

func (taskServ TaskService) UpdateTaskEntityByTaskUuId(task *models.TaskEntity) (err error) {
	return taskServ.Model(&models.TaskEntity{}).Where("uuid=?", task.Uuid).Updates(map[string]interface{}{
		"status": task.Status,
	}).Error
}

func (taskServ TaskService) GetTaskEntity(taskId int64) (*models.TaskEntity, error) {
	var taskEntity models.TaskEntity
	err := taskServ.First(&taskEntity, taskId).Error
	return &taskEntity, err
}

func (taskServ TaskService) GetTaskListNoRewardForFilC2() (list []*models.TaskEntity, err error) {
	err = taskServ.Model(&models.TaskEntity{}).Where("sequencer=? and sequence_task_addr =='' and uuid=''", models.TaskSequencer).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

func (taskServ TaskService) GetTaskListNoRewardForMining() (list []*models.TaskEntity, err error) {
	err = taskServ.Model(&models.TaskEntity{}).Where("uuid !='' and (status !=? or status !=? or status !=?) ", models.TASK_TIMEOUT_STATUS,
		models.TASK_VERIFYFAILED_STATUS, models.TASK_VERIFIED_STATUS).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

type JobService struct {
	*gorm.DB
}

func (jobServ JobService) SaveJobEntity(job *models.JobEntity) (err error) {
	return jobServ.Save(job).Error
}

func (jobServ JobService) UpdateJobEntityByJobUuid(job *models.JobEntity) (err error) {
	return jobServ.Where("job_uuid=? and delete_at=?", job.JobUuid, models.UN_DELETEED_FLAG).Updates(job).Error
}

func (jobServ JobService) UpdateJobResultUrlByJobUuid(jobUuid string, resultUrl string) (err error) {
	return jobServ.Model(&models.JobEntity{}).Where("job_uuid=? and delete_at=?", jobUuid, models.UN_DELETEED_FLAG).Update("result_url", resultUrl).Error
}

func (jobServ JobService) GetJobEntityByTaskUuid(taskUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := jobServ.Model(&models.JobEntity{}).Where("task_uuid=? and delete_at=?", taskUuid, models.UN_DELETEED_FLAG).Find(&job).Error
	return job, err
}

func (jobServ JobService) GetJobEntityBySpaceUuid(spaceUuid string) int64 {
	var count int64
	jobServ.Model(&models.JobEntity{}).Where("space_uuid=? and delete_at=?", spaceUuid, models.UN_DELETEED_FLAG).Count(&count)
	return count
}

func (jobServ JobService) GetJobEntityByJobUuid(jobUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := jobServ.Model(&models.JobEntity{}).Where("job_uuid=?", jobUuid).Find(&job).Error
	return job, err
}

func (jobServ JobService) DeleteJobEntityByJobUuId(jobUuid string, jobStatus int) error {
	return jobServ.Model(&models.JobEntity{}).Where("job_uuid=? and delete_at=?", jobUuid, models.UN_DELETEED_FLAG).Updates(map[string]interface{}{
		"delete_at":  models.DELETED_FLAG,
		"status":     jobStatus,
		"pod_status": models.POD_DELETE_STATUS,
	}).Error
}

func (jobServ JobService) DeleteJobEntityBySpaceUuId(spaceUuid, jobUuid string, jobStatus int) error {
	return jobServ.Model(&models.JobEntity{}).Where("job_uuid=? and space_uuid=? and delete_at=?", jobUuid, spaceUuid, models.UN_DELETEED_FLAG).Updates(map[string]interface{}{
		"delete_at":  models.DELETED_FLAG,
		"status":     jobStatus,
		"pod_status": models.POD_DELETE_STATUS,
	}).Error
}

func (jobServ JobService) GetJobList(status int, tailNum int) (list []*models.JobEntity, err error) {
	if tailNum > 0 {
		if status >= 0 {
			err = jobServ.Model(&models.JobEntity{}).Where("delete_at=?", status).Order("create_time desc").Limit(tailNum).Find(&list).Error
		} else {
			err = jobServ.Model(&models.JobEntity{}).Order("create_time desc").Limit(tailNum).Find(&list).Error
		}
	} else {
		if status >= 0 {
			err = jobServ.Model(&models.JobEntity{}).Where("delete_at=?", status).Find(&list).Error
		} else {
			err = jobServ.Model(&models.JobEntity{}).Find(&list).Error
		}
	}
	if err != nil {
		return nil, err
	}

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return
}

func (jobServ JobService) UpdateJobReward(taskUuid string, amount string) (err error) {
	return jobServ.Model(&models.JobEntity{}).Where("task_uuid=?", taskUuid).Update("reward", amount).Error
}

func (jobServ JobService) UpdateJobScannedBlock(taskUuid string, end uint64) (err error) {
	return jobServ.Model(&models.JobEntity{}).Where(fmt.Sprintf("task_uuid='%s'", taskUuid)).Update("scanned_block", end).Error
}

func (jobServ JobService) GetJobListByNoReward() (list []*models.JobEntity, err error) {
	err = jobServ.Model(&models.JobEntity{}).Where("status in ? and (reward is null or reward ='')", []int{models.JOB_COMPLETED_STATUS, models.TERMINATED}).Find(&list).Error
	return
}

type CpInfoService struct {
	*gorm.DB
}

func (cpServ CpInfoService) GetCpInfoEntityByAccountAddress(accountAddress string) (*models.CpInfoEntity, error) {
	var cp models.CpInfoEntity
	err := cpServ.Model(&models.CpInfoEntity{}).Where("contract_address=?", accountAddress).Find(&cp).Error
	return &cp, err
}

func (cpServ CpInfoService) SaveCpInfoEntity(cp *models.CpInfoEntity) (err error) {
	cpServ.Model(&models.CpInfoEntity{}).Where("contract_address =?", cp.ContractAddress).Delete(&models.CpInfoEntity{})
	return cpServ.Save(cp).Error
}

func (cpServ CpInfoService) UpdateCpInfoByNodeId(cp *models.CpInfoEntity) (err error) {
	return cpServ.Model(&models.CpInfoEntity{}).Where("node_id =?", cp.NodeId).Updates(cp).Error
}

type EcpJobService struct {
	*gorm.DB
}

func (cpServ EcpJobService) GetEcpJobByUuid(uuid string) (*models.EcpJobEntity, error) {
	var job models.EcpJobEntity
	err := cpServ.Model(&models.EcpJobEntity{}).Where("uuid=?", uuid).Limit(1).Find(&job).Error
	return &job, err
}

func (cpServ EcpJobService) GetEcpJobs(jobUuid string) ([]models.EcpJobEntity, error) {
	var job []models.EcpJobEntity
	var err error
	if jobUuid != "" {
		err = cpServ.Model(&models.EcpJobEntity{}).Where("uuid=?", jobUuid).Find(&job).Error
	} else {
		err = cpServ.Model(&models.EcpJobEntity{}).Find(&job).Error
	}
	return job, err
}

func (cpServ EcpJobService) GetEcpJobsByLimit(tailNum int) ([]models.EcpJobEntity, error) {
	var jobList []models.EcpJobEntity
	var err error
	if tailNum > 0 {
		err = cpServ.Model(&models.EcpJobEntity{}).Order("create_time desc").Limit(tailNum).Find(&jobList).Error
	} else {
		err = cpServ.Model(&models.EcpJobEntity{}).Order("create_time desc").Find(&jobList).Error
	}
	if err != nil {
		return nil, err
	}

	for i, j := 0, len(jobList)-1; i < j; i, j = i+1, j-1 {
		jobList[i], jobList[j] = jobList[j], jobList[i]
	}

	return jobList, err
}

func (cpServ EcpJobService) GetEcpJobList(status []string) ([]models.EcpJobEntity, error) {
	var jobs []models.EcpJobEntity
	err := cpServ.Model(&models.EcpJobEntity{}).Where("status in ?", status).Find(&jobs).Error
	return jobs, err
}

func (cpServ EcpJobService) UpdateEcpJobEntity(jobUuid, status string) (err error) {
	return cpServ.Model(&models.EcpJobEntity{}).Where("uuid =?", jobUuid).Update("status", status).Error
}

func (cpServ EcpJobService) UpdateEcpJobEntityContainerName(jobUuid string, containerName string) (err error) {
	return cpServ.Model(&models.EcpJobEntity{}).Where("uuid =?", jobUuid).Update("container_name", containerName).Error
}

func (cpServ EcpJobService) UpdateEcpJobEntityPortsAndServiceUrl(jobUuid, portMap, serviceUrl string) (err error) {
	return cpServ.Model(&models.EcpJobEntity{}).Where("uuid =?", jobUuid).Updates(map[string]interface{}{
		"service_url": serviceUrl,
		"port_map":    portMap,
	}).Error
}

func (cpServ EcpJobService) UpdateEcpJobEntityMessage(jobUuid string, message string) (err error) {
	return cpServ.Model(&models.EcpJobEntity{}).Where("uuid =?", jobUuid).Updates(map[string]interface{}{
		"message":   message,
		"status":    models.TerminatedStatus,
		"delete_at": models.DELETED_FLAG,
	}).Error
}

func (cpServ EcpJobService) UpdateEcpJobEntityRewardAndBlock(jobUuid string, blockNumber int64, reward float64) (err error) {
	return cpServ.Model(&models.EcpJobEntity{}).Where("uuid =?", jobUuid).Updates(map[string]interface{}{
		"reward":            reward,
		"last_block_number": blockNumber,
	}).Error
}

func (cpServ EcpJobService) SaveEcpJobEntity(job *models.EcpJobEntity) (err error) {
	return cpServ.Save(job).Error
}

func (cpServ EcpJobService) DeleteContainerByUuid(uuid string) (err error) {
	return cpServ.Model(&models.EcpJobEntity{}).Where("uuid =?", uuid).Updates(map[string]interface{}{
		"status":    models.TerminatedStatus,
		"delete_at": models.DELETED_FLAG,
	}).Error
}

type CpBalanceService struct {
	*gorm.DB
}

func (cpServ CpBalanceService) SaveCpBalance(cpBalance models.CpBalanceEntity) (err error) {
	return cpServ.Model(&models.CpBalanceEntity{}).Save(&cpBalance).Error
}

func (cpServ CpBalanceService) GetCpBalance(cpAccount string) (*models.CpBalanceEntity, error) {
	var cpBalance models.CpBalanceEntity
	err := cpServ.Model(&models.CpBalanceEntity{}).Where("cp_account=?", cpAccount).Order("id desc").Limit(1).Find(&cpBalance).Error
	return &cpBalance, err
}

func (cpServ CpBalanceService) UpdateCpBalance(cpBalance models.CpBalanceEntity) error {
	err := cpServ.Model(&models.CpBalanceEntity{}).Where("cp_account=? and id=?", cpBalance.CpAccount, cpBalance.Id).Updates(map[string]interface{}{
		"worker_balance":    cpBalance.WorkerBalance,
		"sequencer_balance": cpBalance.SequencerBalance,
	}).Error
	return err
}

var taskSet = wire.NewSet(db.NewDbService, wire.Struct(new(TaskService), "*"))
var jobSet = wire.NewSet(db.NewDbService, wire.Struct(new(JobService), "*"))
var cpInfoSet = wire.NewSet(db.NewDbService, wire.Struct(new(CpInfoService), "*"))
var ecpJobSet = wire.NewSet(db.NewDbService, wire.Struct(new(EcpJobService), "*"))
var cpBalanceSet = wire.NewSet(db.NewDbService, wire.Struct(new(CpBalanceService), "*"))
