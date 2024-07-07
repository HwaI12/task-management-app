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
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Failed to parse login request: %v", err)
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		log.Printf("Attempting login for username: %s", user.Username)

		var storedPassword string
		err = db.QueryRow("SELECT password_hash FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
		if err != nil {
			log.Printf("Failed to retrieve password hash: %v", err)
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		log.Printf("Retrieved password hash from database: %s", storedPassword)

		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.PasswordHash))
		if err != nil {
			log.Printf("Password comparison failed: %v", err)
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		log.Printf("User authenticated: %s", user.Username)

		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			log.Printf("Failed to generate token: %v", err)
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Set cookie with JWT token
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
			Path:    "/",
		})

		log.Printf("Generated and set token for user: %s", user.Username)

		w.WriteHeader(http.StatusOK)
	}
}
