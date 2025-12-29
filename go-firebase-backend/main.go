package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mik-sea/go-firebase-backend/internal/config"
	"github.com/mik-sea/go-firebase-backend/routes"
)

func main() {
	godotenv.Load()
	config.InitFirebase()
	// routes.RegisterRoutes()
	routes.GetAllTransactions()

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
