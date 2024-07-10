package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
	"github.com/gorilla/mux"
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
		taskID, err := models.CreateTask(db, task)
		if err != nil {
			log.Printf("タスクの保存に失敗しました: %v", err)
			http.Error(w, "タスクの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		task.ID = int(taskID)

		log.Println("タスクが正常に追加されました")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}

// GetTasks は指定されたユーザーのタスクを取得するハンドラ関数です
func GetTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		tasks, err := models.GetTasksByUserID(db, userID)
		if err != nil {
			log.Printf("タスクの取得に失敗しました: %v", err)
			http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			log.Printf("タスクのエンコードに失敗しました: %v", err)
			http.Error(w, "タスクの取得に失敗しました", http.StatusInternalServerError)
		}
	}
}

// GetUserTasks は指定されたタスクIDのタスクを取得するハンドラ関数です
func GetUserTasks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskID := vars["task_id"]
		if taskID == "" {
			http.Error(w, "タスクIDが指定されていません", http.StatusBadRequest)
			return
		}

		task, err := models.GetTaskByID(db, taskID)
		if err != nil {
			if err == sql.ErrNoRows {
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
