package entity

import (
	"log"
	"time"

	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    uint   `gorm:"primaryKey" json:"id"`
	FullName  string `gorm:"type:varchar(100);not null" json:"full_name"`
	Email     string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
	Role      string `gorm:"type:varchar(100);not null" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword() errs.MessageErr {
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)
	if err != nil {
		log.Println(err)
		return errs.NewInternalServerError("Error hashing password")
	}

	u.Password = string(bs)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
