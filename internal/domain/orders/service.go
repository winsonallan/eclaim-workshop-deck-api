package orders

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"
	"fmt"
	"mime/multipart"

	"gorm.io/gorm"
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

func (s *Service) GetNegotiatingOrders(workshopId uint) ([]models.Order, error) {
	return s.repo.GetNegotiatingOrders(workshopId)
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

func (s *Service) SubmitNegotiation(
	req *SubmitNegotiationRequest,
	files []*multipart.FileHeader,
	uploadFn func(file multipart.File, header *multipart.FileHeader, folder string) (string, error),
) (*models.WorkOrder, error) {
	if req.WorkOrderNo == 0 {
		return nil, errors.New("work_order_no is required")
	}
	if len(req.OrderPanels) == 0 {
		return nil, errors.New("order_panels cannot be empty")
	}
	if req.Reason == "" {
		return nil, errors.New("reason is required")
	}

	if len(req.Photos) != len(files) {
		return nil, fmt.Errorf("file count mismatch: expected %d, got %d", len(req.Photos), len(files))
	}

	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/webp": true,
	}
	maxSize := int64(10 << 20)

	for _, fh := range files {
		if fh.Size > maxSize {
			return nil, fmt.Errorf("file %s exceeds 10MB", fh.Filename)
		}
		contentType := fh.Header.Get("Content-Type")
		if !allowedTypes[contentType] {
			return nil, fmt.Errorf("invalid file type %s for %s", contentType, fh.Filename)
		}
	}

	workOrder, err := s.repo.FindWorkOrderById(req.WorkOrderNo)
	if err != nil {
		return nil, err
	}

	if workOrder == nil {
		return nil, errors.New("work order not found")
	}

	groupCount := workOrder.AdditionalWorkOrderCount

	order, err := s.repo.FindOrderById(workOrder.OrderNo)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order not found")
	}

	order.Status = "negotiating"
	order.LastModifiedBy = &req.LastModifiedBy

	// Build photo mapping: order_panel_no -> file headers
	photosByPanel := make(map[uint][]*multipart.FileHeader)
	for _, photoMeta := range req.Photos {
		if photoMeta.FileIndex >= 0 && photoMeta.FileIndex < len(files) {
			photosByPanel[photoMeta.OrderPanelNo] = append(
				photosByPanel[photoMeta.OrderPanelNo],
				files[photoMeta.FileIndex],
			)
		}
	}

	negotiationsCreated := make([]NegotiationCreatedInfo, 0)

	err = s.repo.WithTransaction(func(tx *gorm.DB) error {
		for _, panelReq := range req.OrderPanels {
			orderPanel, err := s.repo.GetOrderPanelWithLock(tx, panelReq.OrderPanelNo)
			if err != nil {
				return fmt.Errorf("failed to lock order panel %d: %w", panelReq.OrderPanelNo, err)
			}

			// Validate order panel is in correct state
			if orderPanel.NegotiationStatus != "pending_workshop" {
				return fmt.Errorf("order panel %d is not pending workshop action (current status: %s)",
					panelReq.OrderPanelNo, orderPanel.NegotiationStatus)
			}

			// Calculate new round number
			newRound := orderPanel.CurrentRound + 1

			// Create negotiation history entry
			negotiationHistory := &models.NegotiationHistory{
				OrderPanelNo:      panelReq.OrderPanelNo,
				RoundCount:        newRound,
				OldPanelPricingNo: orderPanel.InsurancePanelPricingNo,
				OldPrice:          orderPanel.InsurerPrice,
				OldMeasurementNo:  orderPanel.InsurerMeasurementNo,
				OldServiceType:    orderPanel.InsurerServiceType,
				OldQty:            orderPanel.InsurerQty,

				ProposedPanelPricingNo: panelReq.WorkshopPanelPricingNo,
				ProposedPrice:          panelReq.WorkshopPrice,
				ProposedServiceType:    panelReq.WorkshopServiceType,
				ProposedQty:            panelReq.WorkshopQty,
				WorkshopNotes:          req.Reason,
				InsuranceDecision:      "pending",
				CreatedBy:              &req.LastModifiedBy,
			}

			if panelReq.WorkshopMeasurementNo != 0 {
				negotiationHistory.ProposedMeasurementNo = &panelReq.WorkshopMeasurementNo
			}

			err = s.repo.CreateNegotiationHistory(tx, negotiationHistory)
			if err != nil {
				return fmt.Errorf("failed to create negotiation history for panel %d: %w",
					panelReq.OrderPanelNo, err)
			}

			// Update order panel with workshop proposal
			orderPanel.CurrentRound = newRound
			orderPanel.NegotiationStatus = "negotiating"
			orderPanel.WorkshopPanelPricingNo = &panelReq.WorkshopPanelPricingNo
			orderPanel.WorkshopPanelName = panelReq.WorkshopPanelName
			orderPanel.WorkshopPrice = &negotiationHistory.ProposedPrice
			orderPanel.WorkshopServiceType = negotiationHistory.ProposedServiceType
			orderPanel.WorkshopQty = &panelReq.WorkshopQty
			orderPanel.IsIncluded = panelReq.IsIncluded
			orderPanel.IsSpecialRepair = panelReq.IsSpecialRepair
			orderPanel.LastModifiedBy = &req.LastModifiedBy

			if panelReq.WorkshopMeasurementNo != 0 {
				orderPanel.WorkshopMeasurementNo = &panelReq.WorkshopMeasurementNo
			}

			err = s.repo.UpdateOrderPanelTx(tx, orderPanel)
			if err != nil {
				return fmt.Errorf("failed to update order panel %d: %w", panelReq.OrderPanelNo, err)
			}

			// Upload photos and create photo records
			photosUploaded := 0
			panelPhotos := photosByPanel[panelReq.OrderPanelNo]

			if len(panelPhotos) > 0 {
				photoRecords := make([]models.NegotiationPhotos, 0)

				for _, fileHeader := range panelPhotos {
					// Open file
					file, err := fileHeader.Open()
					if err != nil {
						return fmt.Errorf("failed to open file %s: %w", fileHeader.Filename, err)
					}

					// Upload to storage
					photoURL, err := uploadFn(file, fileHeader, "negotiations")
					file.Close()

					if err != nil {
						return fmt.Errorf("failed to upload file %s: %w", fileHeader.Filename, err)
					}

					// Create photo record
					photoRecords = append(photoRecords, models.NegotiationPhotos{
						NegotiationHistoryNo: negotiationHistory.NegotiationHistoryNo,
						PhotoUrl:             photoURL,
						CreatedBy:            &req.LastModifiedBy,
					})

					photosUploaded++
				}

				// Save all photo records
				if len(photoRecords) > 0 {
					err = s.repo.CreateNegotiationPhotos(tx, photoRecords)
					if err != nil {
						return fmt.Errorf("failed to save photo records for panel %d: %w",
							panelReq.OrderPanelNo, err)
					}
				}
			}

			negotiationsCreated = append(negotiationsCreated, NegotiationCreatedInfo{
				OrderPanelNo:         panelReq.OrderPanelNo,
				NegotiationHistoryNo: negotiationHistory.NegotiationHistoryNo,
				RoundCount:           newRound,
				PhotosUploaded:       photosUploaded,
			})
		}

		if groupCount > 0 {
			if err := s.repo.BulkAcceptPanelsByGroupRangeTx(tx, workOrder.WorkOrderNo, 0, groupCount, req.LastModifiedBy); err != nil {
				return fmt.Errorf("failed to bulk accept previous panels: %w", err)
			}
		}

		err = s.repo.UpdateOrderTx(tx, order)
		if err != nil {
			return fmt.Errorf("failed to update order status for work order %d: %w", req.WorkOrderNo, err)
		}

		return nil
	})

	if err != nil {
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

	fmt.Println("Order LastModifiedBy:", req.LastModifiedBy)
	fmt.Println("WorkOrder LastModifiedBy:", workOrder.LastModifiedBy)

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

			op.WorkshopPrice = &op.InsurerPrice
			op.FinalPrice = &op.InsurerPrice

			op.WorkshopServiceType = op.InsurerServiceType
			op.FinalServiceType = op.InsurerServiceType

			if op.InsurerMeasurementNo != nil && *op.InsurerMeasurementNo != 0 {
				op.WorkshopMeasurementNo = op.InsurerMeasurementNo
				op.FinalMeasurementNo = op.InsurerMeasurementNo
			}

			if op.InsurerQty != 0 {
				op.WorkshopQty = &op.InsurerQty
				op.FinalQty = &op.InsurerQty
			}
		} else if *op.WorkshopPanelPricingNo != 0 {
			op.InsurancePanelPricingNo = op.WorkshopPanelPricingNo
			op.FinalPanelPricingNo = op.WorkshopPanelPricingNo

			op.InsurancePanelName = op.WorkshopPanelName
			op.FinalPanelName = op.WorkshopPanelName

			op.InsurerPrice = *op.WorkshopPrice
			op.FinalPrice = op.WorkshopPrice

			op.InsurerServiceType = op.WorkshopServiceType
			op.FinalServiceType = op.WorkshopServiceType

			if op.WorkshopMeasurementNo != nil && *op.WorkshopMeasurementNo != 0 {
				op.InsurerMeasurementNo = op.WorkshopMeasurementNo
				op.FinalMeasurementNo = op.WorkshopMeasurementNo
			}

			if *op.WorkshopQty != 0 {
				op.InsurerQty = *op.WorkshopQty
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

	if !(req.ETA.IsZero()) {
		order.Eta = req.ETA
	}

	if req.DiscountType != "" {
		order.DiscountType = req.DiscountType
		order.Discount = req.Discount
	}

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
	fmt.Println("Order LastModifiedBy:", req.LastModifiedBy)
	fmt.Println("WorkOrder LastModifiedBy:", workOrder.LastModifiedBy)

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
