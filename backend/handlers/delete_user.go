// handlers/delete_user.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
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

		result, err := db.Exec("DELETE FROM users WHERE email = ?", user.Email)
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
			http.Error(w, "指定されたメールアドレスのユーザーが見つかりません", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ユーザーが正常に削除されました"})
	}
}
