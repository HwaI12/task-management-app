package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/models"
	"github.com/gorilla/mux"
)

// GetTask: 指定されたタスクIDのタスクを取得するハンドラ関数
func GetUserTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskID := vars["task_id"]
		if taskID == "" {
			log.Printf("タスクIDが指定されていません")
			http.Error(w, "タスクIDが指定されていません", http.StatusBadRequest)
			return
		}

		var task models.Task
		err := db.QueryRow("SELECT title, deadline, priority, status, purpose, steps, memo, remarks FROM tasks WHERE id = ?", taskID).Scan(
			&task.Title, &task.Deadline, &task.Priority, &task.Status, &task.Purpose, &task.Steps, &task.Memo, &task.Remarks)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("タスクが見つかりません: %s", taskID)
				http.Error(w, "タスクが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("タスクの取得に失敗しました: %v", err)
				http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(task); err != nil {
			log.Printf("タスクのエンコードに失敗しました: %v", err)
			http.Error(w, "タスクのエンコードに失敗しました", http.StatusInternalServerError)
		}
	}
}
