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

func (s *Service) preparePanelPricing(req BasePricingRequest) (*models.PanelPricing, error) {
	// 1. Core Validation
	if req.WorkshopNo == 0 {
		return nil, errors.New("Workshop No is required")
	}
	if req.ServiceType == "" {
		return nil, errors.New("Service Type is required")
	}

	// 2. Pricing Logic Validation
	if *req.IsFixedPrice {
		// Now BasePricingRequest implements PricingRequest interface
		if err := validateFixedPricing(req); err != nil {
			return nil, err
		}
	} else if len(req.Measurements) == 0 {
		return nil, errors.New("measurements must exist since it's conditional pricing")
	}

	// 3. Mapping
	panelPricing := &models.PanelPricing{
		WorkshopNo:       req.WorkshopNo,
		ServiceType:      req.ServiceType,
		IsFixedPrice:     *req.IsFixedPrice,
		SparePartCost:    req.SparePartCost,
		LaborFee:         req.LaborFee,
		VehicleRangeLow:  0,
		VehicleRangeHigh: 999999999999,
	}

	if req.VehicleRangeLow != 0 {
		panelPricing.VehicleRangeLow = req.VehicleRangeLow
	}
	if req.VehicleRangeHigh != 0 && req.VehicleRangeHigh != 999999999999 {
		panelPricing.VehicleRangeHigh = req.VehicleRangeHigh
	}

	if req.InsurerNo != 0 {
		panelPricing.InsurerNo = &req.InsurerNo
	}
	if req.MouNo != 0 {
		panelPricing.MouNo = &req.MouNo
	}
	if req.AdditionalNote != "" {
		panelPricing.AdditionalNotes = req.AdditionalNote
	}

	return panelPricing, nil
}

func (s *Service) CreatePanelPricing(req CreatePanelPricingRequest) (*models.PanelPricing, error) {
	// Pass the embedded BasePricingRequest field
	panelPricing, err := s.preparePanelPricing(req.BasePricingRequest)
	if err != nil {
		return nil, err
	}

	// Handle Custom Panel creation
	if *req.IsCustom {
		if req.CustomPanelName == "" {
			return nil, errors.New("custom panel name is required")
		}
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
		if req.WorkshopPanelNo == 0 {
			return nil, errors.New("workshop_panel_no is required")
		}
		panelPricing.WorkshopPanelNo = req.WorkshopPanelNo
	}

	panelPricing.CreatedBy = &req.CreatedBy

	// Save Main Record
	if err := s.repo.CreatePanelPricing(panelPricing); err != nil {
		return nil, err
	}

	// Handle Measurements
	if !*req.IsFixedPrice {
		for _, m := range req.Measurements {
			measurement := &models.Measurement{
				PanelPricingNo: panelPricing.PanelPricingNo,
				ConditionText:  m.ConditionText,
				Notes:          m.Note,
				LaborFee:       m.LaborFee,
				CreatedBy:      &req.CreatedBy,
			}
			if err := s.repo.CreateMeasurement(measurement); err != nil {
				return nil, err
			}
		}
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
	// 1. Check existence
	existing, err := s.repo.FindPanelPricingById(id)
	if err != nil {
		return nil, errors.New("panel pricing not found")
	}

	// 2. Map & Validate using the embedded Base struct
	updatedData, err := s.preparePanelPricing(req.BasePricingRequest)
	if err != nil {
		return nil, err
	}

	// 3. Handle Custom Panel Logic
	if *req.IsCustom {
		if req.CustomPanelName == "" {
			return nil, errors.New("custom panel name is required")
		}
		workshopPanel := &models.WorkshopPanels{
			WorkshopNo: req.WorkshopNo,
			PanelName:  req.CustomPanelName,
			CreatedBy:  &req.LastModifiedBy,
		}
		if err := s.repo.CreateWorkshopPanel(workshopPanel); err != nil {
			return nil, err
		}
		existing.WorkshopPanelNo = workshopPanel.WorkshopPanelNo
	} else {
		existing.WorkshopPanelNo = req.WorkshopPanelNo
	}

	// 4. Update core fields
	existing.WorkshopNo = updatedData.WorkshopNo
	existing.ServiceType = updatedData.ServiceType
	existing.IsFixedPrice = updatedData.IsFixedPrice
	existing.SparePartCost = updatedData.SparePartCost
	existing.LaborFee = updatedData.LaborFee
	existing.VehicleRangeLow = updatedData.VehicleRangeLow
	existing.VehicleRangeHigh = updatedData.VehicleRangeHigh
	existing.InsurerNo = updatedData.InsurerNo
	existing.MouNo = updatedData.MouNo
	existing.AdditionalNotes = updatedData.AdditionalNotes
	existing.LastModifiedBy = &req.LastModifiedBy

	// 5. Save Main Record
	if err := s.repo.UpdatePanelPricing(existing); err != nil {
		return nil, err
	}

	// 6. Sync Measurements (Soft Delete then Re-insert)
	if err := s.repo.SoftDeleteMeasurementsByPanelPricingNo(id); err != nil {
		return nil, err
	}

	if !*req.IsFixedPrice {
		for _, m := range req.Measurements {
			newM := &models.Measurement{
				PanelPricingNo: id,
				ConditionText:  m.ConditionText,
				Notes:          m.Note,
				LaborFee:       m.LaborFee,
				CreatedBy:      &req.LastModifiedBy,
				IsLocked:       false, // Changed from 0 to false for bool type
			}
			if err := s.repo.CreateMeasurement(newM); err != nil {
				return nil, err
			}
		}
	}

	return s.repo.FindPanelPricingById(id)
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
