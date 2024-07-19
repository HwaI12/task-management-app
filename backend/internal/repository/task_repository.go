package repository

import (
	"database/sql"
	"fmt"

	"github.com/HwaI12/task-management-app/backend/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task model.Task) (int64, error) {
	query := `
        INSERT INTO tasks (user_id, title, deadline, priority, status, purpose, steps, memo, remarks)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	result, err := r.db.Exec(query, task.UserID, task.Title, task.Deadline, task.Priority, task.Status, task.Purpose, task.Steps, task.Memo, task.Remarks)
	if err != nil {
		return 0, fmt.Errorf("タスクの保存に失敗しました: %v", err)
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("挿入されたタスクのID取得に失敗しました: %v", err)
	}

	return taskID, nil
}

func (r *TaskRepository) GetByUserID(userID string) ([]model.Task, error) {
	query := "SELECT id, user_id, title, purpose, deadline, priority, status, steps, memo, remarks FROM tasks WHERE user_id = ?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("タスクの取得に失敗しました: %v", err)
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Purpose, &task.Deadline, &task.Priority, &task.Status, &task.Steps, &task.Memo, &task.Remarks); err != nil {
			return nil, fmt.Errorf("タスクのスキャンに失敗しました: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) GetByID(taskID string) (model.Task, error) {
	var task model.Task
	query := "SELECT id, user_id, title, deadline, priority, status, purpose, steps, memo, remarks FROM tasks WHERE id = ?"
	err := r.db.QueryRow(query, taskID).Scan(
		&task.ID, &task.UserID, &task.Title, &task.Deadline, &task.Priority, &task.Status, &task.Purpose, &task.Steps, &task.Memo, &task.Remarks)
	if err != nil {
		if err == sql.ErrNoRows {
			return task, nil
		}
		return task, fmt.Errorf("タスクの取得に失敗しました: %v", err)
	}

	return task, nil
}
