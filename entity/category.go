package entity

import "time"

type Category struct {
	CategoryID uint   `gorm:"primaryKey;not null"`
	Type       string `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
