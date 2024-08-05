package models

import "time"

type Tour struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Transport   string    `json:"transport"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type TourOrder struct {
	ID      int    `json:"id"`
	TourID  int    `json:"tour_id"`
	UserID  int    `json:"user_id"`
	Status  string `json:"status"`
	Ordered time.Time `json:"ordered"`
}