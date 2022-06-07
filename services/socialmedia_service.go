package services

import (
	"hacktiv8-day12/models"
	"hacktiv8-day12/params"
	"hacktiv8-day12/repositories"
	"net/http"
	"time"
)

type SocialMediaService struct {
	sosmedRepo repositories.SocialMediaRepository
}

func NewSocialMediaService(repo repositories.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{repo}
}

func (s *SocialMediaService) PostSocialMedia(userId uint, sosmed *params.SocialMedia) *params.Response {
	var sosmedModel models.SocialMedia
	sosmedModel.Name = sosmed.Name
	sosmedModel.SocialMediaUrl = sosmed.SocialMediaUrl
	sosmedModel.UserID = userId
	sosmedModel.CreatedAt = time.Now()

	err := s.sosmedRepo.PostSocialMedia(&sosmedModel)
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
			"id":               sosmedModel.ID,
			"name":             sosmedModel.Name,
			"social_media_url": sosmedModel.SocialMediaUrl,
			"user_id":          sosmedModel.UserID,
			"created_at":       sosmedModel.CreatedAt,
		},
	}
}

func (s *SocialMediaService) GetAllSocialMedias() (*[]params.GetAllSocialMedias, *params.Response) {
	sosmeds, err := s.sosmedRepo.GetAllSocialMedias()
	if err != nil {
		return nil, &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "bad request",
			AdditionalInfo: err.Error(),
		}
	}

	var socialMediasList []params.GetAllSocialMedias

	for _, sosmed := range *sosmeds {
		socialMediasList = append(socialMediasList, params.GetAllSocialMedias{
			ID:             sosmed.ID,
			Name:           sosmed.Name,
			SocialMediaUrl: sosmed.SocialMediaUrl,
			UserID:         sosmed.UserID,
			CreatedAt:      &sosmed.CreatedAt,
			UpdatedAt:      &sosmed.UpdatedAt,
			User: &params.User{
				ID:       sosmed.User.ID,
				Username: sosmed.User.Username,
			},
		})
	}
	return &socialMediasList, nil
}

func (s *SocialMediaService) UpdateSocialMedia(idUser, id uint, request params.SocialMedia) *params.Response {
	var sosmed models.SocialMedia
	sosmed.Name = request.Name
	sosmed.SocialMediaUrl = request.SocialMediaUrl
	sosmed.UpdatedAt = time.Now()

	sosmedDB, err := s.sosmedRepo.UpdateSocialMedia(id, &sosmed)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "photo not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "social media success updated",
		Payload: map[string]interface{}{
			"id":               sosmedDB.ID,
			"name":             sosmedDB.Name,
			"social_media_url": sosmedDB.SocialMediaUrl,
			"user_id":          sosmedDB.UserID,
			"updated_at":       sosmedDB.UpdatedAt,
		},
	}
}

func (s *SocialMediaService) CheckUser(email string) (*models.User, error) {
	var userDB *models.User
	userDB, err := s.sosmedRepo.CheckUser(email)
	if err != nil {
		return nil, err
	}
	return userDB, nil
}

func (s *SocialMediaService) DeleteSocialMedia(id uint) *params.Response {
	err := s.sosmedRepo.DeleteSocialMedia(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "social media not found",
			AdditionalInfo: err.Error(),
		}
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "your social media has been successfully deleted",
	}
}

func (s *SocialMediaService) GetSosmedUserId(id uint) (uint, error) {
	sosmedUserId, err := s.sosmedRepo.GetSocialMediaUserId(id)
	if err != nil {
		return 0, err
	}
	return sosmedUserId, nil
}
