package main

import (
	"camarinb2096/cosmetics-shop-go/cmd/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	app.InitDB()

	// Create and start server
	router := app.CreateRouter()

	// Start the server
	appRouter := &app.Router{}
	if err := appRouter.StartServer(router, "8080"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
