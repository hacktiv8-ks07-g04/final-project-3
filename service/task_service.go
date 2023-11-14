package service

import (
	"net/http"

	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/helpers"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/task_repository"
)

type TaskService interface {
	CreateNewTask(userId uint, payload *dto.NewTaskRequest) (*dto.NewTaskDataResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo task_repository.Repository
}

func NewTaskService(taskRepo task_repository.Repository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (t *taskService) CreateNewTask(userId uint, payload *dto.NewTaskRequest) (*dto.NewTaskDataResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	task := entity.Task{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      false,
		UserID:      userId,
		CategoryID:  payload.CategoryID,
	}

	err = t.taskRepo.CreateNewTask(&task)
	if err != nil {
		return nil, err
	}

	response := dto.NewTaskDataResponse{
		Status:  http.StatusCreated,
		Message: "Success create new task",
		Data: dto.NewTaskResponse{
			ID:          task.TaskID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			UserID:      task.UserID,
			CategoryID:  task.CategoryID,
			CreatedAt:   task.CreatedAt,
		},
	}

	return &response, nil
}
