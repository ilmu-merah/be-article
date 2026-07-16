package auth

import (
	"encoding/json"
	"net/http"
)


func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":"ini api login",
	})
}

func Register(w http.ResponseWriter, r *http.Request) {}