package routes

import (
	"net/http"

	"github.com/mik-sea/go-firebase-backend/internal/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/api/users", handlers.CreateUser)
}

func GetAllTransactions() {
	http.HandleFunc("/api/transactions", handlers.GetTransactions)
}
