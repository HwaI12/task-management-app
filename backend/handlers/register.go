// handlers/register.go

package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// RegisterRequest は登録リクエストのJSON構造体です
type RegisterRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// Register は新規ユーザー登録を処理するハンドラ関数です
func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエストボディをパース
		var req RegisterRequest
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

		// メールアドレスが既に存在するか確認
		var emailExists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&emailExists)
		if err != nil {
			log.Printf("メールアドレスの存在確認中にエラーが発生しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}
		if emailExists {
			log.Printf("メールアドレス %s はすでに使用されています", req.Email)
			http.Error(w, "このメールアドレスは既に使用されています", http.StatusConflict)
			return
		}

		// ユーザーをデータベースに登録
		query := "INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)"
		_, err = db.Exec(query, req.Username, req.Email, hashedPassword)
		if err != nil {
			log.Printf("ユーザーの保存に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		log.Println("ユーザーが正常に登録されました.")
		w.WriteHeader(http.StatusCreated)
	}
}
