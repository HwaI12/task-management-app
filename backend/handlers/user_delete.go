// handlers/user_delete.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
	"golang.org/x/crypto/bcrypt"
)

// DeleteUser はユーザーの削除を処理するハンドラ関数です
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("リクエストの解析に失敗しました: %v", err)
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		// ユーザーのパスワードハッシュをデータベースから取得
		var storedPassword string
		err = db.QueryRow("SELECT password_hash FROM users WHERE user_id = ?", user.User_id).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザが見つかりません: %s", user.User_id)
				http.Error(w, "ユーザが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		// パスワードの比較
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
		if err != nil {
			log.Printf("パスワードの比較に失敗しました: %v", err)
			http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		// ユーザーの削除
		result, err := db.Exec("DELETE FROM users WHERE user_id = ?", user.User_id)
		if err != nil {
			log.Printf("ユーザーの削除に失敗しました: %v", err)
			http.Error(w, "ユーザーの削除に失敗しました", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("行数の取得に失敗しました: %v", err)
			http.Error(w, "ユーザーの削除結果の確認に失敗しました", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "指定されたユーザーが見つかりません", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ユーザーが正常に削除されました"})
	}
}
