package services

import (
	"hacktiv8-day12/models"
	"hacktiv8-day12/params"
	"hacktiv8-day12/repositories"
	"net/http"
	"time"
)

type PhotoService struct {
	photoRepo repositories.PhotoRepository
}

func NewPhotoService(repo repositories.PhotoRepository) *PhotoService {
	return &PhotoService{repo}
}

func (p *PhotoService) PostPhoto(userId uint, photo *params.Photo) *params.Response {
	var photoModel models.Photo
	photoModel.Title = photo.Title
	photoModel.Caption = photo.Caption
	photoModel.PhotoUrl = photo.PhotoUrl
	photoModel.UserID = userId
	photoModel.CreatedAt = time.Now()

	err := p.photoRepo.PostPhoto(userId, &photoModel)
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
			"id":         photoModel.ID,
			"title":      photoModel.Title,
			"caption":    photoModel.Caption,
			"photo_url":  photoModel.PhotoUrl,
			"user_id":    photoModel.UserID,
			"created_at": photoModel.CreatedAt,
		},
	}
}

func (p *PhotoService) GetAllPhotos() (*[]params.GetAllPhotos, *params.Response) {
	photos, err := p.photoRepo.GetAllPhotos()
	if err != nil {
		return nil, &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "internal server error",
			AdditionalInfo: err.Error(),
		}
	}

	var photosList []params.GetAllPhotos
	for _, photo := range *photos {
		photosList = append(photosList, params.GetAllPhotos{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: &photo.CreatedAt,
			UpdatedAt: &photo.UpdatedAt,
			User: &params.User{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	return &photosList, nil
}

func (p *PhotoService) UpdatePhoto(idUser, id uint, request params.Photo) *params.Response {
	var photo models.Photo
	photo.Title = request.Title
	photo.Caption = request.Caption
	photo.PhotoUrl = request.PhotoUrl
	photo.UpdatedAt = time.Now()

	err := p.photoRepo.UpdatePhoto(id, &photo)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "photo not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "photo success updated",
		Payload: map[string]interface{}{
			"id":         id,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoUrl,
			"user_id":    idUser,
			"updated_at": photo.UpdatedAt,
		},
	}
}

func (p *PhotoService) CheckUser(email string) (*models.User, error) {
	var userDB *models.User
	userDB, err := p.photoRepo.CheckUser(email)
	if err != nil {
		return nil, err
	}
	return userDB, nil
}

func (p *PhotoService) DeletePhoto(id uint) *params.Response {
	err := p.photoRepo.DeletePhoto(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "photo not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "your photo has been successfully deleted",
	}
}

func (p *PhotoService) GetPhotoUserId(id uint) (uint, error) {
	userPhotoId, err := p.photoRepo.GetPhotoUserId(id)
	if err != nil {
		return 0, err
	}
	return userPhotoId, nil
}
