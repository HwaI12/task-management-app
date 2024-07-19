package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/model"
	"github.com/HwaI12/task-management-app/backend/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("リクエストのパースに失敗しました: %v", err)
		http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
		return
	}

	if err := c.userService.RegisterUser(req); err != nil {
		log.Printf("ユーザーの登録に失敗しました: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("新しいユーザー %s が登録されました", req.UserID)
	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("ログインリクエストの解析に失敗しました: %v", err)
		http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
		return
	}

	tokenString, err := c.userService.LoginUser(user)
	if err != nil {
		log.Printf("ログインに失敗しました: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// クッキーにトークンを設定
	c.userService.SetTokenCookie(w, tokenString)

	w.Header().Set("Content-Type", "application/json")
	log.Printf("ユーザ %s が正常にログインしました", user.UserID)
	json.NewEncoder(w).Encode(map[string]string{"message": "ログインに成功しました"})
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("リクエストの解析に失敗しました: %v", err)
		http.Error(w, "リクエストの形式が正しくありません", http.StatusBadRequest)
		return
	}

	if err := c.userService.DeleteUser(user); err != nil {
		log.Printf("ユーザーの削除に失敗しました: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("ユーザー %s が正常に削除されました", user.UserID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ユーザーが正常に削除されました"})
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		log.Printf("ユーザーIDが指定されていません")
		http.Error(w, "ユーザーIDが指定されていません", http.StatusBadRequest)
		return
	}

	user, err := c.userService.GetUser(userId)
	if err != nil {
		log.Printf("ユーザー情報の取得に失敗しました: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("ユーザー %s の情報を取得しました", userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
