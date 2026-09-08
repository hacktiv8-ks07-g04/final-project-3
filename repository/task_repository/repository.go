package task_repository

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

type Repository interface {
	CreateNewTask(task *entity.Task) errs.MessageErr
	GetTaskWithUser() ([]entity.Task, errs.MessageErr)
	// ISSUE: GetTaskById is exposed on the interface but never called by any service
	// or handler — it's only used internally by the update/delete methods in task_pg.
	// Consider making it private to the implementation.
	GetTaskById(id uint) (*entity.Task, errs.MessageErr)
	UpdateTaskStatus(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskCategory(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskTitleAndDescription(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr)
	DeleteTaskById(id uint, userId uint) errs.MessageErr
}
