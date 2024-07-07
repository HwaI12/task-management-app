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
				log.Println("Successfully connected to the database.")
				break
			}
		}
		log.Printf("Failed to connect to database. Retrying in 5 seconds... (%d/5)\n", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db)).Methods("POST")
	router.HandleFunc("/delete", handlers.DeleteUser(db)).Methods("POST")

	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
