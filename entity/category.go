package entity

import "time"

type Category struct {
	CategoryID uint   `gorm:"primaryKey;not null" json:"id"`
	Type       string `gorm:"type:varchar(255);not null" json:"type" valid:"required~type is required, type(string)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Tasks      []Task `gorm:"foreignKey:TaskID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
