package models

import (
	"time"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required~title is required"`
	Caption   string `json:"caption"`
	PhotoUrl  string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo url is required"`
	UserID    uint   `json:"user_id"`
	User      User
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
