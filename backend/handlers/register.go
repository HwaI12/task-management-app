package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing password:", err)
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)", user.Username, user.Email, string(hashedPassword))
		if err != nil {
			log.Println("Error saving user:", err)
			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
