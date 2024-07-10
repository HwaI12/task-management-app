package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID           int    `json:"id"`
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Password     string `json:"password,omitempty"`
	CreatedAt    string `json:"created_at"`
}

type RegisterRequest struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser は新しいユーザーをデータベースに追加します
func CreateUser(db *sql.DB, user User) error {
	query := "INSERT INTO users (user_id, username, email, password_hash) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, user.UserID, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("ユーザーの保存に失敗しました: %v", err)
	}
	return nil
}

// GetUserByID は指定されたユーザーIDのユーザー情報をデータベースから取得します
func GetUserByID(db *sql.DB, userID string) (User, error) {
	var user User
	query := "SELECT id, user_id, username, email, password_hash, created_at FROM users WHERE user_id = ?"
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.UserID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, fmt.Errorf("ユーザーの取得に失敗しました: %v", err)
	}
	return user, nil
}

// UserExists は指定されたユーザーIDまたはメールアドレスが既に存在するかを確認します
func UserExists(db *sql.DB, userID, email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE user_id = ? OR email = ?)"
	err := db.QueryRow(query, userID, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ユーザーIDまたはメールアドレスの存在確認中にエラーが発生しました: %v", err)
	}
	return exists, nil
}

// DeleteUser は指定されたユーザーIDのユーザーをデータベースから削除します
func DeleteUser(db *sql.DB, userID string) error {
	query := "DELETE FROM users WHERE user_id = ?"
	result, err := db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("ユーザーの削除に失敗しました: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("行数の取得に失敗しました: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("指定されたユーザーが見つかりません")
	}

	return nil
}
