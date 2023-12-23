package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[SERVER]: Failed to load Environment variables")
	} else {
		log.Println("[SERVER]: Successfully loaded environment variables")
	}
}

func StartServer(router http.Handler) {
	port := os.Getenv("SERVER_PORT")
	log.Printf("[SERVER]: Server is running on port %s\n", port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Use a context to handle shutdown
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals to initiate a graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		log.Println("[SERVER]: Shutting down gracefully...")
		cancel()
		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer timeoutCancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Printf("[SERVER]: Error during server shutdown: %v\n", err)
		}
	}()

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[SERVER]: Error starting server: %v\n", err)
	}
}
