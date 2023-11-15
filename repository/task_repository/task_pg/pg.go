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

// update task title and description
func (t *taskPg) UpdateTaskById(id uint, task *entity.Task) errs.MessageErr {
	if err := t.db.Model(&task).Where("id = ?", id).Updates(task).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}
