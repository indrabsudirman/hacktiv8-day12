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

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(service *services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{*service}
}

func (s *SocialMediaController) PostSocialMedia(c *gin.Context) {
	var sosmed params.SocialMedia

	idUserToken := c.MustGet("id")
	idUser := int(idUserToken.(float64))

	err := c.ShouldBindJSON(&sosmed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	response := s.socialMediaService.PostSocialMedia(uint(idUser), &sosmed)
	c.JSON(response.Status, response)
}

//Get all social medias
func (s *SocialMediaController) GetAllSocialMedias(c *gin.Context) {
	response, responseError := s.socialMediaService.GetAllSocialMedias()
	if responseError != nil {
		c.JSON(responseError.Status, response)
	}
	c.JSON(http.StatusOK, response)
}

func (s *SocialMediaController) UpdateSocialMedia(c *gin.Context) {
	var sosmedUpdate params.SocialMedia
	socialMediaId := c.Param("socialMediaId")
	idUserToken := c.MustGet("id")
	emailUserToken := c.MustGet("email")
	idUser := int(idUserToken.(float64))
	email := fmt.Sprintf("%v", emailUserToken)

	err := c.ShouldBindJSON(&sosmedUpdate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	id, err := strconv.Atoi(socialMediaId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		})
		return
	}
	// var userDB *models.User
	_, err = s.socialMediaService.CheckUser(email)
	log.Default().Println("error controller", err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	response := s.socialMediaService.UpdateSocialMedia(uint(idUser), uint(id), sosmedUpdate)
	c.JSON(response.Status, response)
}

func (s *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	socialMediaId := c.Param("socialMediaId")
	idUserToken := c.MustGet("id")
	idUser := int(idUserToken.(float64))
	emailUserToken := c.MustGet("email")
	email := fmt.Sprintf("%v", emailUserToken)
	idUint, err1 := strconv.Atoi(socialMediaId)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err1.Error(),
		})
		return
	}
	// var userDB *models.User
	_, err := s.socialMediaService.CheckUser(email)
	log.Default().Println("error controller", err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	userPhotoId, err := s.socialMediaService.GetSosmedUserId(uint(idUint))
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
	response := s.socialMediaService.DeleteSocialMedia(uint(idUint))
	c.JSON(response.Status, response)
}
