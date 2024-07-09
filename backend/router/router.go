package router

import (
	"database/sql"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// SetupRouter は指定されたデータベースを使用してHTTPリクエストを処理するルーターを設定します。
func SetupRouter(db *sql.DB) http.Handler {
	// Gorilla Muxを使用して新しいルーターを作成
	router := mux.NewRouter()

	// CORSミドルウェアの設定
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                   // 許可するオリジン
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 許可するHTTPメソッド
		AllowedHeaders:   []string{"Content-Type", "Authorization"},           // 許可するヘッダー
		AllowCredentials: true,                                                // クレデンシャル（Cookie等）を許可するかどうか
	})

	// ハンドラ関数の登録
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")    // POST /register に対するハンドラ
	router.HandleFunc("/login", handlers.Login(db)).Methods("POST")          // POST /login に対するハンドラ
	router.HandleFunc("/delete", handlers.DeleteUser(db)).Methods("POST")    // POST /delete に対するハンドラ
	router.HandleFunc("/api/tasks", handlers.CreateTask(db)).Methods("POST") // POST /api/tasks に対するハンドラ
	router.HandleFunc("/api/tasks", handlers.GetTasks(db)).Methods("GET")    // GET /api/tasks に対するハンドラ
	router.HandleFunc("/api/user", handlers.GetUser(db)).Methods("GET")

	// CORSミドルウェアをルーターに適用
	handler := corsHandler.Handler(router)

	return handler
}
