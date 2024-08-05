package handler

import (
	"encoding/json"
	"main/hw10/models"
	"net/http"
	"github.com/gorilla/mux"
)

type TaskHandlers struct {
    Manager *models.TaskManager
}

func NewTaskHandlers(manager *models.TaskManager) *TaskHandlers {
    return &TaskHandlers{Manager: manager}
}

func (h *TaskHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    h.Manager.AddTask(task)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    task, exists := h.Manager.GetTask(id)
    if !exists {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
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

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedTask)
}

func (h *TaskHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    if success := h.Manager.DeleteTask(id); !success {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
