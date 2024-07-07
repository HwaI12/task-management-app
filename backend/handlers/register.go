package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("Invalid request payload: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Printf("Registering user: %s, email: %s", req.Username, req.Email)

		// パスワードのハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		query := "INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)"
		_, err = db.Exec(query, req.Username, req.Email, hashedPassword)
		if err != nil {
			log.Printf("Error saving user: %v", err)
			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		log.Println("User successfully registered.")
		w.WriteHeader(http.StatusCreated)
	}
}
