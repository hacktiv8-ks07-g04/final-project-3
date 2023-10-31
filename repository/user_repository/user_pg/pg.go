package user_pg

import (
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterNewUser(user entity.User) (entity.User, error)
}

type userPG struct {
	db *gorm.DB
}

func UserInit(db *gorm.DB) UserRepository {
	return &userPG{db}
}

func (u *userPG) RegisterNewUser(user entity.User) (entity.User, error) {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})

	return user, err
}
