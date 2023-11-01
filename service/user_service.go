package service

import (
	"net/http"

	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/helpers"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/user_repository"
)

type UserService interface {
	CreateNewUser(payload dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr)
	LoginUser(newUserRequest dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.Repository
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) CreateNewUser(payload dto.RegisterRequest) (*dto.RegisterResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Password: payload.Password,
	}

	err = user.HashPassword()
	if err != nil {
		return nil, err
	}

	err = u.userRepo.RegisterNewUser(user)
	if err != nil {
		return nil, err
	}

	response := dto.RegisterResponse{
		StatusCode: 200,
		Message:    "Successfully registered new user",
		Data: dto.UserDataResponse{
			UserID:    user.UserID,
			FullName:  user.FullName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	}

	return &response, nil
}

// Login
func (u *userService) LoginUser(newUserRequest dto.LoginRequest) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newUserRequest)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(newUserRequest.Email)
	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequest("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(newUserRequest.Password)
	if !isValidPassword {
		return nil, errs.NewBadRequest("invalid email/password")
	}

	token := user.GenerateToken()

	response := dto.LoginResponse{
		StatusCode: http.StatusOK,
		Message:    "successfully logged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}

	return &response, nil
}
