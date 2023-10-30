package entity

import "time"

type Task struct {
	ID          uint       `gorm:"primaryKey;not null"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:varchar(255);not null"`
	Status      bool       `gorm:"type:boolean;not null"`
	UserID      []User     `gorm:"foreignKey:ID;not null"`
	CategoryID  []Category `gorm:"foreignKey:ID;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
