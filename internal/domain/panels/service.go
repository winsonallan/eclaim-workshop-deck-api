package panels

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

func (s *Service) GetAllPanels() ([]models.Panel, error) {
	return s.repo.GetAllPanels()
}
