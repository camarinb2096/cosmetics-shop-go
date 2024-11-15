package app

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
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
func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to SQLite database:", err)
	}

	log.Println("database initialized in memory")
	return db
}

// CreateRouter initializes the router and returns a Router instance.
func CreateRouter() *Router {
	router := chi.NewRouter()

	// Endpoint for testing
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// HTTP server configuration
	return &Router{
		server: &http.Server{
			Addr:    ":" + os.Getenv("SERVER_PORT"),
			Handler: router,
		},
		port: os.Getenv("SERVER_PORT"),
	}
}

// StartServer starts the HTTP server.
func (r *Router) StartServer() error {
	if r.port == "" {
		return errors.New("server port is not defined")
	}

	log.Println("Server running on port " + r.port)

	// Start the server
	err := r.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println(ErrStartingServer.Error())
		return ErrStartingServer
	}

	return nil
}
