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

// Register は新規ユーザー登録を処理します
func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("リクエストのパースに失敗しました: %v", err)
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		// パスワードをハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		// ユーザーIDまたはメールアドレスが既に存在するか確認
		exists, err := models.UserExists(db, req.UserID, req.Email)
		if err != nil {
			log.Printf("ユーザーIDまたはメールアドレスの存在確認中にエラーが発生しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}
		if exists {
			log.Printf("ユーザーID %s またはメールアドレス %s はすでに使用されています", req.UserID, req.Email)
			http.Error(w, "このユーザーIDまたはメールアドレスは既に使用されています", http.StatusConflict)
			return
		}

		// ユーザーをデータベースに登録
		user := models.User{
			UserID:       req.UserID,
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
		}
		err = models.CreateUser(db, user)
		if err != nil {
			log.Printf("ユーザーの保存に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		log.Println("ユーザーが正常に登録されました.")
		w.WriteHeader(http.StatusCreated)
	}
}

// Login はユーザーのログインを処理します
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
		storedUser, err := models.GetUserByID(db, user.UserID)
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
		err = bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(user.Password))
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

// DeleteUser はユーザーの削除を処理します
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("リクエストの解析に失敗しました: %v", err)
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		// ユーザーのパスワードハッシュをデータベースから取得
		storedUser, err := models.GetUserByID(db, user.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザが見つかりません: %s", user.UserID)
				http.Error(w, "ユーザが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		// パスワードの比較
		err = bcrypt.CompareHashAndPassword([]byte(storedUser.PasswordHash), []byte(user.Password))
		if err != nil {
			log.Printf("パスワードの比較に失敗しました: %v", err)
			http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		// ユーザーの削除
		err = models.DeleteUser(db, user.UserID)
		if err != nil {
			log.Printf("ユーザーの削除に失敗しました: %v", err)
			http.Error(w, "ユーザーの削除に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ユーザーが正常に削除されました"})
	}
}

// GetUser はユーザー情報を取得するハンドラ関数です
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		user, err := models.GetUserByID(db, userID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "ユーザーが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
