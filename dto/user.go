package dto

import (
	"time"
)

type RegisterRequest struct {
	FullName string `json:"full_name" valid:"required~full name is required, type(string)"`
	Password string `json:"password" valid:"required~password is required,minstringlength(6)~password must be at least 6 characters"`
	Email    string `json:"email" valid:"email~email is not valid, required~email is required, type(string)"`
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

type LoginRequest struct {
	Email    string `json:"email" valid:"required~full name is required, type(string)"`
	Password string `json:"password"`
}

type LoginResponse struct {
	StatusCode int           `json:"status"`
	Message    string        `json:"message"`
	Data       TokenResponse `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" valid:"required~full name is required, type(string)"`
	Email    string `json:"email" valid:"email~email is not valid, required~email is required, type(string)"`
}

type UpdateUserResponse struct {
	StatusCode int                    `json:"status"`
	Message    string                 `json:"message"`
	Data       UpdateUserDataResponse `json:"data"`
}

type UpdateUserDataResponse struct {
	UserID    uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
