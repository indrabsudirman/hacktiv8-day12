package models

import (
	"time"
)

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~social media url is required"`
	UserID         uint   `json:"user_id"`
	User           User
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	UpdatedAt      time.Time  `json:"udated_at,omitempty"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
