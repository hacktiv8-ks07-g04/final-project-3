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
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
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
