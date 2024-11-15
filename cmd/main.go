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

	err = router.StartServer()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
