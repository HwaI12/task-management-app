// handlers/login.go

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
			log.Printf("リクエストの解析に失敗しました: %v", err)
			http.Error(w, "無効なリクエスト形式", http.StatusBadRequest)
			return
		}

		log.Printf("ユーザー名 %s でのログインを試行中", user.Username)

		var storedPassword string
		err = db.QueryRow("SELECT password_hash FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
		if err != nil {
			log.Printf("パスワードハッシュの取得に失敗しました: %v", err)
			http.Error(w, "無効なユーザー名またはパスワード", http.StatusUnauthorized)
			return
		}

		log.Printf("データベースから取得したパスワードハッシュ: %s", storedPassword)

		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.PasswordHash))
		if err != nil {
			log.Printf("パスワードの比較に失敗しました: %v", err)
			log.Printf("データベースのパスワードハッシュ: %s", storedPassword)
			log.Printf("提供されたパスワード: %s", user.PasswordHash)
			http.Error(w, "無効なユーザー名またはパスワード", http.StatusUnauthorized)
			return
		}

		log.Printf("ユーザーが認証されました: %s", user.Username)

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
			log.Printf("トークンの生成に失敗しました: %v", err)
			http.Error(w, "トークンの生成に失敗しました", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		log.Printf("ユーザーに対してトークンを生成してセットしました: %s", user.Username)

		w.WriteHeader(http.StatusOK)
	}
}
