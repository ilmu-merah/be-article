package app

import (
	"net/http"
	"github.com/ilmu-merah/be-article/internal/auth"
)

func NewRouter(authHandler *auth.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", publicRouter(authHandler)))
	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdminRouter()))

	return mux;
}

func publicRouter(authHandler *auth.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// auth routes
	mux.HandleFunc("GET /auth/login", authHandler.Login)
	mux.HandleFunc("POST /auth/register", authHandler.Register)

	return mux
}


func onlyAdminRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux;
}