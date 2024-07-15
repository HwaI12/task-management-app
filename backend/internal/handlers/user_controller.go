package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/models"
)

// GetUserHandler はユーザー情報を取得するハンドラ関数です
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			log.Printf("ユーザーIDが指定されていません")
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		user, err := models.FetchUser(db, userId)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザーが見つかりません: %s", userId)
				http.Error(w, "ユーザーが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		log.Printf("ユーザー %s の情報を取得しました", userId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
