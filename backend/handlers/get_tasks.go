package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

// GetTasks: 指定されたユーザーのタスクを取得するハンドラ関数
func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			log.Printf("ユーザーIDが指定されていません")
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT id, user_id, title, purpose, deadline, priority, status, steps, memo, remarks FROM tasks WHERE user_id = ?", userID)
		if err != nil {
			log.Printf("タスクの取得に失敗しました: %v", err)
			http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tasks []models.Task
		for rows.Next() {
			var task models.Task
			if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Purpose, &task.Deadline, &task.Priority, &task.Status, &task.Steps, &task.Memo, &task.Remarks); err != nil {
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
