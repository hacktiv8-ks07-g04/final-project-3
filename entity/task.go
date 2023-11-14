package entity

import "time"

type Task struct {
	TaskID      uint   `gorm:"primaryKey;not null;" json:"id"`
	Title       string `gorm:"type:varchar(255);not null" json:"title" valid:"required~title is required, type(string)"`
	Description string `gorm:"type:varchar(255);not null" json:"description" valid:"required~description is required, type(string)"`
	Status      bool   `gorm:"type:boolean;not null" json:"status" valid:"required~status is required, type(boolean)"`
	UserID      uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id" valid:"required~user_id is required, type(uint)"`
	CategoryID  uint   `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category_id" valid:"required~category_id is required, type(uint)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
