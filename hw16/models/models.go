package models

import (
    "database/sql"
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("record not found")

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
        return nil, fmt.Errorf("failed to open database connection: %w", err)
    }
    return &TaskManager{db: db}, nil
}

func (tm *TaskManager) AddTask(task Task) error {
    _, err := tm.db.Exec("INSERT INTO tasks (title, completed) VALUES ($1, $2)", task.Title, task.Completed)
    if err != nil {
        return fmt.Errorf("failed to add task: %w", err)
    }
    return nil
}

func (tm *TaskManager) GetTask(id int) (Task, error) {
    var task Task
    row := tm.db.QueryRow("SELECT id, title, completed FROM tasks WHERE id = $1", id)
    err := row.Scan(&task.ID, &task.Title, &task.Completed)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return task, fmt.Errorf("%w: task with id %d", ErrNotFound, id)
        }
        return task, fmt.Errorf("failed to get task with id %d: %w", id, err)
    }
    return task, nil
}

func (tm *TaskManager) UpdateTask(task Task) error {
    result, err := tm.db.Exec("UPDATE tasks SET title = $1, completed = $2 WHERE id = $3", task.Title, task.Completed, task.ID)
    if err != nil {
        return fmt.Errorf("failed to update task with id %d: %w", task.ID, err)
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("%w: task with id %d", ErrNotFound, task.ID)
    }
    return nil
}

func (tm *TaskManager) DeleteTask(id int) error {
    result, err := tm.db.Exec("DELETE FROM tasks WHERE id = $1", id)
    if err != nil {
        return fmt.Errorf("failed to delete task with id %d: %w", id, err)
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("%w: task with id %d", ErrNotFound, id)
    }
    return nil
}

func (tm *TaskManager) GetAllTasks() ([]Task, error) {
    rows, err := tm.db.Query("SELECT id, title, completed FROM tasks")
    if err != nil {
        return nil, fmt.Errorf("failed to get all tasks: %w", err)
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
            return nil, fmt.Errorf("failed to scan task: %w", err)
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}
