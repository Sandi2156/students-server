package main

import (
	"log"
	"net/http"

	"github.com/sandipan/students-api/internal/config"
	student "github.com/sandipan/students-api/internal/http/handlers"
)

func main() {
	// load configuration
	cfg := config.MustLoad()
	// database setup
	// setup routes
	router := http.NewServeMux()

	router.HandleFunc("POST /v1/api/students", student.New())

	
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

	// done := make(chan os.Signal, 1)

	// signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	
	// func () {
	// 	log.Printf("Server listening on %v", cfg.Address)
	// 	err := server.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatal("Failed to start server.")
	// 	}
	// }()

	// <- done

	// slog.Info("Shutting down the server")

	// ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	// defer cancel()

	// if err := server.Shutdown(ctx); err != nil {
	// 	slog.Error("We are unable to shutdown the server. ", slog.String("error ", err.Error()))
	// } 

	// slog.Info("Shut down the server gracefully.")

}