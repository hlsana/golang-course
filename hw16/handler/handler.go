package handler

import (
    "encoding/json"
    "main/hw16/models"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
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

    err := h.Manager.AddTask(task)
    if err != nil {
        http.Error(w, "Unable to create task", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    task, err := h.Manager.GetTask(id)
    if err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func (h *TaskHandlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
    tasks, err := h.Manager.GetAllTasks()
    if err != nil {
        http.Error(w, "Unable to fetch tasks", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    var updatedTask models.Task
    if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    updatedTask.ID = id

    err := h.Manager.UpdateTask(updatedTask)
    if err != nil {
        http.Error(w, "Unable to update task", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedTask)
}

func (h *TaskHandlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    err := h.Manager.DeleteTask(id)
    if err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
