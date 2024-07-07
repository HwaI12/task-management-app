package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/task-management-app/backend/models"
)

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		_, err := db.Exec("DELETE FROM users WHERE email = ?", user.Email)
		if err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
