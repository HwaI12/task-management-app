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

// Claims はJWTのクレームを表す構造体です
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login はユーザーのログインを処理するハンドラ関数です
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("ログインリクエストの解析に失敗しました: %v", err)
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		log.Printf("ユーザ名 %s でのログイン試行", user.Username)

		var storedPassword string
		err = db.QueryRow("SELECT password_hash FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザが見つかりません: %s", user.Username)
				http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
		if err != nil {
			log.Printf("ユーザ %s のパスワード比較に失敗しました: %v", user.Username, err)
			http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		// トークンの有効期限を設定
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		// JWTトークンを生成
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			log.Printf("トークン生成に失敗しました: %v", err)
			http.Error(w, "トークンの生成に失敗しました", http.StatusInternalServerError)
			return
		}

		// クッキーにトークンを設定
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, // 開発環境ではfalse、本番環境ではtrueに設定
			SameSite: http.SameSiteLaxMode,
		})

		// レスポンスとして成功メッセージを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "ログインに成功しました"})
	}
}
