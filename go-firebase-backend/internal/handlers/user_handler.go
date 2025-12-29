package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mik-sea/go-firebase-backend/internal/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user services.User
	json.NewDecoder(r.Body).Decode(&user)

	err := services.SaveUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created",
	})
}
