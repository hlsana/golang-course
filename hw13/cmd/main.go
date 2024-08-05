package main

import (
	"net/http"
	"travel-agency/internal/handler"
	"travel-agency/internal/repository"
	"travel-agency/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewTourRepository()
	svc := service.NewTourService(repo)
	handler := handler.NewTourHandler(svc)

	router := mux.NewRouter()
	router.HandleFunc("/tours", handler.ListAvailableTours).Methods("GET")
	router.HandleFunc("/order/{id:[0-9]+}", handler.OrderTour).Methods("POST")
	router.HandleFunc("/orders", handler.ListUserOrders).Methods("GET")

	http.ListenAndServe(":8080", router)
}