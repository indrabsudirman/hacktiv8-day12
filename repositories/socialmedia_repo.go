package repositories

import (
	"errors"
	"hacktiv8-day12/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialMediaRepository interface {
	PostSocialMedia(*models.SocialMedia) error
	GetAllSocialMedias() (*[]models.SocialMedia, error)
	UpdateSocialMedia(uint, *models.SocialMedia) (*models.SocialMedia, error)
	CheckUser(string) (*models.User, error)
	DeleteSocialMedia(id uint) error
	GetSocialMediaUserId(id uint) (uint, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db}
}

func (s *socialMediaRepository) PostSocialMedia(sosmed *models.SocialMedia) error {
	err := s.db.Create(sosmed).Error
	return err
}

func (s *socialMediaRepository) GetAllSocialMedias() (*[]models.SocialMedia, error) {
	var sosmeds []models.SocialMedia
	err := s.db.Preload(clause.Associations).Find(&sosmeds).Error //s.db.Find(&sosmeds).Error
	return &sosmeds, err
}

func (s *socialMediaRepository) UpdateSocialMedia(id uint, sosmedReq *models.SocialMedia) (*models.SocialMedia, error) {
	var sosmedUpdate models.SocialMedia
	err := s.db.First(&sosmedUpdate, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	err = s.db.Model(&sosmedUpdate).Where("id=?", id).Updates(
		models.SocialMedia{
			Name:           sosmedReq.Name,
			SocialMediaUrl: sosmedReq.SocialMediaUrl,
		},
	).Error
	if err != nil {
		return nil, err
	}
	return &sosmedUpdate, nil
}

func (s *socialMediaRepository) CheckUser(email string) (*models.User, error) {
	var userDB models.User
	if err := s.db.Where("email=?", email).Take(&userDB).Error; err != nil {
		return nil, errors.New("invalid id/email")
	}
	return &userDB, nil
}

//Delete social media
func (s *socialMediaRepository) DeleteSocialMedia(id uint) error {
	var socialMediaDelete models.SocialMedia
	err := s.db.First(&socialMediaDelete, "id=?", id).Error
	if err != nil {
		return err
	}
	err = s.db.Delete(&socialMediaDelete, "id=?", id).Error
	return err
}

func (s *socialMediaRepository) GetSocialMediaUserId(id uint) (uint, error) {
	var socialMedia models.SocialMedia
	err := s.db.First(&socialMedia, "id=?", id).Error
	if err != nil {
		return 0, err
	}
	return socialMedia.UserID, nil
}
