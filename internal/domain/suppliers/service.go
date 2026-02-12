package suppliers

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetSuppliers() ([]models.Supplier, error) {
	return s.repo.GetSuppliers()
}

func (s *Service) GetWorkshopSuppliers(workshopId uint) ([]models.Supplier, error) {
	return s.repo.GetWorkshopSuppliers(workshopId)
}

func (s *Service) AddSupplier(id uint, req AddSupplierRequest) (*models.Supplier, error) {
	req.BaseSupplierRequest.WorkshopNo = id
	supplier, err := s.prepareSupplier(req.BaseSupplierRequest)
	if err != nil {
		return nil, err
	}

	supplier.CreatedBy = &req.CreatedBy

	if err := s.repo.AddSupplier(supplier); err != nil {
		return nil, err
	}

	return s.repo.FindSupplierByID(supplier.SupplierNo)
}

// Update
func (s *Service) UpdateSupplier(id uint, req UpdateSupplierRequest) (*models.Supplier, error) {
	req.BaseSupplierRequest.WorkshopNo = id

	// 1. Check existence
	existing, err := s.repo.FindSupplierByID(id)
	if err != nil {
		return nil, errors.New("supplier not found")
	}

	// 2. Map & Validate using the embedded Base struct
	updatedData, err := s.prepareSupplier(req.BaseSupplierRequest)
	if err != nil {
		return nil, err
	}

	// 4. Update core fields
	existing.WorkshopNo = updatedData.WorkshopNo
	existing.SupplierName = updatedData.SupplierName
	existing.SupplierAddress = updatedData.SupplierAddress
	existing.SupplierEmail = updatedData.SupplierEmail
	existing.SupplierPhone = updatedData.SupplierPhone
	existing.CityNo = updatedData.CityNo
	existing.CityType = updatedData.CityType
	existing.CityName = updatedData.CityName
	existing.ProvinceNo = updatedData.ProvinceNo
	existing.ProvinceName = updatedData.ProvinceName
	existing.LastModifiedBy = &req.LastModifiedBy

	// 5. Save Main Record
	if err := s.repo.UpdateSupplier(existing); err != nil {
		return nil, err
	}

	return s.repo.FindSupplierByID(id)
}

// Delete
func (s *Service) DeleteSupplier(id uint, req DeleteSupplierRequest) (*models.Supplier, error) {
	panelPricing, err := s.repo.FindSupplierByID(id)

	if err != nil {
		return nil, errors.New("supplier not found")
	}

	panelPricing.IsLocked = true
	panelPricing.LastModifiedBy = &req.LastModifiedBy

	if err := s.repo.UpdateSupplier(panelPricing); err != nil {
		return nil, err
	}

	return panelPricing, nil
}
