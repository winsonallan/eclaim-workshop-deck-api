package orders

import "eclaim-workshop-deck-api/internal/models"

func (s *Service) GetRepairingOrders(workshopId uint) ([]models.Order, error) {
	return s.repo.GetRepairingOrders(workshopId)
}
