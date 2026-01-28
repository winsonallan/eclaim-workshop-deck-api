package location

import (
	"eclaim-workshop-deck-api/internal/models"
)

type Service struct {
	repo      *Repository
	jwtSecret string
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetCities() ([]models.City, error) {
	return s.repo.GetCities()
}
