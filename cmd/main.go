package main

import (
	"log"
	"student-service/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start the gRPC server
	server.StartServer()
}
