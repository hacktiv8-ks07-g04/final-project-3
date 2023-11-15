package task_repository

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

type Repository interface {
	CreateNewTask(task *entity.Task) errs.MessageErr
	GetTaskWithUser() ([]entity.Task, errs.MessageErr)
	GetTaskById(id uint) (*entity.Task, errs.MessageErr)
	UpdateTaskStatus(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskCategory(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskTitleAndDescription(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	DeleteTaskById(id uint, userId uint) errs.MessageErr
}
