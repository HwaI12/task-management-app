// handlers/user_register.go

package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/models"
)

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
		err = models.CreateUser(db, req.UserID, req.Username, req.Email, req.PasswordHash)
		if err != nil {
			log.Printf("ユーザーの保存に失敗しました: %v", err)
			http.Error(w, "ユーザーの登録に失敗しました", http.StatusInternalServerError)
			return
		}

		log.Println("ユーザーが正常に登録されました.")
		w.WriteHeader(http.StatusCreated)
	}
}
