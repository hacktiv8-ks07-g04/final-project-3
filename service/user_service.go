package service

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/dto"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/user_repository"
)

type UserService interface {
	RegisterNewUser(user entity.User) (*dto.RegisterResponse, error)
}

type userService struct {
	userRepo user_repository.Repository
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterNewUser(user entity.User) (*dto.RegisterResponse, error) {
	user, err := s.userRepo.RegisterNewUser(user)
	if err != nil {
		return nil, err
	}

	response := dto.RegisterResponse{
		StatusCode: 200,
		Message:    "Successfully registered new user",
		Data: dto.UserDataResponse{
			ID:        int(user.UserID),
			FullName:  user.FullName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
	}
	return &response, nil
}
