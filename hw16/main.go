package main

import (
    "log"
    "main/hw16/handler"
    "main/hw16/models"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    connStr := "user=user password=password dbname=taskdb sslmode=disable"
    taskManager, err := models.NewTaskManager(connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
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