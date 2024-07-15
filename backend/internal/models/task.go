package models

import (
	"database/sql"
	"fmt"
)

// Task represents a task in the system
type Task struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	Title    string `json:"title"`
	Purpose  string `json:"purpose"`
	Deadline string `json:"deadline"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	Steps    string `json:"steps"`
	Memo     string `json:"memo"`
	Remarks  string `json:"remarks"`
}

// データベースにタスクを保存する
func CreateTask(db *sql.DB, task Task) (int64, error) {
	query := `
        INSERT INTO tasks (user_id, title, deadline, priority, status, purpose, steps, memo, remarks)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	result, err := db.Exec(query, task.UserID, task.Title, task.Deadline, task.Priority, task.Status, task.Purpose, task.Steps, task.Memo, task.Remarks)
	if err != nil {
		return 0, fmt.Errorf("タスクの保存に失敗しました: %v", err)
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("挿入されたタスクのID取得に失敗しました: %v", err)
	}

	return taskID, nil
}

// データベースからユーザーIDに紐づくタスクを取得する
func GetTasksByUserID(db *sql.DB, userID string) ([]Task, error) {
	rows, err := db.Query("SELECT id, user_id, title, purpose, deadline, priority, status, steps, memo, remarks FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("タスクの取得に失敗しました: %v", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Purpose, &task.Deadline, &task.Priority, &task.Status, &task.Steps, &task.Memo, &task.Remarks); err != nil {
			return nil, fmt.Errorf("タスクのスキャンに失敗しました: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// データベースからタスクIDに紐づくタスクを取得する
func GetTaskByID(db *sql.DB, taskID string) (Task, error) {
	var task Task
	err := db.QueryRow("SELECT id, user_id, title, deadline, priority, status, purpose, steps, memo, remarks FROM tasks WHERE id = ?", taskID).Scan(
		&task.ID, &task.UserID, &task.Title, &task.Deadline, &task.Priority, &task.Status, &task.Purpose, &task.Steps, &task.Memo, &task.Remarks)
	if err != nil {
		if err == sql.ErrNoRows {
			return task, nil
		}
		return task, fmt.Errorf("タスクの取得に失敗しました: %v", err)
	}

	return task, nil
}
