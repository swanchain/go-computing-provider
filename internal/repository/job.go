package repository

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/gorm"
)

type Job struct {
	*gorm.DB
}

func (j Job) SaveJobEntity(job *models.JobEntity) (err error) {
	return j.Save(job).Error
}

func (j Job) UpdateJobEntityBySpaceUuid(job *models.JobEntity) (err error) {
	return j.Where("space_uuid=?", job.SpaceUuid).Updates(job).Error
}

func (j Job) UpdateJobEntityByJobUuid(job *models.JobEntity) (err error) {
	return j.Where("job_uuid=?", job.JobUuid).Updates(job).Error
}

func (j Job) UpdateJobResultUrlByJobUuid(jobUuid string, resultUrl string) (err error) {
	return j.Model(&models.JobEntity{}).Where("job_uuid=?", jobUuid).Update("result_url", resultUrl).Error
}

func (j Job) GetJobEntityByTaskUuid(taskUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := j.Where("task_uuid=?", taskUuid).Find(&job).Error
	return job, err
}

func (j Job) GetJobEntityBySpaceUuid(spaceUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := j.Where("space_uuid=?", spaceUuid).Find(&job).Error
	return job, err
}

func (j Job) GetJobEntityByJobUuid(jobUuid string) (models.JobEntity, error) {
	var job models.JobEntity
	err := j.Where("job_uuid=?", jobUuid).Find(&job).Error
	return job, err
}

func (j Job) DeleteJobEntityBySpaceUuId(spaceUuid string) error {
	return j.Where("space_uuid=?", spaceUuid).Delete(&models.JobEntity{}).Error
}

func (j Job) GetJobList() (list []*models.JobEntity, err error) {
	err = j.Model(&models.JobEntity{}).Find(&list).Error
	return
}

func (j Job) DeleteJobs(spaceIds []string) (err error) {
	return j.Where("space_uuid in ?", spaceIds).Delete(&models.JobEntity{}).Error
}

var jobRepository = wire.NewSet(db.NewDbService, wire.Struct(new(Job), "*"))
