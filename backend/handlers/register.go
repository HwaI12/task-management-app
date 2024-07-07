package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/HwaI12/task-management-app/backend/models"
)

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", user.Username, user.Email, string(hashedPassword))
		if err != nil {
			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
