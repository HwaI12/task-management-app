package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	// Import your user model package
	"github.com/HwaI12/task-management-app/backend/models"
)

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		result, err := db.Exec("DELETE FROM users WHERE email = ?", user.Email)
		if err != nil {
			log.Printf("Error deleting user: %v", err)
			http.Error(w, "ユーザーの削除に失敗しました", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("Error getting rows affected: %v", err)
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
