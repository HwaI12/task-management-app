package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

// CreateTask: 新しいタスクを作成するハンドラ関数
func CreateTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			log.Printf("リクエストの解析に失敗しました: %v", err)
			http.Error(w, "リクエスト形式が正しくありません", http.StatusBadRequest)
			return
		}

		// データベースにタスクを挿入
		query := `
        INSERT INTO tasks (user_id, title, deadline, priority, status, purpose, steps, memo, remarks)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
        `
		result, err := db.Exec(query, task.UserID, task.Title, task.Deadline, task.Priority, task.Status, task.Purpose, task.Steps, task.Memo, task.Remarks)
		if err != nil {
			log.Printf("タスクの保存に失敗しました: %v", err)
			http.Error(w, "タスクの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		// 挿入されたタスクのIDを取得
		taskID, err := result.LastInsertId()
		if err != nil {
			log.Printf("挿入されたタスクのID取得に失敗しました: %v", err)
			http.Error(w, "タスクの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		task.ID = int(taskID)

		log.Printf("新しいタスクが作成されました: %v", task)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}
