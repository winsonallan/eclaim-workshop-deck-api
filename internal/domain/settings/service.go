package settings

import (
	"eclaim-workshop-deck-api/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAccount(id uint) ([]models.User, error) {
	return s.repo.GetAccount(id)
}

func (s *Service) GetProfileDetails(id uint) ([]models.UserProfile, error) {
	return s.repo.GetProfileDetails(id)
}

func (s *Service) GetWorkshopDetails(id uint) ([]models.WorkshopDetails, error) {
	return s.repo.GetWorkshopDetails(id)
}
