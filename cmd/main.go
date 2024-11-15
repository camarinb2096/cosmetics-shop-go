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
		log.Fatal("error loading .env file:", err)
	}

	db, err := app.InitDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create and start server
	router := app.CreateRouter(db)

	// Start the server
	appRouter := &app.Router{}
	if err := appRouter.StartServer(router, "8080"); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
