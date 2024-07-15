package main

import (
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/db"
	"github.com/HwaI12/task-management-app/backend/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("データベース接続を開始します")
	dbConn := db.InitDB()

	log.Println("ルーターを設定します")
	router := router.SetupRouter(dbConn)

	log.Println("サーバーを起動します")
	startServer(router)

}

func startServer(router http.Handler) {
	log.Println("サーバーをポート :8000 で起動しています")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
