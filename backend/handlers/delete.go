package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM users WHERE email = ?", user.Email)
		if err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
