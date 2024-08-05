package main

import (
	"log"
	"github.com/gorilla/mux"
        "main/hw10/handler"
	"main/hw10/models"
	"net/http"
)

func main() {
	taskManager := models.NewTaskManager()
    taskHandlers := handler.NewTaskHandlers(taskManager)

    r := mux.NewRouter()
    r.HandleFunc("/tasks", taskHandlers.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", taskHandlers.GetTask).Methods("GET")
    r.HandleFunc("/tasks/{id}", taskHandlers.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", taskHandlers.DeleteTask).Methods("DELETE")
    r.HandleFunc("/tasks", taskHandlers.GetAllTasks).Methods("GET")
	
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
