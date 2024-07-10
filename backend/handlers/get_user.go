package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// User はユーザー情報を表す構造体です
type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

// GetUser: ユーザー情報を取得するハンドラ関数
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			log.Printf("ユーザーIDが指定されていません")
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		var user User
		err := db.QueryRow("SELECT user_id, username FROM users WHERE user_id = ?", userId).Scan(&user.UserID, &user.Username)
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
