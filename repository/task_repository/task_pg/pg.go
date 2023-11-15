package task_pg

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/task_repository"
	"gorm.io/gorm"
)

type taskPg struct {
	db *gorm.DB
}

func NewTaskPg(db *gorm.DB) task_repository.Repository {
	return &taskPg{db: db}
}

func (t *taskPg) CreateNewTask(task *entity.Task) errs.MessageErr {
	if err := t.db.Create(&task).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}

// get task with all associated user
func (t *taskPg) GetTaskWithUser() ([]entity.Task, errs.MessageErr) {
	var task []entity.Task

	if err := t.db.Preload("User").Find(&task).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return task, nil
}

// get task id
func (t *taskPg) GetTaskById(id uint) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	if err := t.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &task, nil
}

func (t *taskPg) UpdateTaskTitleAndDescription(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskById(id)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find task")
	}

	if task.UserID != userId {
		return nil, errs.NewUnauthorizedError("You are not authorized to update this task")
	}

	if err := t.db.Model(task).Updates(map[string]interface{}{
		"Title":       taskPayload.Title,
		"Description": taskPayload.Description,
	}).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to update task")
	}

	return task, nil
}

func (t *taskPg) UpdateTaskStatus(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskById(id)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find task")
	}

	if task.UserID != userId {
		return nil, errs.NewUnauthorizedError("You are not authorized to update this task")
	}

	if err := t.db.Model(task).Update("Status", taskPayload.Status).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to update task")
	}

	return task, nil
}

func (t *taskPg) UpdateTaskCategory(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskById(id)

	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find task")
	}

	if task.UserID != userId {
		return nil, errs.NewUnauthorizedError("You are not authorized to update this task")
	}

	if err := t.db.Model(task).Updates(entity.Task{CategoryID: taskPayload.CategoryID}).Error; err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to update task")
	}

	return task, nil
}
