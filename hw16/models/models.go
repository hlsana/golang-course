package models

import (
    "database/sql"
    _ "github.com/lib/pq"
)

type Task struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

type TaskManager struct {
    db *sql.DB
}

func NewTaskManager(dataSourceName string) (*TaskManager, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            completed BOOLEAN NOT NULL
        )
    `)
    if err != nil {
        return nil, err
    }

    return &TaskManager{db: db}, nil
}

func (tm *TaskManager) AddTask(task Task) error {
    _, err := tm.db.Exec("INSERT INTO tasks (title, completed) VALUES ($1, $2)", task.Title, task.Completed)
    return err
}

func (tm *TaskManager) GetTask(id int) (Task, error) {
    var task Task
    row := tm.db.QueryRow("SELECT id, title, completed FROM tasks WHERE id = $1", id)
    err := row.Scan(&task.ID, &task.Title, &task.Completed)
    return task, err
}

func (tm *TaskManager) UpdateTask(task Task) error {
    _, err := tm.db.Exec("UPDATE tasks SET title = $1, completed = $2 WHERE id = $3", task.Title, task.Completed, task.ID)
    return err
}

func (tm *TaskManager) DeleteTask(id int) error {
    _, err := tm.db.Exec("DELETE FROM tasks WHERE id = $1", id)
    return err
}

func (tm *TaskManager) GetAllTasks() ([]Task, error) {
    rows, err := tm.db.Query("SELECT id, title, completed FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}
