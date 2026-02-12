package orders

import (
	"eclaim-workshop-deck-api/internal/models"
)

func (s *Service) prepareOrderPanels(req OrderPanelRequest, createdBy, workOrderNo uint) (*models.OrderPanel, error) {

	// 1. Mapping
	orderPanel := &models.OrderPanel{
		WorkOrderNo:          workOrderNo,
		CreatedBy:            &createdBy,
		WorkOrderGroupNumber: 0,
	}

	if req.WOGroupNumber != 0 {
		orderPanel.WorkOrderGroupNumber = req.WOGroupNumber
	}

	if req.InsurancePanelPricingNo != 0 {
		orderPanel.InsurancePanelPricingNo = &req.InsurancePanelPricingNo
	}

	if req.InsurancePanelName != "" {
		orderPanel.InsurancePanelName = req.InsurancePanelName
	}

	if req.InsurerPrice != 0 {
		orderPanel.InsurerPrice = req.InsurerPrice
	}

	if req.InsurerMeasurementNo != 0 {
		orderPanel.InsurerMeasurementNo = &req.InsurerMeasurementNo
	}

	if req.InsurerServiceType != "" {
		orderPanel.InsurerServiceType = req.InsurerServiceType
	}

	if req.InsurerQty != 0 {
		orderPanel.InsurerQty = req.InsurerQty
	}

	if req.WorkshopPanelPricingNo != 0 {
		orderPanel.WorkshopPanelPricingNo = &req.WorkshopPanelPricingNo
	}

	if req.WorkshopPanelName != "" {
		orderPanel.WorkshopPanelName = req.WorkshopPanelName
	}

	if req.WorkshopPrice != 0 {
		orderPanel.WorkshopPrice = req.WorkshopPrice
	}

	if req.WorkshopMeasurementNo != 0 {
		orderPanel.WorkshopMeasurementNo = &req.WorkshopMeasurementNo
	}

	if req.WorkshopServiceType != "" {
		orderPanel.WorkshopServiceType = req.WorkshopServiceType
	}

	if req.WorkshopQty != 0 {
		orderPanel.WorkshopQty = req.WorkshopQty
	}

	if req.IsIncluded == true {
		orderPanel.IsIncluded = true
	} else {
		orderPanel.IsIncluded = false
	}

	if req.IsSpecialRepair == true {
		orderPanel.IsSpecialRepair = true
	} else {
		orderPanel.IsSpecialRepair = false
	}

	return orderPanel, nil
}
