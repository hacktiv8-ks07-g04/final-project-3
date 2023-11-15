package dto

import "time"

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserTask struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	User        User      `json:"User"`
}

type NewTaskRequest struct {
	Title       string `json:"title" valid:"required~title is required, type(string)"`
	Description string `json:"description" valid:"required~description is required, type(string)"`
	CategoryID  uint   `json:"category_id" valid:"required~category_id is required, type(uint)"`
}

type NewTaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type NewTaskDataResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    NewTaskResponse `json:"data"`
}

type TaskListResponse struct {
	StatusCode int        `json:"status"`
	Message    string     `json:"message"`
	Data       []UserTask `json:"data"`
}

type UpdateDetailTaskRequest struct {
	Title       string `json:"title" valid:"required~title is required, type(string)"`
	Description string `json:"description" valid:"required~description is required, type(string)"`
}

type UpdateDetailTaskData struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTaskDetailResponse struct {
	StatusCode int                  `json:"status"`
	Message    string               `json:"message"`
	Data       UpdateDetailTaskData `json:"data"`
}

type UpdateTaskStatusRequest struct {
	Status bool `json:"status" valid:"required~status is required, type(boolean)"`
}

type UpdateTaskCategoryRequest struct {
	CategoryID uint `json:"category_id" valid:"required~category_id is required, type(uint)"`
}
