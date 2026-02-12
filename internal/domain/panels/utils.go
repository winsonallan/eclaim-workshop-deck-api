package panels

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"
)

func validateFixedPricing(req PricingRequest) error {
	serviceType := req.GetServiceType()
	laborFee := req.GetLaborFee()
	sparePart := req.GetSparePartCost()

	switch serviceType {
	case "repair":
		if laborFee == 0 {
			return errors.New("Since it's a repair, labor_fee must be filled")
		}
	case "replacement":
		if laborFee == 0 && sparePart == 0 {
			return errors.New("Since it's a replacement, labor_fee and spare_part_cost must be filled")
		}
	}
	return nil
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
