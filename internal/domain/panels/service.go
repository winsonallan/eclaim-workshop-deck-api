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

func (s *Service) GetWorkshopPanelPricings(workshopId uint) ([]models.PanelPricing, error) {
	return s.repo.GetWorkshopPanelPricings(workshopId)
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
	// 1. Core Validation
	if req.WorkshopNo == 0 {
		return nil, errors.New("Workshop No is required")
	}
	if req.ServiceType == "" {
		return nil, errors.New("Service Type is required")
	}

	// 2. Custom vs Standard Panel Logic
	if *req.IsCustom {
		if req.CustomPanelName == "" {
			return nil, errors.New("Since it's a custom panel, custom panel name is required")
		}
	} else if req.WorkshopPanelNo == 0 {
		return nil, errors.New("Since it's not a custom panel, workshop_panel_no is required")
	}

	// 3. Pricing Logic (Fixed vs Conditional)
	if *req.IsFixedPrice {
		if err := validateFixedPricing(req); err != nil {
			return nil, err
		}
	} else {
		// Validation: Ensure the slice has at least one entry
		if len(req.Measurements) == 0 {
			return nil, errors.New("measurements must exist since it's conditional pricing")
		}
	}

	// 4. Data Mapping
	panelPricing := &models.PanelPricing{
		WorkshopNo:       req.WorkshopNo,
		ServiceType:      req.ServiceType,
		IsFixedPrice:     *req.IsFixedPrice,
		VehicleRangeLow:  0,
		VehicleRangeHigh: 999999999999,
		SparePartCost:    req.SparePartCost,
		LaborFee:         req.LaborFee,
	}

	if req.VehicleRangeLow != 0 {
		panelPricing.VehicleRangeLow = req.VehicleRangeLow
	}

	if req.VehicleRangeHigh != 999999999999 && req.VehicleRangeHigh != 0 {
		panelPricing.VehicleRangeHigh = req.VehicleRangeHigh
	}

	// Handle Optional Pointers
	if req.InsurerNo != 0 {
		panelPricing.InsurerNo = &req.InsurerNo
	}
	if req.MouNo != 0 {
		panelPricing.MouNo = &req.MouNo
	}
	if req.AdditionalNote != "" {
		panelPricing.AdditionalNotes = req.AdditionalNote
	}
	if req.CreatedBy != 0 {
		panelPricing.CreatedBy = &req.CreatedBy
	}

	if *req.IsCustom {
		workshopPanel := &models.WorkshopPanels{
			WorkshopNo: req.WorkshopNo,
			PanelName:  req.CustomPanelName,
			CreatedBy:  &req.CreatedBy,
		}

		if err := s.repo.CreateWorkshopPanel(workshopPanel); err != nil {
			return nil, err
		}

		panelPricing.WorkshopPanelNo = workshopPanel.WorkshopPanelNo
	} else {
		panelPricing.WorkshopPanelNo = req.WorkshopPanelNo
	}

	// 5. Database Operations
	if err := s.repo.CreatePanelPricing(panelPricing); err != nil {
		return nil, err
	}

	return s.repo.FindPanelPricingById(panelPricing.PanelPricingNo)
}

func (s *Service) CreateWorkshopPanel(req CreateWorkshopPanelRequest) (*models.WorkshopPanels, error) {
	if req.WorkshopNo == 0 {
		return nil, errors.New("Workshop_no is required")
	}

	if req.CreatedBy == 0 {
		return nil, errors.New("Created_by is required")
	}

	if req.PanelName == "" {
		return nil, errors.New("Panel Name is required!")
	}

	workshopPanel := &models.WorkshopPanels{
		WorkshopNo: req.WorkshopNo,
		PanelNo:    req.PanelNo,
		PanelName:  req.PanelName,
		CreatedBy:  &req.CreatedBy,
	}

	if err := s.repo.CreateWorkshopPanel(workshopPanel); err != nil {
		return nil, err
	}

	return s.repo.FindWorkshopPanelById(workshopPanel.WorkshopPanelNo)
}

// Update
func (s *Service) UpdatePanelPricing(id uint, req UpdatePanelPricingRequest) (*models.PanelPricing, error) {
	panelPricing, err := s.repo.FindPanelPricingById(id)
	if err != nil {
		return nil, errors.New("panel pricing not found")
	}
	// 1. Core Validation
	if req.WorkshopNo == 0 {
		return nil, errors.New("Workshop No is required")
	}
	if req.ServiceType == "" {
		return nil, errors.New("Service Type is required")
	}

	// 2. Custom vs Standard Panel Logic
	if *req.IsCustom {
		if req.CustomPanelName == "" {
			return nil, errors.New("Since it's a custom panel, custom panel name is required")
		}
	} else if req.WorkshopPanelNo == 0 {
		return nil, errors.New("Since it's not a custom panel, workshop_panel_no is required")
	}

	// 3. Pricing Logic (Fixed vs Conditional)
	if *req.IsFixedPrice {
		if err := validateFixedPricing(req); err != nil {
			return nil, err
		}
	} else {
		// Validation: Ensure the slice has at least one entry
		if len(req.Measurements) == 0 {
			return nil, errors.New("measurements must exist since it's conditional pricing")
		}
	}

	panelPricing.WorkshopNo = req.WorkshopNo
	panelPricing.WorkshopPanelNo = req.WorkshopPanelNo
	panelPricing.ServiceType = req.ServiceType
	panelPricing.IsFixedPrice = *req.IsFixedPrice
	panelPricing.SparePartCost = req.SparePartCost
	panelPricing.LaborFee = req.LaborFee

	if req.VehicleRangeLow != 0 {
		panelPricing.VehicleRangeLow = req.VehicleRangeLow
	}

	if req.VehicleRangeHigh != 999999999999 && req.VehicleRangeHigh != 0 {
		panelPricing.VehicleRangeHigh = req.VehicleRangeHigh
	}

	// Handle Optional Pointers
	if req.InsurerNo != 0 {
		panelPricing.InsurerNo = &req.InsurerNo
	}

	if req.MouNo != 0 {
		panelPricing.MouNo = &req.MouNo
	}

	if req.AdditionalNote != "" {
		panelPricing.AdditionalNotes = req.AdditionalNote
	}

	if req.LastModifiedBy != 0 {
		panelPricing.LastModifiedBy = &req.LastModifiedBy
	}

	if *req.IsCustom {
		workshopPanel := &models.WorkshopPanels{
			WorkshopNo: req.WorkshopNo,
			PanelName:  req.CustomPanelName,
			CreatedBy:  &req.LastModifiedBy,
		}

		if err := s.repo.CreateWorkshopPanel(workshopPanel); err != nil {
			return nil, err
		}

		panelPricing.WorkshopPanelNo = workshopPanel.WorkshopPanelNo
	} else {
		panelPricing.WorkshopPanelNo = req.WorkshopPanelNo
	}

	// 5. Database Operations
	if err := s.repo.UpdatePanelPricing(panelPricing); err != nil {
		return nil, err
	}

	return s.repo.FindPanelPricingById(panelPricing.PanelPricingNo)
}

// Delete
func (s *Service) DeletePanelPricing(id uint, req DeletePanelPricingRequest) (*models.PanelPricing, error) {
	panelPricing, err := s.repo.FindPanelPricingById(id)

	if err != nil {
		return nil, errors.New("panel pricing not found")
	}

	panelPricing.IsLocked = true
	panelPricing.LastModifiedBy = &req.LastModifiedBy

	if err := s.repo.UpdatePanelPricing(panelPricing); err != nil {
		return nil, err
	}

	return panelPricing, nil
}
