package entity

import "time"

type Task struct {
	TaskID      uint       `gorm:"primaryKey;not null"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:varchar(255);not null"`
	Status      bool       `gorm:"type:boolean;not null"`
	UserID      uint       `gorm:"not null"`
	CategoryID  uint       `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

