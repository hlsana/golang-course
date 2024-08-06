package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"main/hw17/handler"
	"main/hw17/models"
	"github.com/go-redis/redis/v8"
)

func main() {
	taskManager := models.NewTaskManager()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	taskHandlers := handler.NewTaskHandlers(taskManager, rdb)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandlers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandlers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks", taskHandlers.GetAllTasks).Methods("GET")

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
