package controllers

import (
	"hacktiv8-day12/params"
	"hacktiv8-day12/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{*service}
}

//Register User
func (u *UserController) RegisterUser(c *gin.Context) {
	var reqister params.Register

	err := c.ShouldBindJSON(&reqister)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	response := u.userService.RegisterUser(&reqister)
	c.JSON(response.Status, response)
}

//Login User
func (u *UserController) LoginUser(c *gin.Context) {
	var login params.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	response := u.userService.LoginUser(&login)
	c.JSON(response.Status, response)
}

//Update User
func (u *UserController) UpdateUser(c *gin.Context) {
	var update *params.Update
	userId := c.Param("userId")
	idUserToken := c.MustGet("id")

	idUser := int(idUserToken.(float64))

	err := c.ShouldBindJSON(&update)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		})
		return
	}
	if id != idUser {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	response := u.userService.UpdateUser(uint(id), *update)
	c.JSON(response.Status, response)
}

//Delete User
func (u *UserController) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	//Get from token JWT
	idUserToken := c.MustGet("id")
	idUser := int(idUserToken.(float64))

	idUint, err1 := strconv.Atoi(userId)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err1.Error(),
		})
		return
	}
	if idUint != idUser {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}

	response := u.userService.DeleteUser(uint(idUint))
	c.JSON(response.Status, response)

}
