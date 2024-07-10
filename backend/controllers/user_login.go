// handlers/user_login.go

package controllers

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
	UserID string `json:"user_id"`
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

		log.Printf("ユーザ名 %s でのログイン試行", user.UserID)

		// ユーザーのパスワードハッシュをデータベースから取得
		storedPassword, err := models.GetUserPassword(db, user.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザが見つかりません: %s", user.UserID)
				http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		// パスワードの比較
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
		if err != nil {
			log.Printf("ユーザ %s のパスワード比較に失敗しました: %v", user.UserID, err)
			http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		// トークンの有効期限を設定
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			UserID: user.UserID,
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
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		})

		// レスポンスとして成功メッセージを返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "ログインに成功しました"})
	}
}
