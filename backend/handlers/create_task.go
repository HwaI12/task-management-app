package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

// CreateTask は新しいタスクを作成するハンドラ関数です
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
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
		_, err = db.Exec(query, task.UserID, task.Title, task.Deadline, task.Priority, task.Status, task.Purpose, task.Steps, task.Memo, task.Remarks)
		if err != nil {
			log.Printf("タスクの保存に失敗しました: %v", err)
			http.Error(w, "タスクの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		log.Println("タスクが正常に追加されました")
		w.WriteHeader(http.StatusCreated)
	}
}
