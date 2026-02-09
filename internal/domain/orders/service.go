package orders

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"
)

type Service struct {
	repo      *Repository
	jwtSecret string
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetOrders() ([]models.Order, error) {
	return s.repo.GetOrders()
}

func (s *Service) prepareClient(req AddClientRequest) (*models.Client, error) {
	var client *models.Client

	if req.ClientName == "" {
		return nil, errors.New("client name is required")
	}
	if req.ClientPhone == "" {
		return nil, errors.New("client email is required")
	}
	if req.CityNo == 0 {
		return nil, errors.New("city no is required")
	}
	if req.CityName == "" {
		return nil, errors.New("city name is required")
	}
	if req.VehicleBrandName == "" {
		return nil, errors.New("vehicle brand is required")
	}
	if req.VehicleSeriesName == "" {
		return nil, errors.New("vehicle series name is required")
	}
	if req.VehicleChassisNo == "" {
		return nil, errors.New("vehicle chassis no is required")
	}
	if req.VehicleLicensePlate == "" {
		return nil, errors.New("vehicle license plate is required")
	}
	if req.VehiclePrice == 0 {
		return nil, errors.New("vehicle price is required")
	}

	client = &models.Client{
		ClientName:          req.ClientName,
		ClientPhone:         req.ClientPhone,
		CityNo:              req.CityNo,
		CityType:            req.CityType,
		CityName:            req.CityName,
		Address:             req.Address,
		VehicleBrandName:    req.VehicleBrandName,
		VehicleSeriesName:   req.VehicleSeriesName,
		VehicleChassisNo:    req.VehicleChassisNo,
		VehicleLicensePlate: req.VehicleLicensePlate,
		VehiclePrice:        req.VehiclePrice,
	}

	if req.ClientEmail != "" {
		client.ClientEmail = req.ClientEmail
	}

	return client, nil
}

func (s *Service) AddClient(req AddClientRequest) (*models.Client, error) {
	client, err := s.prepareClient(req)
	if err != nil {
		return nil, err
	}

	if err := s.repo.AddClient(client); err != nil {
		return nil, err
	}

	return s.repo.FindClientById(client.ClientNo)
}

func (s *Service) CreateOrder(req CreateOrderRequest) (*models.Order, error) {
	var orderType string
	if req.InsuranceNo != 0 {
		orderType = "insurance"
	} else {
		orderType = "manual"
	}

	var clientNo uint
	if req.ClientNo != 0 {
		clientNo = req.ClientNo
	} else {
		client, err := s.prepareClient(*req.ClientDetails)
		if err != nil {
			return nil, err
		}
		client.CreatedBy = &req.CreatedBy
		if err := s.repo.AddClient(client); err != nil {
			return nil, err
		}

		clientNo = client.ClientNo
	}

	if req.WorkshopNo == 0 {
		return nil, errors.New("workshop no is required")
	}
	if req.ClaimDetails == "" {
		return nil, errors.New("claim details is required")
	}
	if req.CreatedBy == 0 {
		return nil, errors.New("created by is required")
	}
	if req.Status == "" {
		return nil, errors.New("status is required")
	}

	order := &models.Order{
		WorkshopNo:   req.WorkshopNo,
		OrderType:    orderType,
		ClaimDetails: req.ClaimDetails,
		ClientNo:     clientNo,
		CreatedBy:    &req.CreatedBy,
		Status:       req.Status,
	}

	if req.InsuranceNo != 0 {
		order.InsuranceNo = &req.InsuranceNo
	}

	if !req.ETA.IsZero() {
		order.Eta = req.ETA
	}

	if err := s.repo.CreateOrder(order); err != nil {
		return nil, err
	}

	return s.repo.FindOrderById(order.OrderNo)
}
