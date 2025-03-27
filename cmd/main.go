package main

import (
	"log"
	"student-service/internal/server"
)

func main() {
	// Start the gRPC server
	server.StartServer()

	log.Println("Server started on port 50051")
}
