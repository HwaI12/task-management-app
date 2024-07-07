module github.com/HwaI12/task-management-app/backend

go 1.22

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.8.1
	github.com/gorilla/mux v1.8.1
	github.com/rs/cors v1.11.0
	golang.org/x/crypto v0.25.0
)

// Indirect dependencies
require filippo.io/edwards25519 v1.1.0 // indirect
