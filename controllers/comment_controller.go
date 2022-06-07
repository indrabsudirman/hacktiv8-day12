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

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{*service}
}

func (cmt *CommentController) PostComment(c *gin.Context) {
	var comment params.Comment
	idUserToken := c.MustGet("id")
	idUser := int(idUserToken.(float64))

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}

	response := cmt.commentService.PostComment(uint(idUser), &comment)
	c.JSON(response.Status, response)
}

//Get all comments
func (cmt *CommentController) GetAllComments(c *gin.Context) {
	response, _ := cmt.commentService.GetAllComments()
	log.Default().Println("response:", response)
	c.JSON(http.StatusOK, response)
}

//Update comment
func (cmt *CommentController) UpdateComment(c *gin.Context) {
	var commentUpdate params.UpdateComment
	commentId := c.Param("commentId")
	idUserToken := c.MustGet("id")
	emailUserToken := c.MustGet("email")
	idUser := int(idUserToken.(float64))
	email := fmt.Sprintf("%v", emailUserToken)

	err := c.ShouldBindJSON(&commentUpdate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			params.Response{
				Status:         http.StatusBadRequest,
				Error:          "bad request",
				AdditionalInfo: err.Error(),
			})
		return
	}
	id, err := strconv.Atoi(commentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		})
		return
	}
	// var userDB *models.User
	_, err = cmt.commentService.CheckUser(email)
	log.Default().Println("error controller", err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	response := cmt.commentService.UpdateComment(uint(idUser), uint(id), commentUpdate)
	c.JSON(response.Status, response)
}

func (cmt *CommentController) DeleteComment(c *gin.Context) {
	commentId := c.Param("commentId")
	idUserToken := c.MustGet("id")
	idUser := int(idUserToken.(float64))
	emailUserToken := c.MustGet("email")
	email := fmt.Sprintf("%v", emailUserToken)
	idUint, err1 := strconv.Atoi(commentId)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err1.Error(),
		})
		return
	}
	// var userDB *models.User
	_, err := cmt.commentService.CheckUser(email)
	log.Default().Println("error controller", err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
			Status:  http.StatusUnauthorized,
			Error:   "unauthorized",
			Message: "your are not allowed to access this data",
		})
		return
	}
	userPhotoId, err := cmt.commentService.GetCommentUserId(uint(idUint))
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
	response := cmt.commentService.DeleteComment(uint(idUint))
	c.JSON(response.Status, response)
}
