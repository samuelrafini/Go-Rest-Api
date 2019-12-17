package middleware

import (
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware (h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, e := w.Write([]byte("You are not authorized")); if e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("An error accured: %v", e)
			}
			return
		}
	})
}