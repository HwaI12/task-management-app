package models

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id"`
	User_id      string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Password     string `json:"password" gorm:"-"`
	CreatedAt    string `json:"created_at"`
}

type RegisterRequest struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// UserExists はユーザーIDまたはメールアドレスが既に存在するかを確認します
func UserExists(db *sql.DB, userID, email string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE user_id = ? OR email = ?)", userID, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// CreateUser は新しいユーザーをデータベースに登録します
func CreateUser(db *sql.DB, userID, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("パスワードのハッシュ化に失敗しました: %v", err)
		return err
	}

	query := "INSERT INTO users (user_id, username, email, password_hash) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, userID, username, email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

// GetUserPassword はユーザーのパスワードハッシュをデータベースから取得します
func GetUserPassword(db *sql.DB, userID string) (string, error) {
	var storedPassword string
	err := db.QueryRow("SELECT password_hash FROM users WHERE user_id = ?", userID).Scan(&storedPassword)
	if err != nil {
		return "", err
	}
	return storedPassword, nil
}
