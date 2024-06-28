package repository

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	*gorm.DB
}

func (t Task) GetAllTask(tailNum int) (list []*models.TaskEntity, err error) {
	if tailNum > 0 {
		err = t.Model(&models.TaskEntity{}).Order("create_time desc").Limit(tailNum).Find(&list).Error
		if err != nil {
			return nil, err
		}
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	} else {
		err = t.Model(&models.TaskEntity{}).Order("create_time").Find(&list).Error
		if err != nil {
			return nil, err
		}
	}
	return
}

func (t Task) GetTaskList(taskStatus, tailNum int) (list []*models.TaskEntity, err error) {
	if tailNum > 0 {
		err = t.Where(&models.TaskEntity{Status: taskStatus}).Order("create_time desc").Limit(tailNum).Find(&list).Error
		if err != nil {
			return nil, err
		}
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	} else {
		err = t.Where(&models.TaskEntity{Status: taskStatus}).Order("create_time").Find(&list).Error
		if err != nil {
			return nil, err
		}
	}
	return
}

func (t Task) SaveTaskEntity(task *models.TaskEntity) (err error) {
	if task.Status == models.TASK_FAILED_STATUS || task.Status == models.TASK_SUCCESS_STATUS {
		task.EndTime = time.Now().Unix()
	}
	return t.Save(task).Error
}

func (t Task) GetTaskEntity(taskId int64) (*models.TaskEntity, error) {
	var taskEntity models.TaskEntity
	err := t.First(&taskEntity, taskId).Error
	return &taskEntity, err
}

func (t Task) GetTaskListNoReward() (list []*models.TaskEntity, err error) {
	err = t.Where("status=? and reward_status=?", models.TASK_SUCCESS_STATUS, models.REWARD_UNCLAIMED).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return
}

var taskRepository = wire.NewSet(db.NewDbService, wire.Struct(new(Task), "*"))
