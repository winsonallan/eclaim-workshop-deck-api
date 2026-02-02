package panels

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"
	"time"
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

func (s *Service) GetAllWorkshopPanels(workshopId uint) ([]models.WorkshopPanels, error) {
	return s.repo.GetAllWorkshopPanels(workshopId)
}

func (s *Service) GetMOUs(insurerId uint, workshopId uint, mouId uint, activeOnly bool) ([]models.MOU, error) {
	return s.repo.GetMOUs(insurerId, workshopId, mouId, activeOnly)
}

func (s *Service) GetPanelPricings(insurerId, workshopId, mouId uint) ([]models.PanelPricing, error) {
	return s.repo.GetPanelPricings(insurerId, workshopId, mouId)
}

func (s *Service) CreateMOU(req CreateMOURequest) (*models.MOU, error) {
	if req.MouDocumentNumber == "" {
		return nil, errors.New("MOU Document Number is required")
	}
	if req.InsurerNo == 0 {
		return nil, errors.New("Insurer No is required")
	}
	if req.WorkshopNo == 0 {
		return nil, errors.New("Workshop No is required")
	}

	mou := &models.MOU{
		MouDocumentNumber: req.MouDocumentNumber,
		InsurerNo:         req.InsurerNo,
		WorkshopNo:        req.WorkshopNo,
	}

	if req.CreatedBy != 0 {
		mou.CreatedBy = &req.CreatedBy
	}

	if req.MouExpiryDate != "" {
		// Layout format: YYYY-MM-DD. Change this if your input format is different!
		layout := "2006-01-02"

		parsedDate, err := time.Parse(layout, req.MouExpiryDate)
		if err != nil {
			// Handle the error (e.g., return an "invalid date format" error)
			return nil, err
		}

		mou.MouExpiryDate = parsedDate
	}

	if err := s.repo.CreateMOU(mou); err != nil {
		return nil, err
	}

	return s.repo.FindMOUByID(mou.MouNo)
}
