package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)



func ErrRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC ERROR] message: %s", err)
				
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
	
				json.NewEncoder(w).Encode(map[string]string{
					"error":"terjadi kesalahan di server",
				})
			}

		}()

		next.ServeHTTP(w, r)
	})
}