package params

import (
	"time"
)

type GetAllPhotos struct {
	ID        uint       `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Caption   string     `json:"caption,omitempty"`
	PhotoUrl  string     `json:"photo_url,omitempty"`
	UserID    uint       `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	User      *User      `json:"user,omitempty"`
}

type User struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type GetAllComments struct {
	ID        uint          `json:"id,omitempty"`
	Message   string        `json:"message,omitempty"`
	PhotoID   uint          `json:"photo_id,omitempty"`
	UserID    uint          `json:"user_id,omitempty"`
	CreatedAt *time.Time    `json:"created_at,omitempty"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
	User      *User         `json:"user"`
	Photo     *GetAllPhotos `json:"photos"`
}

type GetAllSocialMedias struct {
	ID             uint       `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	SocialMediaUrl string     `json:"social_media_url,omitempty"`
	UserID         uint       `json:"user_id,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	User           *User      `json:"user"`
}
