package computing

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/gorm"
	"time"
)

type TaskService struct {
	*gorm.DB
}

func (taskServ TaskService) GetAllTask(tailNum int) (list []*models.TaskEntity, err error) {
	if tailNum > 0 {
		err = taskServ.Model(&models.TaskEntity{}).Order("create_time desc").Limit(tailNum).Find(&list).Error
		if err != nil {
			return nil, err
		}
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	} else {
		err = taskServ.Model(&models.TaskEntity{}).Order("create_time").Find(&list).Error
		if err != nil {
			return nil, err
		}
	}
	return
}

func (taskServ TaskService) GetTaskList(taskStatus, tailNum int) (list []*models.TaskEntity, err error) {
	if tailNum > 0 {
		err = taskServ.Where(&models.TaskEntity{Status: taskStatus}).Order("create_time desc").Limit(tailNum).Find(&list).Error
		if err != nil {
			return nil, err
		}
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	} else {
		err = taskServ.Where(&models.TaskEntity{Status: taskStatus}).Order("create_time").Find(&list).Error
		if err != nil {
			return nil, err
		}
	}
	return
}

func (taskServ TaskService) SaveTaskEntity(task *models.TaskEntity) (err error) {
	if task.Status == models.TASK_FAILED_STATUS || task.Status == models.TASK_SUCCESS_STATUS {
		task.EndTime = time.Now().Unix()
	}
	return taskServ.Save(task).Error
}

func (taskServ TaskService) GetTaskEntity(taskId int64) (*models.TaskEntity, error) {
	var taskEntity models.TaskEntity
	err := taskServ.First(&taskEntity, taskId).Error
	return &taskEntity, err
}

func (taskServ TaskService) GetTaskListNoReward() (list []*models.TaskEntity, err error) {
	err = taskServ.Where("status=? and reward_status=?", models.TASK_SUCCESS_STATUS, models.REWARD_UNCLAIMED).Find(&list).Error
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

func (jobServ JobService) UpdateJobEntityBySpaceUuid(job *models.JobEntity) (err error) {
	return jobServ.Where("space_uuid=?", job.SpaceUuid).Updates(job).Error
}

func (jobServ JobService) UpdateJobEntityByJobUuid(job *models.JobEntity) (err error) {
	return jobServ.Where("job_uuid=?", job.JobUuid).Updates(job).Error
}

func (jobServ JobService) UpdateJobResultUrlByJobUuid(jobUuid string, resultUrl string) (err error) {
	return jobServ.Model(&models.JobEntity{}).Where("job_uuid=?", jobUuid).Update("result_url", resultUrl).Error
}

func (jobServ JobService) GetJobEntityByTaskUuid(taskUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := jobServ.Where("task_uuid=?", taskUuid).Find(&job).Error
	return job, err
}

func (jobServ JobService) GetJobEntityBySpaceUuid(spaceUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := jobServ.Where("space_uuid=?", spaceUuid).Find(&job).Error
	return job, err
}

func (jobServ JobService) GetJobEntityByJobUuid(jobUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := jobServ.Where("job_uuid=?", jobUuid).Find(&job).Error
	return job, err
}

func (jobServ JobService) DeleteJobEntityBySpaceUuId(spaceUuid string) error {
	return jobServ.Where("space_uuid=?", spaceUuid).Delete(&models.JobEntity{}).Error
}

func (jobServ JobService) GetJobList() (list []*models.JobEntity, err error) {
	err = jobServ.Model(&models.JobEntity{}).Find(&list).Error
	return
}

func (jobServ JobService) DeleteJobs(spaceIds []string) (err error) {
	return jobServ.Where("space_uuid in ?", spaceIds).Delete(&models.JobEntity{}).Error
}

type CpInfoService struct {
	*gorm.DB
}

func (cpServ CpInfoService) GetCpInfoEntityByAccountAddress(accountAddress string) (*models.CpInfoEntity, error) {
	var cp models.CpInfoEntity
	err := cpServ.Where("contract_address=?", accountAddress).Find(&cp).Error
	return &cp, err
}

func (cpServ CpInfoService) SaveCpInfoEntity(cp *models.CpInfoEntity) (err error) {
	cpServ.Where("contract_address =?", cp.ContractAddress).Delete(&models.CpInfoEntity{})
	return cpServ.Save(cp).Error
}

func (cpServ CpInfoService) UpdateCpInfoByNodeId(cp *models.CpInfoEntity) (err error) {
	return cpServ.Where("node_id =?", cp.NodeId).Updates(cp).Error
}

var taskSet = wire.NewSet(db.NewDbService, wire.Struct(new(TaskService), "*"))
var jobSet = wire.NewSet(db.NewDbService, wire.Struct(new(JobService), "*"))
var cpInfoSet = wire.NewSet(db.NewDbService, wire.Struct(new(CpInfoService), "*"))
