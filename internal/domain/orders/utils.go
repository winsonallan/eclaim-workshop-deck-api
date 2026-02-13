package orders

import (
	"eclaim-workshop-deck-api/internal/domain/panels"
	"eclaim-workshop-deck-api/internal/models"
	"errors"
)

func (s *Service) prepareOrderPanels(req OrderPanelRequest, createdBy, workOrderNo uint) (*models.OrderPanel, error) {
	panelService := panels.NewRepository(s.repo.db)

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

		insurancePanel, err := panelService.FindPanelPricingById(req.InsurancePanelPricingNo)

		if err != nil {
			return nil, errors.New("insurance panel pricing no is invalid")
		}

		orderPanel.InsurancePanelName = insurancePanel.WorkshopPanels.PanelName
		orderPanel.InsurerServiceType = insurancePanel.ServiceType

		switch insurancePanel.ServiceType {
		case "repair":
			if req.InsurerMeasurementNo != 0 {
				orderPanel.InsurerMeasurementNo = &req.InsurerMeasurementNo

				var chosenMeasurement models.Measurement
				for _, m := range insurancePanel.Measurements {
					if m.MeasurementNo == *orderPanel.InsurerMeasurementNo {
						chosenMeasurement = m
					}
				}

				orderPanel.InsurerPrice = chosenMeasurement.LaborFee
			} else {
				orderPanel.InsurerPrice = insurancePanel.LaborFee
			}
		case "replacement":
			orderPanel.InsurerPrice = insurancePanel.LaborFee + insurancePanel.SparePartCost
		default:
			return nil, errors.New("service type error while preparing order panels (insurance panel)")
		}
	}

	if req.InsurerQty != 0 {
		orderPanel.InsurerQty = req.InsurerQty
	}

	if req.WorkshopPanelPricingNo != 0 {
		orderPanel.WorkshopPanelPricingNo = &req.WorkshopPanelPricingNo

		workshopPanel, err := panelService.FindPanelPricingById(req.WorkshopPanelPricingNo)

		if err != nil {
			return nil, errors.New("insurance panel pricing no is invalid")
		}

		orderPanel.WorkshopPanelName = workshopPanel.WorkshopPanels.PanelName
		orderPanel.WorkshopServiceType = workshopPanel.ServiceType

		switch workshopPanel.ServiceType {
		case "repair":
			if req.WorkshopMeasurementNo != 0 {
				orderPanel.WorkshopMeasurementNo = &req.WorkshopMeasurementNo

				var chosenMeasurement models.Measurement
				for _, m := range workshopPanel.Measurements {
					if m.MeasurementNo == *orderPanel.WorkshopMeasurementNo {
						chosenMeasurement = m
					}
				}

				orderPanel.WorkshopPrice = chosenMeasurement.LaborFee
			} else {
				orderPanel.WorkshopPrice = workshopPanel.LaborFee
			}
		case "replacement":
			orderPanel.WorkshopPrice = workshopPanel.LaborFee + workshopPanel.SparePartCost
		default:
			return nil, errors.New("service type error while preparing order panels (insurance panel)")
		}
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
