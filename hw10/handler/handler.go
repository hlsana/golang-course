package handler

import (
	"encoding/json"
	"fmt"
	"main/hw10/models"
	"net/http"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	for _, task := range models.TaskList {
		tasks = append(tasks, task)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = generateID()
	models.TaskList[task.ID] = task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask models.Task
	json.NewDecoder(r.Body).Decode(&updatedTask)
	if task, exists := models.TaskList[id]; exists {
		task.Title = updatedTask.Title
		task.Completed = updatedTask.Completed
		models.TaskList[id] = task
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, exists := models.TaskList[id]; exists {
		delete(models.TaskList, id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

func generateID() string {
	return fmt.Sprintf("%d", len(models.TaskList)+1)
}
