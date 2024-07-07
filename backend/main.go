package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/HwaI12/task-management-app/backend/handlers"
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

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db)).Methods("POST")
	router.HandleFunc("/delete", handlers.DeleteUser(db)).Methods("POST")

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
