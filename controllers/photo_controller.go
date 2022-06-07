package controllers

import (
	"fmt"
	"hacktiv8-day12/params"
	"hacktiv8-day12/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService services.PhotoService
}

func NewPhotoController(service *services.PhotoService) *PhotoController {
	return &PhotoController{photoService: *service}
}

//Post photo
func (p *PhotoController) PostPhoto(c *gin.Context) {
	var photo params.Photo
	idUserToken := c.MustGet("id")
	idUser := int(idUserToken.(float64))

	err := c.ShouldBindJSON(&photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	response := p.photoService.PostPhoto(uint(idUser), &photo)
	c.JSON(response.Status, response)
}

//Get all photos
func (p *PhotoController) GetAllPhotos(c *gin.Context) {
	response, responseErr := p.photoService.GetAllPhotos()
	if responseErr != nil {
		c.JSON(responseErr.Status, responseErr)
	}
	c.JSON(http.StatusOK, response)
}

//Update photo
func (p *PhotoController) UpdatePhoto(c *gin.Context) {
	var photoUpdate *params.Photo
	photoId := c.Param("photoId")
	idUserToken := c.MustGet("id")
	emailUserToken := c.MustGet("email")
	idUser := int(idUserToken.(float64))
	email := fmt.Sprintf("%v", emailUserToken)

	err := c.ShouldBindJSON(&photoUpdate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	id, err := strconv.Atoi(photoId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		})
		return
	}
	// var userDB *models.User
	_, err = p.photoService.CheckUser(email)
	log.Default().Println("error controller", err) //Println()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}

	response := p.photoService.UpdatePhoto(uint(idUser), uint(id), *photoUpdate)
	c.JSON(response.Status, response)
}

func (p *PhotoController) DeletePhoto(c *gin.Context) {
	photoId := c.Param("photoId")
	idUserToken := c.MustGet("id")
	idUser := int(idUserToken.(float64))
	emailUserToken := c.MustGet("email")
	email := fmt.Sprintf("%v", emailUserToken)
	idUint, err1 := strconv.Atoi(photoId)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err1.Error(),
		})
		return
	}
	// var userDB *models.User
	_, err := p.photoService.CheckUser(email)
	log.Default().Println("error controller", err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	userPhotoId, err := p.photoService.GetPhotoUserId(uint(idUint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	if userPhotoId != uint(idUser) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	response := p.photoService.DeletePhoto(uint(idUint))
	c.JSON(response.Status, response)
}
