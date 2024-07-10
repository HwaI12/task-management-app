// handlers/user_get.go

package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

// GetUser はユーザー情報を取得するハンドラ関数です
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		user, err := models.GetUserByID(db, userId)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "ユーザーが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
