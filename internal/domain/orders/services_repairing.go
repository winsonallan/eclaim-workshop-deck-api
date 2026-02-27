package orders

import (
	"eclaim-workshop-deck-api/internal/models"
	"errors"
)

func (s *Service) GetRepairingOrders(workshopId uint) ([]models.Order, error) {
	return s.repo.GetRepairingOrders(workshopId)
}

func (s *Service) ExtendDeadline(req ExtendDeadlineRequest) (*models.Order, error) {
	if req.LastModifiedBy == 0 {
		return nil, errors.New("last_modified_by is needed")
	}
	if req.NewDeadline.IsZero() {
		return nil, errors.New("new_deadline is needed")
	}
	if req.OrderNo == 0 {
		return nil, errors.New("order_no is needed")
	}

	order, err := s.repo.FindOrderById(req.OrderNo)
	if err != nil {
		return nil, errors.New("order not found")
	}

	order.LastModifiedBy = &req.LastModifiedBy
	order.Eta = req.NewDeadline
	if req.Reason != nil && *req.Reason != "" {
		order.Notes = req.Reason
	}

	if err := s.repo.UpdateOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}
