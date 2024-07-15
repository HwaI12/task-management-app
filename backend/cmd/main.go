package main

import (
	"context"
	"log"
	"net/http"

	"github.com/HwaI12/task-management-app/backend/db"
	logger "github.com/HwaI12/task-management-app/backend/internal/log"
	"github.com/HwaI12/task-management-app/backend/internal/transaction"
	"github.com/HwaI12/task-management-app/backend/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.InitializeLogger()
	transaction.InitializeGlobalTransaction()

	ctx := context.Background()
	ctx = transaction.InitializeTransaction(ctx)

	entry := logger.WithTransaction(ctx)

	entry.Infof("データベース接続を開始します")
	dbConn := db.InitDB()

	entry.Infof("ルーターを設定します")
	router := router.SetupRouter(dbConn)

	entry.Infof("サーバーを起動します")
	startServer(router)

	defer entry.Infof("サーバーが終了しました")

}

func startServer(router http.Handler) {
	log.Println("サーバーをポート :8000 で起動しています")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
