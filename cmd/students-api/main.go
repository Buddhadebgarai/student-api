package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Buddhadebgarai/student-api/internal/config"
)

func main() {
	fmt.Println("Welcome to the Student API!")

	// Load configuration
	cfg := config.MustLoad()
	fmt.Printf("Configuration loaded: %+v\n", cfg)
	// database setup

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WelCome to students api"))
	})
	//setup server

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	//making channel to listen for shutdown signal
	done := make(chan os.Signal, 1)

	// Listen for interrupt signal to gracefully shutdown the server
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		fmt.Println("Server is running on", cfg.HTTPServer.Addr)
		err := server.ListenAndServe()

		if err != nil {
			log.Fatalf("failed to start server: %v", err.Error())
		}
	}()

	<-done // wait for signal to shutdown server
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err.Error())
	}

	slog.Info("Server gracefully stopped")

}
