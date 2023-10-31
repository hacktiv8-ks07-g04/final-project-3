package entity

import "time"

type User struct {
	UserID        uint   `gorm:"primaryKey" json:"id"`
	FullName  string `gorm:"type:varchar(100);not null" json:"full_name"`
	Email     string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
	Role      string `gorm:"type:varchar(100);not null" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
