package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ilmu-merah/be-article/internal/app"
	"github.com/ilmu-merah/be-article/internal/auth"
	"github.com/ilmu-merah/be-article/internal/config"
	"github.com/ilmu-merah/be-article/internal/middleware"
	"github.com/ilmu-merah/be-article/internal/repository"
)


func main() {
	DB, err := config.InitDB()
	if err != nil {
		log.Fatalf("[Error] message: %s", err)
	}
	defer config.CloseDB()

	DB.AutoMigrate(&repository.User{}, &repository.EmailVerification{})

	// Inisialisasi Repository, Service, dan Handler
	userRepo := repository.NewUserRepository(DB)
	authService := auth.NewAuthService(userRepo)
	authHandler := auth.NewAuthHandler(authService)

	mux := app.NewRouter(authHandler)

	fmt.Printf("database name: %s \n", DB.Migrator().CurrentDatabase())

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message" : "success",
		})
	})


	var handler http.Handler = mux
	handler = middleware.ErrRecoveryMiddleware(handler)
	handler = middleware.CorsMiddleware(handler)
	

	port := app.GetEnv("APP_PORT", ":8080")
	fmt.Printf("[SERVER] berjalan di: %s%s \n", app.GetEnv("APP_HOST", "localhost"), port)
	log.Fatal(http.ListenAndServe(port, handler))
}