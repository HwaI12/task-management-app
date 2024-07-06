package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		data := Data{Message: "Hello from Go backend!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe(":8000", r)
}
