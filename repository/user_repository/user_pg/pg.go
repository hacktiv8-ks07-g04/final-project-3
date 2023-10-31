package user_pg

import (
	"fmt"

	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"github.com/hacktiv8-ks07-g04/final-project-3/repository/user_repository"
	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func UserInit(db *gorm.DB) user_repository.Repository {
	return &userPG{
		db: db,
	}
}

func (u *userPG) RegisterNewUser(user entity.User) errs.MessageErr {
	if err := u.db.Create(&user).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}

// get user by email
func (u *userPG) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFoundError(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}
