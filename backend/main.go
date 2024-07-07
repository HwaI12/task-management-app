// main.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HwaI12/task-management-app/backend/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors" // corsミドルウェアのインポート
)

var db *sql.DB

func main() {
	var err error

	// 環境変数を取得
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	if user == "" || password == "" || host == "" || database == "" {
		log.Fatal("環境変数が設定されていません")
	}

	// DSNの作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	log.Println("DSN:", dsn)

	// データベース接続をリトライ
	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("データベースに正常に接続しました")
				break
			}
		}
		log.Printf("データベースへの接続に失敗しました。リトライします... (%d/5)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}

	router := mux.NewRouter()

	// CORS設定
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // フロントエンドのURL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// ハンドラを登録
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db)).Methods("POST")
	router.HandleFunc("/delete", handlers.DeleteUser(db)).Methods("POST")

	// CORSハンドラを適用
	handler := corsHandler.Handler(router)

	log.Println("サーバーをポート :8000 で起動しています")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
