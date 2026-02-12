package suppliers

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"
)

func (s *Service) prepareSupplier(req BaseSupplierRequest) (*models.Supplier, error) {
	if req.WorkshopNo == 0 {
		return nil, errors.New("Workshop No is required")
	}
	if req.SupplierName == "" {
		return nil, errors.New("Supplier name is required")
	}
	if req.SupplierAddress == "" {
		return nil, errors.New("Supplier address is required")
	}
	if req.SupplierCityNo == 0 {
		return nil, errors.New("Supplier city no is required")
	}
	if req.SupplierCityType == "" {
		return nil, errors.New("Supplier city type is required")
	}
	if req.SupplierCityName == "" {
		return nil, errors.New("Supplier city name is required")
	}
	if req.SupplierProvinceNo == 0 {
		return nil, errors.New("Supplier province no is required")
	}
	if req.SupplierProvinceName == "" {
		return nil, errors.New("Supplier province name is required")
	}
	if req.SupplierPhone == "" {
		return nil, errors.New("Supplier phone is required")
	}
	if req.SupplierEmail == "" {
		return nil, errors.New("Supplier email is required")
	}

	supplier := &models.Supplier{
		WorkshopNo:      req.WorkshopNo,
		SupplierName:    req.SupplierName,
		SupplierPhone:   req.SupplierPhone,
		SupplierEmail:   req.SupplierEmail,
		SupplierAddress: req.SupplierAddress,
		CityNo:          req.SupplierCityNo,
		CityType:        req.SupplierCityType,
		CityName:        req.SupplierCityName,
		ProvinceNo:      req.SupplierProvinceNo,
		ProvinceName:    req.SupplierProvinceName,
	}

	return supplier, nil
}
