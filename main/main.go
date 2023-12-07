package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a function to handle the health check endpoint
	http.HandleFunc("/health", healthcheck.HealthCheckHandler)

	// Start the HTTP server on port 8080
	port := 8080
	fmt.Printf("Server listening on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
