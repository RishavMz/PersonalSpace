package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

const (
	FILE_PATH = "log/server.log"
	FILE_PERMISSION = 0666
	STOP_SIGNAL_CHANNEL_SIZE = 1
)

func ConfigureLogger() *os.File {
	file, err := os.OpenFile(FILE_PATH, os.O_APPEND, FILE_PERMISSION)
	if err != nil {
		fmt.Println("Error opening log file:", err)
	}
	return file;
}

func ReadEnvironmentVariables() {
	log.Println("Reading environment variables")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if os.Getenv("PORT") == "" {
		log.Fatal("PORT not defined in environment")
	}
}

func RunServer() {
	port := os.Getenv("PORT")
	server := http.Server{Addr: ":" + port}

	go func() {
		log.Println("Starting server on port:", port)
		fmt.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, STOP_SIGNAL_CHANNEL_SIZE)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	fmt.Println("Shutting down server...")

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}
	log.Println("Server gracefully stopped")
}
