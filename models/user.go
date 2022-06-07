package models

import (
	"errors"
	"hacktiv8-day12/helper"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Username     string        `gorm:"unique" json:"username" form:"username" valid:"required~your username is required"`
	Email        string        `gorm:"unique" json:"email" form:"email" valid:"required~your email is required, email~invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~your password is required,minstringlength(6)~password has to have a minimum length of 6 characters"`
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
	Photos       []Photo       `json:"Photos,omitempty"`
	Comments     []Comment     `json:"Comments,omitempty"`
	SocialMedias []SocialMedia `json:"SocialMedias,omitempty"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"udated_at,omitempty"`
	DeletedAt    *time.Time    `sql:"index" json:"deleted_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	} else if u.Age < 8 || u.Age > 85 {
		err = errors.New("age must greater than 8 and smaller than 86")
		return
	}
	u.Password, _ = helper.HashPass(u.Password)

	err = nil
	return
}
