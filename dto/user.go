package dto

import (
	"time"
)

type RegisterRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserDataResponse struct {
	UserID    uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterResponse struct {
	StatusCode int              `json:"status"`
	Message    string           `json:"message"`
	Data       UserDataResponse `json:"data"`
}
