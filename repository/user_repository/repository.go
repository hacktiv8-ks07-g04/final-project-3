package user_repository

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
)

type Repository interface {
	RegisterNewUser(user entity.User) errs.MessageErr
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
	UpdateUser(user entity.User) errs.MessageErr
}
