// backend/internal/router/router.go

package router

import (
	"database/sql"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/controller"
	"github.com/HwaI12/task-management-app/backend/internal/middleware"
	"github.com/HwaI12/task-management-app/backend/internal/repository"
	"github.com/HwaI12/task-management-app/backend/internal/service"
	"github.com/gorilla/mux"
)

// SetupRouter は指定されたデータベースを使用してHTTPリクエストを処理するルーターを設定します。
func SetupRouter(db *sql.DB) http.Handler {
	// Gorilla Muxを使用して新しいルーターを作成
	router := mux.NewRouter()

	// リポジトリの作成
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	// サービスの作成
	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	// コントローラーの作成
	taskController := controller.NewTaskController(taskService)
	userController := controller.NewUserController(userService)

	// ハンドラ関数の登録
	router.HandleFunc("/register", userController.Register).Methods("POST")
	router.HandleFunc("/login", userController.Login).Methods("POST")
	router.HandleFunc("/delete", userController.DeleteUser).Methods("POST")
	router.HandleFunc("/api/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{task_id}", taskController.GetUserTasks).Methods("GET")
	router.HandleFunc("/api/user", userController.GetUser).Methods("GET")

	// CORSミドルウェアをルーターに適用
	handler := middleware.CORSHandler()(router)

	return handler
}
