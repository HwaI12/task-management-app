package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/HwaI12/task-management-app/backend/models"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Invalid request payload: %v", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		log.Printf("User attempting login: %s", user.Email)

		var storedPassword string
		err = db.QueryRow("SELECT password_hash FROM users WHERE email = ?", user.Email).Scan(&storedPassword)
		if err != nil {
			log.Printf("Error retrieving password hash: %v", err)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		log.Printf("Stored password hash: %s", storedPassword)

		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.PasswordHash))
		if err != nil {
			log.Printf("Password comparison failed: %v", err)
			log.Printf("Stored password: %s", storedPassword)
			log.Printf("Provided password: %s", user.PasswordHash)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		log.Printf("User authenticated: %s", user.Email)

		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			log.Printf("Error generating token: %v", err)
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		log.Printf("Token generated and set for user: %s", user.Email)

		w.WriteHeader(http.StatusOK)
	}
}
