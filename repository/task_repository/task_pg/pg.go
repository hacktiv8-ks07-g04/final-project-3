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

// ISSUE: GetTaskById returns InternalServerError (500) when a task is not found.
// It should return NotFoundError (404) for that case. It's also exposed on the
// public Repository interface but never called by any service/handler directly —
// it's only used internally by the update/delete methods, so it shouldn't be exported.
func (t *taskPg) GetTaskById(id uint) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	if err := t.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &task, nil
}

// ISSUE (DUPLICATION): Every update/delete method below repeats the same pattern:
// (1) GetTaskById, (2) check task.UserID != userId, (3) return UnauthorizedError.
// This ownership check is also business logic that belongs in the SERVICE or a
// middleware layer, not in the data-access (repository) layer. Extract to a shared
// helper (e.g. getOwnedTask) and move authorization up the stack.
func (t *taskPg) UpdateTaskTitleAndDescription(id uint, userId uint, taskPayload *entity.Task) (*entity.Task, errs.MessageErr) {
	task, err := t.GetTaskById(id)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to find task")
	}

	if task.UserID != userId {
		// ISSUE: Authorization in repository layer (should be in service/middleware).
		// Also the task-not-found case above returns 500 from GetTaskById instead of 404.
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

func (t *taskPg) DeleteTaskById(id uint, userId uint) errs.MessageErr {
	task, err := t.GetTaskById(id)

	if err != nil {
		return errs.NewInternalServerError("Error occurred while trying to find task")
	}

	if task.UserID != userId {
		return errs.NewUnauthorizedError("You are not authorized to delete this task")
	}

	if err := t.db.Delete(&task).Error; err != nil {
		return errs.NewInternalServerError("Error occurred while trying to delete task")
	}

	return nil
}
