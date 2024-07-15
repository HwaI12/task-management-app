// backend/internal/middleware/cors.go

package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// CORSHandler はCORSミドルウェアを返します。
func CORSHandler() func(http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                   // 許可するオリジン
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 許可するHTTPメソッド
		AllowedHeaders:   []string{"Content-Type", "Authorization"},           // 許可するヘッダー
		AllowCredentials: true,                                                // クレデンシャル（Cookie等）を許可するかどうか
	})
	return corsHandler.Handler
}
