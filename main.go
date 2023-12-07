package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	// Define a function to handle the health check endpoint
	http.HandleFunc("/health", healthCheckHandler)

	// Start the HTTP server on port 8080
	port := 8080
	fmt.Printf("Server listening on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// HealthInfo represents the health information to be included in the response
type HealthInfo struct {
	Version          string `json:"version"`
	GoroutinesCount  int    `json:"goroutines_count"`
	TotalAllocBytes  uint64 `json:"total_alloc_bytes"`
	HeapObjectsCount uint64 `json:"heap_objects_count"`
	AllocBytes       uint64 `json:"alloc_bytes"`
}

// healthCheckHandler is the handler function for the /health endpoint
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Perform health check logic here
	// You can check the status of your library or any other dependencies

	// Gather runtime information
	version := runtime.Version()
	goroutinesCount := runtime.NumGoroutine()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Create a HealthInfo struct with the collected information
	healthInfo := HealthInfo{
		Version:          version,
		GoroutinesCount:  goroutinesCount,
		TotalAllocBytes:  m.TotalAlloc,
		HeapObjectsCount: m.HeapObjects,
		AllocBytes:       m.Alloc,
	}

	// Convert the HealthInfo struct to JSON
	responseJSON, err := json.Marshal(healthInfo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response with the health information
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
