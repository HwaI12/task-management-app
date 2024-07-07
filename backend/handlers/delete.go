package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

// DeleteUser はログインしているユーザーが自身のアカウントを削除するためのハンドラ関数です。
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエストボディからユーザー情報をデコード
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		// データベースからユーザーを削除
		_, err = db.Exec("DELETE FROM users WHERE email = ?", user.Email)
		if err != nil {
			http.Error(w, "ユーザーの削除に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
