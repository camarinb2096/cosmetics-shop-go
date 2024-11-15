package app

import (
	"camarinb2096/cosmetics-shop-go/internal/entities"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Router structure contains the HTTP server and its port.
type Router struct {
	server *http.Server
	port   string
}

var (
	ErrStartingServer = errors.New("error starting server")
)

// InitDB initializes a SQLite database connection using GORM.
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to SQLite database: " + err.Error())
	}

	// AutoMigrate creates tables for the specified entities.
	err = db.AutoMigrate(
		&entities.Buyer{},
		&entities.Customer{},
		&entities.Invoice{},
		&entities.Product{},
	)
	if err != nil {
		return nil, errors.New("failed to auto-migrate database: " + err.Error())
	}

	log.Println("Database initialized in memory")
	return db, nil
}

// CreateRouter initializes the Gin router and configures endpoints.
func CreateRouter() *gin.Engine {
	router := gin.Default()

	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	return router
}

// StartServer starts the HTTP server using Gin.
func (r *Router) StartServer(router *gin.Engine, port string) error {
	if port == "" {
		return errors.New("server port is not defined")
	}

	r.port = port
	r.server = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server running on port " + port)
	if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("Error starting server:", err)
		return err
	}

	return nil
}
