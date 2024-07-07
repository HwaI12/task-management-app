package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/HwaI12/task-management-app/backend/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp(db:3306)/"+os.Getenv("MYSQL_DATABASE"))
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
