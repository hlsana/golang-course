package handler

import (
	"context"
	"encoding/json"
	"main/hw17/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/go-redis/redis/v8"
)

type TaskHandlers struct {
	Manager *models.TaskManager
	Cache   *redis.Client
}

func NewTaskHandlers(manager *models.TaskManager, cache *redis.Client) *TaskHandlers {
	return &TaskHandlers{
		Manager: manager,
		Cache:   cache,
	}
}

func (h *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	h.Manager.AddTask(task)

	h.Cache.Set(context.Background(), task.ID, task, 10*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	taskJSON, err := h.Cache.Get(context.Background(), id).Result()
	if err == redis.Nil {
		task, exists := h.Manager.GetTask(id)
		if !exists {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		taskJSON, _ := json.Marshal(task)
		h.Cache.Set(context.Background(), id, taskJSON, 10*time.Minute)

		w.Header().Set("Content-Type", "application/json")
		w.Write(taskJSON)
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(taskJSON))
	}
}

func (h *TaskHandlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.Manager.GetAllTasks()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedTask.ID = id

	if success := h.Manager.UpdateTask(updatedTask); !success {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	taskJSON, _ := json.Marshal(updatedTask)
	h.Cache.Set(context.Background(), id, taskJSON, 10*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.Write(taskJSON)
}

func (h *TaskHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if success := h.Manager.DeleteTask(id); !success {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	h.Cache.Del(context.Background(), id)

	w.WriteHeader(http.StatusNoContent)
}
