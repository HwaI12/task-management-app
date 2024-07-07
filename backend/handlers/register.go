package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// RegisterRequestは新規ユーザー登録のリクエストを表す
type RegisterRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// Registerは新規ユーザーを登録する
func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// リクエストデータの読み込みとパース
		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("リクエストのパースに失敗しました: %v", err)
			http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
			return
		}

		log.Printf("ユーザーを登録中: ユーザー名: %s, メールアドレス: %s", req.Username, req.Email)

		// パスワードのハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		// ユーザーをデータベースに登録する前に、同じメールアドレスが存在しないか確認
		var emailExists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", req.Email).Scan(&emailExists)
		if err != nil {
			log.Printf("ユーザーの存在確認中にエラーが発生しました: %v", err)
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
