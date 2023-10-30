package entity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FullName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
