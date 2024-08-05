package service

import (
	"time"
	"log"
	"travel-agency/internal/models"
	"travel-agency/internal/repository"
)

type TourService struct {
	repo *repository.TourRepository
}

func NewTourService(repo *repository.TourRepository) *TourService {
	return &TourService{repo: repo}
}

func (s *TourService) ListAvailableTours() []models.Tour {
	return s.repo.GetAvailableTours()
}

func (s *TourService) OrderTour(tourID, userID int) (models.TourOrder, error) {
	tour, err := s.repo.GetTourByID(tourID)
	if err != nil {
		return models.TourOrder{}, err
	}

	order := models.TourOrder{
		TourID:  tour.ID,
		UserID:  userID,
		Status:  "future",
		Ordered: time.Now(),
	}

	s.repo.SaveOrder(order)
	log.Printf("Email sent: Order for tour %s confirmed.\n", tour.Name)

	return order, nil
}

func (s *TourService) ListUserOrders(userID int) []models.TourOrder {
	orders := s.repo.GetOrdersByUserID(userID)
	for i, order := range orders {
		tour, _ := s.repo.GetTourByID(order.TourID)
		now := time.Now()
		if now.Before(tour.StartDate) {
			orders[i].Status = "future"
		} else if now.After(tour.EndDate) {
			orders[i].Status = "completed"
		} else if now.Before(tour.StartDate.Add(-time.Hour * 24 * 3)) {
			orders[i].Status = "soon"
		} else {
			orders[i].Status = "ongoing"
		}
	}
	return orders
}