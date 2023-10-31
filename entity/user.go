package entity

import (
	"log"
	"time"

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

func (u *User) HashPassword() string {
	salt := 8

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
