package main

import (
	"log"
	"net/http"

	"github.com/sandipan/students-api/internal/config"
)

func main() {
	// load configuration
	cfg := config.MustLoad()
	// database setup
	// setup routes
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})
	// setup server
	server := http.Server {
		Addr: cfg.Address,
		Handler: router,
	}

	log.Printf("Server listening on %v", cfg.Address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server.")
	}
}