package app

import (
	"net/http"
	"github.com/ilmu-merah/be-article/internal/auth"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", publicRouter()))
	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdminRouter()))

	return mux;
}

func publicRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// auth routes
	mux.HandleFunc("GET /auth/login", auth.Login)
	mux.HandleFunc("POST /auth/register", auth.Register)

	return mux
}


func onlyAdminRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux;
}