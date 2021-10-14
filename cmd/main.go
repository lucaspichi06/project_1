package main

import (
	"log"
	"os"

	"github.com/project_1/cmd/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := app.Start()
	if err := router.Run(":" + port); err != nil {
		log.Println("[main][message: error running server]", err)
	}
}
