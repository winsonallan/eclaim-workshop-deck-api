package admin

import (
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
		UserProfileCityName: req.CityName,
		UserProfileAddress:  req.Address,
		UserProfileEmail:    req.Email,
		UserProfilePhone:    req.Phone,
	}

	if err := s.repo.CreateUserProfile(userProfile); err != nil {
		return nil, err
	}

	return s.repo.FindUserProfileByID(userProfile.UserProfileNo)
}
