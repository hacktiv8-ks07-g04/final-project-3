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

type CategoryListResponse struct {
	StatusCode int        `json:"status"`
	Message    string     `json:"message"`
	Data       []Category `json:"data"`
}

type UpdateCategoryRequest struct {
	Type string `json:"type" valid:"required~type is required, type(string)"`
}

type UpdateCategoryDataResponse struct {
	CategoryID uint      `json:"id"`
	Type       string    `json:"type"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateCategoryResponse struct {
	StatusCode int                        `json:"status"`
	Message    string                     `json:"message"`
	Data       UpdateCategoryDataResponse `json:"data"`
}

type DeleteCategoryResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
