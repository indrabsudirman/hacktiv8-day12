package services

import (
	"hacktiv8-day12/helper"
	"hacktiv8-day12/models"
	"hacktiv8-day12/params"
	"hacktiv8-day12/repositories"
	"log"
	"net/http"
	"time"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (u *UserService) RegisterUser(requestParam *params.Register) *params.Response {
	user := models.User{
		Username: requestParam.Username,
		Email:    requestParam.Email,
		Password: requestParam.Password,
		Age:      requestParam.Age,
	}
	err := u.userRepo.RegisterUser(&user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusCreated,
		Message: "created success",
		Payload: map[string]interface{}{
			"age":      user.Age,
			"email":    user.Email,
			"id":       user.ID,
			"username": user.Username,
		},
	}
}

//Login User
func (u *UserService) LoginUser(requestLogin *params.Login) *params.Response {
	user := models.User{
		Email:    requestLogin.Email,
		Password: requestLogin.Password,
	}
	var userDB *models.User
	userDB, err := u.userRepo.LoginUser(user.Email, user.Password)
	if err != nil {
		return &params.Response{
			Status:         http.StatusUnauthorized,
			Error:          "unauthorized",
			AdditionalInfo: err.Error(),
		}
	}
	log.Default().Println("user DB data", userDB)
	token := helper.GeneratedToken(userDB.ID, userDB.Email)
	return &params.Response{
		Status:  http.StatusOK,
		Message: "login success",
		Payload: token,
	}
}

//Update User
func (u *UserService) UpdateUser(id uint, requestUpdate params.Update) *params.Response {
	var user models.User
	user.Email = requestUpdate.Email
	user.Username = requestUpdate.Username
	user.UpdatedAt = time.Now()

	response, err := u.userRepo.UpdateUser(id, &user)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "data not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status: http.StatusOK,
		Error:  "update success",
		Payload: map[string]interface{}{
			"id":         response.ID,
			"email":      response.Email,
			"username":   response.Username,
			"age":        response.Age,
			"updated_at": response.UpdatedAt,
		},
	}

}

//Delete User
func (u *UserService) DeleteUser(id uint) *params.Response {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "user not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "your account has been successfully deleted",
	}
}
