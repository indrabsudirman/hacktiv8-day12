package repositories

import (
	"errors"
	"hacktiv8-day12/helper"
	"hacktiv8-day12/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(*models.User) error
	LoginUser(string, string) (*models.User, error)
	UpdateUser(uint, *models.User) (*models.User, error)
	DeleteUser(uint) error
	CheckUser(uint, string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

//Register User
func (u *userRepository) RegisterUser(register *models.User) error {
	err := u.db.Create(register).Error
	return err
}

//Login User
func (u *userRepository) LoginUser(email, password string) (*models.User, error) {
	var userDB models.User
	if err := u.db.Where("email=?", email).Take(&userDB).Error; err != nil {
		return nil, errors.New("invalid email")
	} else if err := helper.CheckPass(userDB.Password, password); err != nil {
		return nil, errors.New("invalid password")
	}
	return &userDB, nil
}

//Update User
func (u *userRepository) UpdateUser(id uint, update *models.User) (*models.User, error) {
	var user models.User

	err := u.db.First(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}

	err = u.db.Model(&user).Where("id=?", id).Updates(models.User{
		Username: update.Username,
		Email:    update.Email,
	}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//Delete User
func (u *userRepository) DeleteUser(id uint) error {
	var user models.User
	err := u.db.First(&user, "id=?", id).Error
	if err != nil {
		return err
	}
	err = u.db.Delete(&user, "id=?", id).Error
	return err
}

func (u *userRepository) CheckUser(id uint, email string) (*models.User, error) {
	var userDB models.User
	if err := u.db.Where("id=? email=?", id, email).Take(&userDB).Error; err != nil {
		return nil, errors.New("invalid id/email")
	}
	return &userDB, nil
}
