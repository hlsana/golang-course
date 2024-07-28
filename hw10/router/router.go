package router

import (
    "github.com/gorilla/mux"
    "main/hw10/handler"
)

func InitializeRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
    r.HandleFunc("/tasks", handler.AddTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
    return r
}

