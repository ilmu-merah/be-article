package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/ilmu-merah/be-article/internal/app"
	"github.com/ilmu-merah/be-article/internal/middleware"
)

func main() {
	mux := app.NewRouter()


	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "success",
		})
	})


	var handler http.Handler = mux
	handler = middleware.ErrRecoveryMiddleware(handler)
	

	port := app.GetEnv("APP_PORT", ":8080")
	fmt.Printf("[SERVER] berjalan di: %s%s \n", app.GetEnv("APP_HOST", "localhost"), port)
	log.Fatal(http.ListenAndServe(port, handler))
}