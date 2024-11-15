package main

import (
	"camarinb2096/cosmetics-shop-go/cmd/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Crear y arrancar el servidor
	router := app.CreateRouter()

	err = router.StartServer()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
