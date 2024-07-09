package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

// GetTasks は指定されたユーザーのタスクを取得するハンドラ関数です
func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT title, deadline, priority, status FROM tasks WHERE user_id = ?", userID)
		if err != nil {
			log.Printf("タスクの取得に失敗しました: %v", err)
			http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tasks []models.Task
		for rows.Next() {
			var task models.Task
			if err := rows.Scan(&task.Title, &task.Deadline, &task.Priority, &task.Status); err != nil {
				log.Printf("タスクのスキャンに失敗しました: %v", err)
				http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			log.Printf("タスクのエンコードに失敗しました: %v", err)
			http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
		}
	}
}
