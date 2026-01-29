package admin

import (
	"eclaim-workshop-deck-api/internal/domain/settings"
	"eclaim-workshop-deck-api/internal/models"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUserProfile(req CreateUserProfileRequest) (*models.UserProfile, error) {
	if req.Type == "" {
		return nil, errors.New("type is required")
	}
	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	userProfile := &models.UserProfile{
		UserProfileType:     req.Type,
		UserProfileName:     req.Name,
		UserProfileCityNo:   req.CityNo,
		UserProfileCityType: req.CityType,
		UserProfileCityName: req.CityName,
		UserProfileAddress:  req.Address,
		UserProfileEmail:    req.Email,
		UserProfilePhone:    req.Phone,
		CreatedBy:           &req.CreatedBy,
	}

	if err := s.repo.CreateUserProfile(userProfile); err != nil {
		return nil, err
	}

	if req.Type == "workshop" {
		settingsRepo := settings.NewRepository(s.repo.db)

		workshopDetails := &models.WorkshopDetails{
			UserProfileNo: userProfile.UserProfileNo,
			IsAuthorized:  req.IsAuthorized,
			CreatedBy:     &req.CreatedBy,
		}

		if req.Capacity != 0 {
			workshopDetails.Capacity = req.Capacity
		}

		if req.Description != "" {
			workshopDetails.Description = req.Description
		}

		if req.Specialist != "" {
			workshopDetails.Specialist = req.Specialist
		}

		if err := settingsRepo.CreateWorkshopDetails(workshopDetails); err != nil {
			return nil, err
		}
	}

	return s.repo.FindUserProfileByID(userProfile.UserProfileNo)
}
