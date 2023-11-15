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
	GetTaskWithUser() (*dto.TaskListResponse, errs.MessageErr)
	// UpdateTaskById(id uint, userId uint, payload *dto.UpdateDetailTaskRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr)
	UpdateTaskById(id uint, userId uint, payload *dto.UpdateDetailTaskRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr)
	UpdateTaskStatus(id uint, userId uint, payload *dto.UpdateTaskStatusRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr)
	// UpdateTaskCategory(id uint, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr)
	UpdateTaskCategory(id uint, userId uint, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr)
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
			ID:          task.ID,
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

func (t *taskService) GetTaskWithUser() (*dto.TaskListResponse, errs.MessageErr) {
	allTasks, err := t.taskRepo.GetTaskWithUser()
	if err != nil {
		return nil, errs.NewInternalServerError("Error when trying to get data")
	}

	userTask := []dto.UserTask{}
	for _, v := range allTasks {
		task := dto.UserTask{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Status:      v.Status,
			UserID:      v.UserID,
			CategoryID:  v.CategoryID,
			CreatedAt:   v.CreatedAt,
			User: dto.User{
				ID:       v.User.ID,
				Email:    v.User.Email,
				FullName: v.User.FullName,
			},
		}
		userTask = append(userTask, task)
	}

	response := dto.TaskListResponse{
		StatusCode: http.StatusOK,
		Message:    "Success get all tasks",
		Data:       userTask,
	}

	return &response, nil
}

// update task title and description
func (t *taskService) UpdateTaskById(id uint, userId uint, payload *dto.UpdateDetailTaskRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	updateTask := entity.Task{
		Title:       payload.Title,
		Description: payload.Description,
	}

	result, err := t.taskRepo.UpdateTaskTitleAndDescription(id, userId, &updateTask)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to edit payload")
	}

	response := dto.UpdateTaskDetailResponse{
		StatusCode: http.StatusOK,
		Message:    "Success update task title and description field",
		Data: dto.UpdateDetailTaskData{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Status:      result.Status,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			UpdatedAt:   result.UpdatedAt,
		},
	}

	return &response, nil
}

// update status of a task
func (t *taskService) UpdateTaskStatus(id uint, userId uint, payload *dto.UpdateTaskStatusRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	updateStatus := entity.Task{
		Status: payload.Status,
	}

	result, err := t.taskRepo.UpdateTaskStatus(id, userId, &updateStatus)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to edit payload")
	}
	response := dto.UpdateTaskDetailResponse{
		StatusCode: http.StatusOK,
		Message:    "Success update task status field",
		Data: dto.UpdateDetailTaskData{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Status:      result.Status,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			UpdatedAt:   result.UpdatedAt,
		},
	}

	return &response, nil
}

func (t *taskService) UpdateTaskCategory(id uint, userId uint, payload *dto.UpdateTaskCategoryRequest) (*dto.UpdateTaskDetailResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	updateCategory := entity.Task{
		CategoryID: payload.CategoryID,
	}

	result, err := t.taskRepo.UpdateTaskCategory(id, userId, &updateCategory)
	if err != nil {
		return nil, errs.NewInternalServerError("Error occurred while trying to edit payload")
	}
	response := dto.UpdateTaskDetailResponse{
		StatusCode: http.StatusOK,
		Message:    "Success update task category field",
		Data: dto.UpdateDetailTaskData{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Status:      result.Status,
			UserID:      result.UserID,
			CategoryID:  result.CategoryID,
			UpdatedAt:   result.UpdatedAt,
		},
	}

	return &response, nil
}
