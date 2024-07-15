package models

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
