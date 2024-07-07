package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/HwaI12/task-management-app/backend/models"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		var storedPassword string
		err := db.QueryRow("SELECT password_hash FROM users WHERE email = ?", user.Email).Scan(&storedPassword)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.PasswordHash))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
