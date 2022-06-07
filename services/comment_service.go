package services

import (
	"hacktiv8-day12/models"
	"hacktiv8-day12/params"
	"hacktiv8-day12/repositories"
	"log"
	"net/http"
	"time"
)

type CommentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(repo repositories.CommentRepository) *CommentService {
	return &CommentService{repo}
}

func (c *CommentService) PostComment(idUser uint, reqComment *params.Comment) *params.Response {
	var comment models.Comment
	comment.UserID = idUser
	comment.PhotoID = reqComment.PhotoId
	comment.Message = reqComment.Message
	comment.CreatedAt = time.Now()

	err := c.commentRepo.PostComment(&comment)
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
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"created_at": comment.CreatedAt,
		},
	}
}

func (c *CommentService) GetAllComments() (*[]params.GetAllComments, *params.Response) {
	comments, err := c.commentRepo.GetAllComments()
	if err != nil {
		return nil, &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "internal server error",
			AdditionalInfo: err.Error(),
		}
	}

	var commentlists []params.GetAllComments
	for _, comment := range *comments {
		commentlists = append(commentlists, params.GetAllComments{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			CreatedAt: &comment.CreatedAt,
			UpdatedAt: &comment.UpdatedAt,
			User: &params.User{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: &params.GetAllPhotos{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserID:   comment.Photo.UserID,
			},
		})

	}
	return &commentlists, nil
}

func (c *CommentService) UpdateComment(idUser, id uint, request params.UpdateComment) *params.Response {
	var comment models.Comment
	comment.Message = request.Message
	comment.UpdatedAt = time.Now()
	log.Default().Println("message", comment.Message)

	commentDB, err := c.commentRepo.UpdateComment(id, &comment)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "photo not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "comment success updated",
		Payload: map[string]interface{}{
			"id":         id,
			"message":    commentDB.Message,
			"photo_id":   commentDB.PhotoID,
			"user_id":    idUser,
			"updated_at": commentDB.UpdatedAt,
		},
	}
}

func (c *CommentService) CheckUser(email string) (*models.User, error) {
	var userDB *models.User
	userDB, err := c.commentRepo.CheckUser(email)
	if err != nil {
		return nil, err
	}
	return userDB, nil
}

func (c *CommentService) DeleteComment(id uint) *params.Response {
	err := c.commentRepo.DeleteComment(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "comment not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "your comment has been successfully deleted",
	}
}

func (c *CommentService) GetCommentUserId(id uint) (uint, error) {
	userCommentId, err := c.commentRepo.GetCommentUserId(id)
	if err != nil {
		return 0, err
	}
	return userCommentId, nil
}
