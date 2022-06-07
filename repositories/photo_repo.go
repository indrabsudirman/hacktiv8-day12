package repositories

import (
	"errors"
	"hacktiv8-day12/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository interface {
	PostPhoto(uint, *models.Photo) error
	GetAllPhotos() (*[]models.Photo, error)
	UpdatePhoto(uint, *models.Photo) error
	CheckUser(string) (*models.User, error)
	DeletePhoto(uint) error
	GetPhotoUserId(uint) (uint, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{
		db: db,
	}
}

//Post Photo
func (p *photoRepository) PostPhoto(userId uint, photo *models.Photo) error {
	err := p.db.Create(photo).Error
	return err
}

//Get all photos
func (p *photoRepository) GetAllPhotos() (*[]models.Photo, error) {
	var photos []models.Photo
	err := p.db.Preload(clause.Associations).Find(&photos).Error //Exec("select photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id, photos.created_at, photos.updated_at,users.username, users.email from photos inner join users on photos.user_id = users.id ").Scan(&photos).Error//Error //Model(models.Photo{}).Joins("JOIN users ON users.id = photos.user_id").Scan(&photos).Error
	return &photos, err
}

//Update photo
func (p *photoRepository) UpdatePhoto(id uint, request *models.Photo) error {
	var photo models.Photo
	err := p.db.First(&photo, "id=?", id).Error
	if err != nil {
		return err
	}
	err = p.db.Model(&photo).Where("id=?", id).Updates(models.Photo{Title: request.Title, Caption: request.Caption, PhotoUrl: request.PhotoUrl}).Error
	return err
}

func (p *photoRepository) CheckUser(email string) (*models.User, error) {
	var userDB models.User
	if err := p.db.Where("email=?", email).Take(&userDB).Error; err != nil {
		return nil, errors.New("invalid id/email")
	}
	return &userDB, nil
}

//Delete photo
func (p *photoRepository) DeletePhoto(id uint) error {
	var photo models.Photo
	err := p.db.First(&photo, "id=?", id).Error
	if err != nil {
		return err
	}
	err = p.db.Delete(&photo, "id=?", id).Error
	return err
}

func (p *photoRepository) GetPhotoUserId(id uint) (uint, error) {
	var photo models.Photo
	err := p.db.First(&photo, "id=?", id).Error
	if err != nil {
		return 0, err
	}
	return photo.UserID, nil
}
