package models

import (
    "sync"
)

type Task struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

type TaskManager struct {
    tasks map[string]Task
    mu    sync.Mutex
}

func NewTaskManager() *TaskManager {
    return &TaskManager{
        tasks: make(map[string]Task),
    }
}

func (tm *TaskManager) AddTask(task Task) {
    tm.mu.Lock()
    defer tm.mu.Unlock()
    tm.tasks[task.ID] = task
}

func (tm *TaskManager) GetTask(id string) (Task, bool) {
    tm.mu.Lock()
    defer tm.mu.Unlock()
    task, exists := tm.tasks[id]
    return task, exists
}

func (tm *TaskManager) UpdateTask(task Task) bool {
    tm.mu.Lock()
    defer tm.mu.Unlock()
    if _, exists := tm.tasks[task.ID]; !exists {
        return false
    }
    tm.tasks[task.ID] = task
    return true
}

func (tm *TaskManager) DeleteTask(id string) bool {
    tm.mu.Lock()
    defer tm.mu.Unlock()
    if _, exists := tm.tasks[id]; !exists {
        return false
    }
    delete(tm.tasks, id)
    return true
}

func (tm *TaskManager) GetAllTasks() []Task {
    tm.mu.Lock()
    defer tm.mu.Unlock()
    tasks := make([]Task, 0, len(tm.tasks))
    for _, task := range tm.tasks {
        tasks = append(tasks, task)
    }
    return tasks
}
