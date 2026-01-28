package settings

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

func (s *Service) GetAccount(id uint) ([]models.User, error) {
	return s.repo.GetAccount(id)
}

func (s *Service) GetProfileDetails(id uint) (*models.UserProfile, error) {
	return s.repo.GetProfileDetails(id)
}

func (s *Service) GetWorkshopDetails(id uint) ([]models.WorkshopDetails, error) {
	return s.repo.GetWorkshopDetails(id)
}

func (s *Service) GetWorkshopPICs(id uint) ([]models.WorkshopPics, error) {
	return s.repo.GetWorkshopPICs(id)
}

func (s *Service) CreateWorkshopDetails(req CreateWorkshopDetailsRequest) (*models.WorkshopDetails, error) {
	if req.ProfileNo == 0 {
		return nil, errors.New("User Profile Number is required")
	}
	if !req.IsAuthorized {
		return nil, errors.New("Is Authorized is required")
	}

	workshopDetails := &models.WorkshopDetails{
		UserProfileNo: req.ProfileNo,
		Capacity:      req.Capacity,
		Description:   req.Description,
		IsAuthorized:  req.IsAuthorized,
		Specialist:    req.Specialist,
		CreatedBy:     &req.CreatedBy,
	}

	if err := s.repo.CreateWorkshopDetails(workshopDetails); err != nil {
		return nil, err
	}

	return s.repo.FindWorkshopDetailsByID(workshopDetails.WorkshopDetailsNo)
}

func (s *Service) CreateWorkshopPIC(req CreateWorkshopPICRequest) (*models.WorkshopPics, error) {
	if req.WorkshopDetailsNo == 0 {
		return nil, errors.New("Workshop Details No is required")
	}

	if req.PicName == "" {
		return nil, errors.New("Name is required")
	}

	if req.PicTitle == "" {
		return nil, errors.New("Title is required")
	}

	if req.Phone == "" {
		return nil, errors.New("Phone is required")
	}

	if req.Email == "" {
		return nil, errors.New("Email is required")
	}

	workshopPics := &models.WorkshopPics{
		WorkshopDetailsNo: req.WorkshopDetailsNo,
		WorkshopPicName:   req.PicName,
		WorkshopPicTitle:  req.PicTitle,
		Phone:             req.Phone,
		Email:             req.Email,
		CreatedBy:         &req.CreatedBy,
	}

	if err := s.repo.CreateWorkshopPICs(workshopPics); err != nil {
		return nil, err
	}

	return s.repo.FindWorkshopPICByID(workshopPics.WorkshopPicNo)
}

func (s *Service) UpdateWorkshopDetails(workshopDetailsNo uint, req UpdateWorkshopDetailsRequest) (*models.WorkshopDetails, error) {
	workshopDetails, err := s.repo.FindWorkshopDetailsByID(workshopDetailsNo)
	if err != nil {
		return nil, errors.New("workshop details not found")
	}

	userProfileNo := workshopDetails.UserProfileNo
	userProfile, err := s.repo.GetProfileDetails(userProfileNo)
	if err != nil {
		return nil, errors.New("user profile not found")
	}

	if req.Address != "" {
		userProfile.UserProfileAddress = req.Address
	}

	if req.Capacity != 0 {
		workshopDetails.Capacity = req.Capacity
	}

	if req.CityType != "" {
		userProfile.UserProfileCityType = req.CityType
	}
	if req.CityName != "" {
		userProfile.UserProfileCityName = req.CityName
	}

	if req.CityNo != 0 {
		userProfile.UserProfileCityNo = req.CityNo
	}

	if req.Description != "" {
		workshopDetails.Description = req.Description
	}

	if req.Email != "" {
		userProfile.UserProfileEmail = req.Email
	}

	if req.Phone != "" {
		userProfile.UserProfilePhone = req.Phone
	}

	if req.WorkshopName != "" {
		userProfile.UserProfileName = req.WorkshopName
	}

	userProfile.LastModifiedBy = &req.LastModifiedBy
	workshopDetails.LastModifiedBy = &req.LastModifiedBy

	if err := s.repo.UpdateWorkshopDetails(workshopDetails); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateUserProfile(userProfile); err != nil {
		return nil, err
	}

	return workshopDetails, nil
}

func (s *Service) UpdateWorkshopPIC(workshopPICNo uint, req UpdateWorkshopPICRequest) (*models.WorkshopPics, error) {

	workshopPIC, err := s.repo.FindWorkshopPICByID(workshopPICNo)

	if err != nil {
		return nil, errors.New("workshop PIC not found")
	}

	if req.WorkshopPicName != "" {
		workshopPIC.WorkshopPicName = req.WorkshopPicName
	}

	if req.WorkshopPicTitle != "" {
		workshopPIC.WorkshopPicTitle = req.WorkshopPicTitle
	}

	if req.Phone != "" {
		workshopPIC.Phone = req.Phone
	}

	if req.Email != "" {
		workshopPIC.Email = req.Email
	}

	workshopPIC.LastModifiedBy = &req.LastModifiedBy

	if err := s.repo.UpdateWorkshopPIC(workshopPIC); err != nil {
		return nil, err
	}

	return workshopPIC, nil
}

// Delete
func (s *Service) DeleteWorkshopPIC(workshopPICNo uint, req DeleteWorkshopPICRequest) (*models.WorkshopPics, error) {
	workshopPIC, err := s.repo.FindWorkshopPICByID(workshopPICNo)

	if err != nil {
		return nil, errors.New("workshop PIC not found")
	}

	workshopPIC.IsLocked = true
	workshopPIC.LastModifiedBy = &req.LastModifiedBy

	if err := s.repo.UpdateWorkshopPIC(workshopPIC); err != nil {
		return nil, err
	}

	return workshopPIC, nil
}
