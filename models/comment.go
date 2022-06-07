package models

import (
	"time"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"user_id" form:"user_id" valid:"required~user_id is required"`
	PhotoID   uint   `json:"photo_id" form:"photo_id" valid:"required~photo_id is required"`
	Message   string `gorm:"not null" json:"message" form:"message" valid:"required~message is required"`
	User      User
	Photo     Photo
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
