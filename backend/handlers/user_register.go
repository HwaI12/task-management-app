// handlers/user_register.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest は登録リクエストのJSON構造体です

// Register は新規ユーザー登録を処理するハンドラ関数です
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
		var exists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = ? OR email = ?)", req.UserID, req.Email).Scan(&exists)
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
		query := "INSERT INTO users (user_id, username, email, password_hash) VALUES (?, ?, ?, ?)"
		_, err = db.Exec(query, req.UserID, req.Username, req.Email, hashedPassword)
		if err != nil {
			log.Printf("ユーザーの保存に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		log.Printf("新しいユーザー %s が登録されました", req.UserID)
		w.WriteHeader(http.StatusCreated)
	}
}
