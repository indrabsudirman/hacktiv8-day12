package repositories

import (
	"errors"
	"hacktiv8-day12/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository interface {
	PostComment(*models.Comment) error
	GetAllComments() (*[]models.Comment, error)
	UpdateComment(uint, *models.Comment) (*models.Comment, error)
	CheckUser(string) (*models.User, error)
	DeleteComment(id uint) error
	GetCommentUserId(id uint) (uint, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

func (c *commentRepository) PostComment(comment *models.Comment) error {
	err := c.db.Create(comment).Error
	return err
}

func (c *commentRepository) GetAllComments() (*[]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Preload(clause.Associations).Find(&comments).Error //c.db.Find(&comments).Error
	return &comments, err
}

func (c *commentRepository) UpdateComment(id uint, comment *models.Comment) (*models.Comment, error) {
	var commentUpdate models.Comment
	err := c.db.First(&commentUpdate, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	err = c.db.Model(&commentUpdate).Where("id=?", id).Updates(
		models.Comment{
			UserID:  comment.UserID,
			PhotoID: comment.PhotoID,
			Message: comment.Message,
		},
	).Error
	if err != nil {
		return nil, err
	}
	return &commentUpdate, nil
}

func (c *commentRepository) CheckUser(email string) (*models.User, error) {
	var userDB models.User
	if err := c.db.Where("email=?", email).Take(&userDB).Error; err != nil {
		return nil, errors.New("invalid id/email")
	}
	return &userDB, nil
}

//Delete comment
func (c *commentRepository) DeleteComment(id uint) error {
	var commentDelete models.Comment
	err := c.db.First(&commentDelete, "id=?", id).Error
	if err != nil {
		return err
	}
	err = c.db.Delete(&commentDelete, "id=?", id).Error
	return err
}

func (c *commentRepository) GetCommentUserId(id uint) (uint, error) {
	var commentUpdate models.Comment
	err := c.db.First(&commentUpdate, "id=?", id).Error
	if err != nil {
		return 0, err
	}
	return commentUpdate.UserID, nil
}
