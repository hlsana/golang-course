package repository

import (
	"errors"
	"sync"
	"time"
	"travel-agency/internal/models"
)

var (
	ErrTourNotFound = errors.New("tour not found")
)

type TourRepository struct {
	mu          sync.Mutex
	tours       []models.Tour
	orders      []models.TourOrder
	nextTourID  int
	nextOrderID int
}

func NewTourRepository() *TourRepository {
	return &TourRepository{
		tours: []models.Tour{
			{ID: 1, Name: "Kyiv-Odesa", Description: "A beach holiday for your soul.", Price: 499.99, Transport: "Plane", StartDate: time.Now().AddDate(0, 1, 0), EndDate: time.Now().AddDate(0, 1, 7)},
			{ID: 2, Name: "Kyiv-Bukovel", Description: "An meditative mountain trek.", Price: 299.99, Transport: "Train", StartDate: time.Now().AddDate(0, 2, 0), EndDate: time.Now().AddDate(0, 2, 10)},
		},
		nextTourID:  3,
		nextOrderID: 1,
	}
}

func (r *TourRepository) GetAvailableTours() []models.Tour {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.tours
}

func (r *TourRepository) GetTourByID(id int) (models.Tour, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, t := range r.tours {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Tour{}, ErrTourNotFound
}

func (r *TourRepository) SaveOrder(order models.TourOrder) models.TourOrder {
	r.mu.Lock()
	defer r.mu.Unlock()

	order.ID = r.nextOrderID
	r.nextOrderID++
	r.orders = append(r.orders, order)
	return order
}

func (r *TourRepository) GetOrdersByUserID(userID int) []models.TourOrder {
	r.mu.Lock()
	defer r.mu.Unlock()

	var orders []models.TourOrder
	for _, o := range r.orders {
		if o.UserID == userID {
			orders = append(orders, o)
		}
	}
	return orders
}
