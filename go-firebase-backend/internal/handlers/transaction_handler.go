package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mik-sea/go-firebase-backend/internal/services"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	data, err := services.GetAllTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
