package user_pg

import (
	"fmt"
	"log"

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

func (u *userPG) RegisterNewUser(user *entity.User) errs.MessageErr {
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

// update user
func (u *userPG) UpdateUser(user entity.User) errs.MessageErr {
	if err := u.db.Save(&user).Error; err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}

// get user by id
func (u *userPG) GetUserById(id uint) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, errs.NewNotFoundError(fmt.Sprintf("User with id %d is not found", id))
	}

	return &user, nil
}

// delete user
func (u *userPG) DeleteUserById(id uint) errs.MessageErr {
	user, err := u.GetUserById(id)
	if err != nil {
		return err
	}

	if err := u.db.Delete(user).Error; err != nil {
		log.Println(err.Error())
		return errs.NewInternalServerError(fmt.Sprintf("Failed to delete user with id %d", user.ID))
	}

	return nil
}
