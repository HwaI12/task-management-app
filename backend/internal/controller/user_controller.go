package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/auth"
	"github.com/HwaI12/task-management-app/backend/internal/cookie"
	"github.com/HwaI12/task-management-app/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// 新規ユーザー登録を処理するハンドラ関数
func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエストボディをパース
		var req models.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("リクエストのパースに失敗しました: %v", err)
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		// パスワードをハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		// ユーザーIDまたはメールアドレスが既に存在するか確認
		exists, err := models.CheckUserExists(db, req.UserID, req.Email)
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
		err = models.RegisterUser(db, req, string(hashedPassword))
		if err != nil {
			log.Printf("ユーザーの保存に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		log.Printf("新しいユーザー %s が登録されました", req.UserID)
		w.WriteHeader(http.StatusCreated)
	}
}

// ユーザーのログインを処理するハンドラ関数
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
		storedPassword, err := models.GetUserPasswordHash(db, user.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザ %s が見つかりません", user.UserID)
				http.Error(w, "ユーザが見つかりません", http.StatusNotFound)
				return
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
				return
			}
		}

		// パスワードの比較
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
		if err != nil {
			log.Printf("ユーザ %s のパスワード比較に失敗しました: %v", user.UserID, err)
			http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		// JWTトークンを生成
		tokenString, err := auth.GenerateJWT(user.UserID)
		if err != nil {
			log.Printf("トークン生成に失敗しました: %v", err)
			http.Error(w, "トークンの生成に失敗しました", http.StatusInternalServerError)
			return
		}

		// クッキーにトークンを設定
		cookie.SetTokenCookie(w, tokenString)

		// レスポンスとして成功メッセージを返す
		w.Header().Set("Content-Type", "application/json")
		log.Printf("ユーザ %s が正常にログインしました", user.UserID)
		json.NewEncoder(w).Encode(map[string]string{"message": "ログインに成功しました"})
	}
}

// ユーザーの削除を処理するハンドラ関数
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
		storedPassword, err := models.GetUserPasswordHash(db, user.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザ %s が見つかりません", user.UserID)
				http.Error(w, "ユーザが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		// パスワードの比較
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
		if err != nil {
			log.Printf("パスワードの比較に失敗しました: %v", err)
			http.Error(w, "ユーザ名またはパスワードが正しくありません", http.StatusUnauthorized)
			return
		}

		result, err := models.DeleteUser(db, user.UserID)
		if err != nil {
			log.Printf("ユーザーの削除に失敗しました: %v", err)
			http.Error(w, "ユーザーの削除に失敗しました", http.StatusInternalServerError)
			return
		}

		// 削除された行数を取得
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("行数の取得に失敗しました: %v", err)
			http.Error(w, "ユーザーの削除結果の確認に失敗しました", http.StatusInternalServerError)
			return
		}

		// 削除された行数が0の場合はユーザーが見つからなかったとしてエラーを返す
		if rowsAffected == 0 {
			http.Error(w, "指定されたユーザーが見つかりません", http.StatusNotFound)
			return
		}

		log.Printf("ユーザー %s が正常に削除されました", user.UserID)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "ユーザーが正常に削除されました"})
	}
}

// ユーザー情報を取得するハンドラ関数
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			log.Printf("ユーザーIDが指定されていません")
			http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
			return
		}

		user, err := models.FetchUser(db, userId)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("ユーザーが見つかりません: %s", userId)
				http.Error(w, "ユーザーが見つかりません", http.StatusNotFound)
			} else {
				log.Printf("データベースエラー: %v", err)
				http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
			}
			return
		}

		log.Printf("ユーザー %s の情報を取得しました", userId)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
