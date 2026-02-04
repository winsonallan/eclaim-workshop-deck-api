package panels

import "errors"

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
