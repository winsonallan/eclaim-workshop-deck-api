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

// Read
func (s *Service) GetOrders() ([]models.Order, error) {
	return s.repo.GetOrders()
}

func (s *Service) GetIncomingOrders(workshopId uint) ([]models.Order, error) {
	return s.repo.GetIncomingOrders(workshopId)
}

func (s *Service) ViewOrderDetails(orderNo uint) (models.Order, error) {
	return s.repo.ViewOrderDetails(orderNo)
}

// Create
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

func (s *Service) CreateWorkOrder(req CreateWorkOrderRequest) (*models.WorkOrder, error) {
	if req.CreatedBy == 0 {
		return nil, errors.New("created by is needed")
	}

	if len(req.OrderPanels) == 0 {
		return nil, errors.New("order panels are needed")
	}

	workOrder := &models.WorkOrder{
		OrderNo:                  req.OrderNo,
		CreatedBy:                &req.CreatedBy,
		AdditionalWorkOrderCount: 0,
	}

	if req.AdditionalWorkOrderCount != 0 {
		workOrder.AdditionalWorkOrderCount = req.AdditionalWorkOrderCount
	}

	if req.WorkOrderDocumentNumber != "" {
		workOrder.WorkOrderDocumentNumber = req.WorkOrderDocumentNumber
	}

	if req.WorkOrderUrl != "" {
		workOrder.WorkOrderUrl = req.WorkOrderUrl
	}

	if err := s.repo.CreateWorkOrder(workOrder); err != nil {
		return nil, err
	}

	var allPanels []*models.OrderPanel

	for _, o := range req.OrderPanels {
		orderPanel, err := s.prepareOrderPanels(o, req.CreatedBy, workOrder.WorkOrderNo)

		if err != nil {
			return nil, err
		}

		allPanels = append(allPanels, orderPanel)
	}

	if err := s.repo.CreateOrderPanelsBatch(allPanels); err != nil {
		return nil, err
	}

	return workOrder, nil
}

// Update
func (s *Service) ProposeAdditionalWork(req ProposeAdditionalWorkRequest) (*models.WorkOrder, error) {
	if req.LastModifiedBy == 0 {
		return nil, errors.New("last modified by is needed")
	}

	if len(req.OrderPanels) == 0 {
		return nil, errors.New("order panels are needed")
	}

	workOrder, err := s.repo.FindWorkOrderById(uint(req.WorkOrderNo))

	if err != nil {
		return nil, err
	}

	var currentWOGroup = &workOrder.AdditionalWorkOrderCount

	workOrder.AdditionalWorkOrderCount = *currentWOGroup + 1
	workOrder.LastModifiedBy = &req.LastModifiedBy

	var allPanels []*models.OrderPanel

	for _, o := range req.OrderPanels {
		orderPanel, err := s.prepareOrderPanels(o, req.LastModifiedBy, req.WorkOrderNo)
		orderPanel.NegotiationStatus = "proposed_additional"

		if err != nil {
			return nil, err
		}

		allPanels = append(allPanels, orderPanel)
	}

	if err := s.repo.CreateOrderPanelsBatch(allPanels); err != nil {
		return nil, err
	}

	if err := s.repo.UpdateWorkOrder(workOrder); err != nil {
		return nil, err
	}

	order, err := s.repo.ViewOrderDetails(workOrder.OrderNo)
	if err != nil {
		return nil, errors.New("order not found")
	}

	order.LastModifiedBy = &req.LastModifiedBy
	order.Status = "additional_work"

	if err := s.repo.UpdateOrder(&order); err != nil {
		return nil, err
	}

	return workOrder, nil
}

