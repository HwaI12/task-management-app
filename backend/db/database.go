package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB は環境変数から取得した情報を使ってMySQLデータベースに接続します。
// 接続に失敗した場合、最大5回のリトライを行います。
func InitDB() *sql.DB {
	// 環境変数からMySQL接続情報を取得
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	// 必須の環境変数が設定されているか確認
	if user == "" || password == "" || host == "" || database == "" {
		log.Fatal("環境変数が設定されていません")
	}

	// 接続文字列（DSN）を構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	log.Println("DSN:", dsn)

	// 最大5回の接続リトライを試行
	var err error
	maxAttempts := 5
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("データベースへの接続に失敗しました。リトライします... (%d/%d)\n", attempt, maxAttempts)
			time.Sleep(5 * time.Second)
			continue
		}

		err = DB.Ping()
		if err == nil {
			log.Println("データベースに正常に接続しました")
			break
		}

		log.Printf("データベースへの接続に失敗しました。リトライします... (%d/%d)\n", attempt, maxAttempts)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}

	return DB
}
