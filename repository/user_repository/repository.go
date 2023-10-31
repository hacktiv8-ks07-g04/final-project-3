package user_repository

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
)

type Repository interface {
	RegisterNewUser(user entity.User) (entity.User, error)
}
