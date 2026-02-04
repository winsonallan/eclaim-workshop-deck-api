package panels

import "errors"

func validateFixedPricing(req PricingRequest) error {
	serviceType := req.GetServiceType()
	laborFee := req.GetLaborFee()
	sparePart := req.GetSparePartCost()

	switch serviceType {
	case "Repair":
		if laborFee == 0 {
			return errors.New("Since it's a repair, labor_fee must be filled")
		}
	case "Replacement":
		if laborFee == 0 && sparePart == 0 {
			return errors.New("Since it's a replacement, labor_fee and spare_part_cost must be filled")
		}
	}
	return nil
}

func (r CreatePanelPricingRequest) GetServiceType() string { return r.ServiceType }
func (r CreatePanelPricingRequest) GetLaborFee() uint      { return r.LaborFee }
func (r CreatePanelPricingRequest) GetSparePartCost() uint { return r.SparePartCost }
func (r UpdatePanelPricingRequest) GetServiceType() string { return r.ServiceType }
func (r UpdatePanelPricingRequest) GetLaborFee() uint      { return r.LaborFee }
func (r UpdatePanelPricingRequest) GetSparePartCost() uint { return r.SparePartCost }
