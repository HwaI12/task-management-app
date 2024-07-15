package cookie

import (
	"net/http"
	"time"
)

// JWTトークンをクッキーに設定する
func SetTokenCookie(w http.ResponseWriter, tokenString string) {
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}
