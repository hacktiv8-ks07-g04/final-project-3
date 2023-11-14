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
