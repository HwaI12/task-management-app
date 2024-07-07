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

// jwtKey はJWTの署名に使用するシークレットキーです。
var jwtKey = []byte("my_secret_key")

// Claims はJWTトークンのクレームを表します。
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Login はユーザーのログインを処理するハンドラ関数です。
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("リクエストのパースに失敗しました: %v", err)
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		log.Printf("ログインを試行中: ユーザー名: %s", user.Email)

		var storedPassword string
		err = db.QueryRow("SELECT password_hash FROM users WHERE email = ?", user.Email).Scan(&storedPassword)
		if err != nil {
			log.Printf("パスワードハッシュの取得に失敗しました: %v", err)
			http.Error(w, "メールアドレスまたはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		log.Printf("データベースから取得したパスワードハッシュ: %s", storedPassword)

		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.PasswordHash))
		if err != nil {
			log.Printf("パスワードの比較に失敗しました: %v", err)
			log.Printf("データベースのパスワードハッシュ: %s", storedPassword)
			log.Printf("提供されたパスワード: %s", user.PasswordHash)
			http.Error(w, "メールアドレスまたはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		log.Printf("ユーザーが認証されました: %s", user.Email)

		// JWTトークンの有効期限を設定
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// JWTトークンの生成
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			log.Printf("トークンの生成に失敗しました: %v", err)
			http.Error(w, "トークンの生成に失敗しました", http.StatusInternalServerError)
			return
		}

		// クッキーにトークンをセット
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

		log.Printf("ユーザーに対してトークンを生成してセットしました: %s", user.Email)

		// 成功レスポンスを返す
		w.WriteHeader(http.StatusOK)
	}
}
