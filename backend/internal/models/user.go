package models

import "database/sql"

type User struct {
	ID           int    `json:"id"`
	UserID       string `json:"user_id"`
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

// FetchUser はデータベースからユーザーを取得する関数です
func FetchUser(db *sql.DB, userId string) (User, error) {
	query := "SELECT user_id, username FROM users WHERE user_id = ?"
	var user User
	err := db.QueryRow(query, userId).Scan(&user.UserID, &user.Username)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// CheckUserExists はユーザーIDまたはメールアドレスが既に存在するかを確認する関数です
func CheckUserExists(db *sql.DB, userId string, email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE user_id = ? OR email = ?)"
	var exists bool
	err := db.QueryRow(query, userId, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// RegisterUser はデータベースにユーザー情報を新規登録する関数です
func RegisterUser(db *sql.DB, req RegisterRequest, hashedPassword string) error {
	query := "INSERT INTO users (user_id, username, email, password_hash) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, req.UserID, req.Username, req.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

// データベースからユーザのパスワードハッシュを取得する関数
func GetUserPasswordHash(db *sql.DB, userId string) (string, error) {
	query := "SELECT password_hash FROM users WHERE user_id = ?"
	var storedPassword string
	err := db.QueryRow(query, userId).Scan(&storedPassword)
	if err != nil {
		return "", err
	}
	return storedPassword, nil
}

func DeleteUser(db *sql.DB, userId string) (sql.Result, error) {
	query := "DELETE FROM users WHERE user_id = ?"
	result, err := db.Exec(query, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
