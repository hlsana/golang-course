package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"travel-agency/internal/service"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	service *service.TourService
}

func NewTourHandler(s *service.TourService) *TourHandler {
	return &TourHandler{service: s}
}

func (h *TourHandler) ListAvailableTours(w http.ResponseWriter, r *http.Request) {
	tours := h.service.ListAvailableTours()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tours)
}

func (h *TourHandler) OrderTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tourID, _ := strconv.Atoi(vars["id"])
	userID := 1

	order, err := h.service.OrderTour(tourID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *TourHandler) ListUserOrders(w http.ResponseWriter, r *http.Request) {
	userID := 1

	orders := h.service.ListUserOrders(userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}