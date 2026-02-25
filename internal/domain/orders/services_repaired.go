package orders

import "eclaim-workshop-deck-api/internal/models"

func (s *Service) GetRepairedOrders(workshopId uint) ([]models.Order, error) {
	return s.repo.GetRepairedOrders(workshopId)
}
