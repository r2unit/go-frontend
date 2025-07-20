package main

import (
	"log"
	"net/http"

	"github.com/nerdpitch-cloud/frontend/pkg/framework"
)

func main() {
	// Create a new framework instance
	fw, err := framework.New("config/routes.json")
	if err != nil {
		log.Fatalf("Error creating framework: %v", err)
	}

	// Register handlers
	fw.RegisterHandlers()

	// Apply middleware
	handler := fw.NotFoundRedirectMiddleware(http.DefaultServeMux)

	log.Println("Starting frontend on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
