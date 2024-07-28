package models

type Task struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

var TaskList = make(map[string]Task)