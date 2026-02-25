package orders

import "eclaim-workshop-deck-api/internal/models"

func (s *Service) GetDeliveredOrders(workshopId uint) ([]models.Order, error) {
	return s.repo.GetDeliveredOrders(workshopId)
}
