// backend/internal/router/router.go

package router

import (
	"database/sql"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/internal/controller"
	"github.com/HwaI12/task-management-app/backend/internal/middleware" // 新しいCORSミドルウェアのインポート
	"github.com/gorilla/mux"
)

// SetupRouter は指定されたデータベースを使用してHTTPリクエストを処理するルーターを設定します。
func SetupRouter(db *sql.DB) http.Handler {
	// Gorilla Muxを使用して新しいルーターを作成
	router := mux.NewRouter()

	// ハンドラ関数の登録
	router.HandleFunc("/register", controller.Register(db)).Methods("POST")    // POST /register に対するハンドラ
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")          // POST /login に対するハンドラ
	router.HandleFunc("/delete", controller.DeleteUser(db)).Methods("POST")    // POST /delete に対するハンドラ
	router.HandleFunc("/api/tasks", controller.CreateTask(db)).Methods("POST") // POST /api/tasks に対するハンドラ
	router.HandleFunc("/api/tasks", controller.GetTasks(db)).Methods("GET")    // GET /api/tasks に対するハンドラ
	router.HandleFunc("/api/tasks/{task_id}", controller.GetUserTasks(db)).Methods("GET")
	router.HandleFunc("/api/user", controller.GetUser(db)).Methods("GET") // GET /api/user に対するハンドラ

	// CORSミドルウェアをルーターに適用
	handler := middleware.CORSHandler()(router)

	return handler
}
