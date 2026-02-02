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

func (s *Service) CreatePanelPricing(req CreatePanelPricingRequest) (*models.PanelPricing, error) {
	if req.WorkshopNo == 0 {
		return nil, errors.New("Workshop No is required")
	}
	if req.WorkshopPanelNo == 0 {
		return nil, errors.New("WorkshopPanel No is required")
	}
	if req.ServiceType == "" {
		return nil, errors.New("Service Type is required")
	}

	panelPricing := &models.PanelPricing{
		WorkshopNo:      req.WorkshopNo,
		WorkshopPanelNo: req.WorkshopPanelNo,
		ServiceType:     req.ServiceType,
	}

	if req.InsurerNo != 0 {
		panelPricing.InsurerNo = &req.InsurerNo
	}

	if req.MouNo != 0 {
		panelPricing.MouNo = &req.MouNo
	}

	if req.IsFixedPrice == true {
		panelPricing.IsFixedPrice = true
	} else {
		panelPricing.IsFixedPrice = false
	}

	if req.VehicleRangeLow != 0 {
		panelPricing.VehicleRangeLow = req.VehicleRangeLow
	} else {
		panelPricing.VehicleRangeLow = 0
	}

	if req.VehicleRangeHigh != 0 {
		panelPricing.VehicleRangeHigh = req.VehicleRangeHigh
	} else {
		panelPricing.VehicleRangeHigh = 999999999999
	}

	if req.SparePartCost > 0 {
		panelPricing.SparePartCost = req.SparePartCost
	}

	if req.LaborFee > 0 {
		panelPricing.LaborFee = req.LaborFee
	}

	if req.CreatedBy != 0 {
		panelPricing.CreatedBy = &req.CreatedBy
	}

	if err := s.repo.CreatePanelPricing(panelPricing); err != nil {
		return nil, err
	}

	return s.repo.FindPanelPricingById(panelPricing.PanelPricingNo)
}
