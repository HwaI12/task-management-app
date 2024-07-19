package service

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/auth"
	"github.com/HwaI12/task-management-app/backend/internal/cookie"
	"github.com/HwaI12/task-management-app/backend/internal/model"
	"github.com/HwaI12/task-management-app/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(req model.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("パスワードのハッシュ化に失敗しました")
	}

	exists, err := s.userRepo.CheckUserExists(req.UserID, req.Email)
	if err != nil {
		return errors.New("ユーザーIDまたはメールアドレスの存在確認中にエラーが発生しました")
	}
	if exists {
		return errors.New("このユーザーIDまたはメールアドレスは既に使用されています")
	}

	return s.userRepo.RegisterUser(req, string(hashedPassword))
}

func (s *UserService) LoginUser(user model.User) (string, error) {
	storedPassword, err := s.userRepo.GetUserPasswordHash(user.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("ユーザが見つかりません")
		}
		return "", errors.New("サーバー内部エラー")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		return "", errors.New("ユーザ名またはパスワードが正しくありません")
	}

	return auth.GenerateJWT(user.UserID)
}

func (s *UserService) DeleteUser(user model.User) error {
	storedPassword, err := s.userRepo.GetUserPasswordHash(user.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("ユーザが見つかりません")
		}
		return errors.New("サーバー内部エラー")
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		return errors.New("ユーザ名またはパスワードが正しくありません")
	}

	rowsAffected, err := s.userRepo.DeleteUser(user.UserID)
	if err != nil {
		return errors.New("ユーザーの削除に失敗しました")
	}

	if rowsAffected == 0 {
		return errors.New("指定されたユーザーが見つかりません")
	}

	return nil
}

func (s *UserService) GetUser(userId string) (model.User, error) {
	return s.userRepo.FetchUser(userId)
}

func (s *UserService) SetTokenCookie(w http.ResponseWriter, tokenString string) {
	cookie.SetTokenCookie(w, tokenString)
}
