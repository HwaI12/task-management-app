package repository

import (
	"database/sql"
	"fmt"

	"github.com/HwaI12/task-management-app/backend/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CheckUserExists(userID, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_id = ? OR email = ?)`
	var exists bool
	err := r.db.QueryRow(query, userID, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ユーザー存在確認中にエラーが発生しました: %w", err)
	}
	return exists, nil
}

func (r *UserRepository) RegisterUser(req model.RegisterRequest, hashedPassword string) error {
	query := `INSERT INTO users (user_id, username, email, password_hash) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, req.UserID, req.Username, req.Email, hashedPassword)
	if err != nil {
		return fmt.Errorf("ユーザー登録中にエラーが発生しました: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserPasswordHash(userID string) (string, error) {
	query := `SELECT password_hash FROM users WHERE user_id = ?`
	var storedPassword string
	err := r.db.QueryRow(query, userID).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("ユーザーが見つかりません: %s", userID)
		}
		return "", fmt.Errorf("パスワードハッシュ取得中にエラーが発生しました: %w", err)
	}
	return storedPassword, nil
}

func (r *UserRepository) DeleteUser(userID string) (int64, error) {
	query := `DELETE FROM users WHERE user_id = ?`
	result, err := r.db.Exec(query, userID)
	if err != nil {
		return 0, fmt.Errorf("ユーザー削除中にエラーが発生しました: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("削除された行数の取得中にエラーが発生しました: %w", err)
	}
	return rowsAffected, nil
}

func (r *UserRepository) FetchUser(userID string) (model.User, error) {
	query := `SELECT user_id, username FROM users WHERE user_id = ?`
	var user model.User
	err := r.db.QueryRow(query, userID).Scan(&user.UserID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("ユーザーが見つかりません: %s", userID)
		}
		return model.User{}, fmt.Errorf("ユーザー情報取得中にエラーが発生しました: %w", err)
	}
	return user, nil
}
