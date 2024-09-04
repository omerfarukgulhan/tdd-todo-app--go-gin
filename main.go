package main

import (
	"log"
	"todo-app--go-gin/controller"
)

func main() {
	server := controller.InitializeRouter()

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