func (s *Service) AcceptOrder(id uint, req AcceptDeclineOrder) (*models.Order, error) {
	order, err := s.repo.ViewOrderDetails(id)
	if err != nil {
		return nil, errors.New("order not found")
	}

	workOrder := order.WorkOrders[0]
	groupNo := workOrder.AdditionalWorkOrderCount

	var orderPanels []models.OrderPanel

	for _, op := range workOrder.OrderPanels {
		if op.WorkOrderGroupNumber <= groupNo {
			orderPanels = append(orderPanels, op)
		}
	}

	for _, op := range orderPanels {
		if *op.InsurancePanelPricingNo != 0 {
			op.WorkshopPanelPricingNo = op.InsurancePanelPricingNo
			op.FinalPanelPricingNo = op.InsurancePanelPricingNo

			op.WorkshopPanelName = op.InsurancePanelName
			op.FinalPanelName = op.InsurancePanelName

			op.WorkshopPrice = op.InsurerPrice
			op.FinalPrice = op.InsurerPrice

			op.WorkshopServiceType = op.InsurerServiceType
			op.FinalServiceType = op.InsurerServiceType

			if op.InsurerMeasurementNo != nil && *op.InsurerMeasurementNo != 0 {
				op.WorkshopMeasurementNo = op.InsurerMeasurementNo
				op.FinalMeasurementNo = op.InsurerMeasurementNo
			}

			if op.InsurerQty != 0 {
				op.WorkshopQty = op.InsurerQty
				op.FinalQty = op.InsurerQty
			}
		} else if *op.WorkshopPanelPricingNo != 0 {
			op.InsurancePanelPricingNo = op.WorkshopPanelPricingNo
			op.FinalPanelPricingNo = op.WorkshopPanelPricingNo

			op.InsurancePanelName = op.WorkshopPanelName
			op.FinalPanelName = op.WorkshopPanelName

			op.InsurerPrice = op.WorkshopPrice
			op.FinalPrice = op.WorkshopPrice

			op.InsurerServiceType = op.WorkshopServiceType
			op.FinalServiceType = op.WorkshopServiceType

			if op.WorkshopMeasurementNo != nil && *op.WorkshopMeasurementNo != 0 {
				op.InsurerMeasurementNo = op.WorkshopMeasurementNo
				op.FinalMeasurementNo = op.WorkshopMeasurementNo
			}

			if op.WorkshopQty != 0 {
				op.InsurerQty = op.WorkshopQty
				op.FinalQty = op.WorkshopQty
			}
		}

		op.LastModifiedBy = &req.LastModifiedBy
		op.NegotiationStatus = "accepted"
		if err := s.repo.UpdateOrderPanel(&op); err != nil {
			return nil, err
		}
	}

	order.Status = "repairing"
	order.LastModifiedBy = &req.LastModifiedBy

	if err := s.repo.UpdateOrder(&order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *Service) DeclineOrder(id uint, req AcceptDeclineOrder) (*models.Order, error) {
	order, err := s.repo.ViewOrderDetails(id)
	if err != nil {
		return nil, errors.New("order not found")
	}

	workOrder := order.WorkOrders[0]
	groupNo := workOrder.AdditionalWorkOrderCount

	workOrder.IsLocked = true
	workOrder.LastModifiedBy = &req.LastModifiedBy

	if err := s.repo.UpdateWorkOrder(&workOrder); err != nil {
		return nil, err
	}

	var orderPanels []models.OrderPanel

	for _, op := range workOrder.OrderPanels {
		if op.WorkOrderGroupNumber <= groupNo {
			orderPanels = append(orderPanels, op)
		}
	}

	for _, op := range orderPanels {
		op.IsLocked = true
		op.NegotiationStatus = "rejected"
		op.LastModifiedBy = &req.LastModifiedBy

		if err := s.repo.UpdateOrderPanel(&op); err != nil {
			return nil, err
		}
	}

	order.Status = "declined"
	order.LastModifiedBy = &req.LastModifiedBy
	order.IsLocked = true
	if err := s.repo.UpdateOrder(&order); err != nil {
		return nil, err
	}

	return &order, nil
}
