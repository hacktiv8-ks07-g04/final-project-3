package dto

import (
	"time"
)

type Category struct {
	CategoryID uint      `json:"id"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Tasks      []Task    `json:"Tasks"`
}

type NewCategoryRequest struct {
	Type string `json:"type" valid:"required~type is required, type(string)"`
}

type NewCategoryResponse struct {
	StatusCode int                     `json:"status"`
	Message    string                  `json:"message"`
	Data       NewCategoryDataResponse `json:"data"`
}

type NewCategoryDataResponse struct {
	CategoryID uint      `json:"id"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
}
