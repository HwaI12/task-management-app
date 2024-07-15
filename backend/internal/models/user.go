package models

import "database/sql"

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

// データベースにユーザーを保存する
func FetchUser(db *sql.DB, userId string) (User, error) {
	query := "SELECT user_id, username FROM users WHERE user_id = ?"
	var user User
	err := db.QueryRow(query, userId).Scan(&user.User_id, &user.Username)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
