package orders

import "eclaim-workshop-deck-api/internal/models"

func (s *Service) GetNegotiatingOrders(workshopId uint) ([]models.Order, error) {
	return s.repo.GetNegotiatingOrders(workshopId)
}
