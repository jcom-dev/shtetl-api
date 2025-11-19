package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from environment or use default
	port := getEnv("KEHILLA_PORT", "8105")

	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","service":"kehilla","version":"0.1.0"}`)
	})

	// Placeholder REST endpoints (to be implemented in Story 1.4)
	mux.HandleFunc("/api/v1/schedules/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, `{"error":"not implemented yet"}`)
	})

	mux.HandleFunc("/api/v1/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, `{"error":"not implemented yet"}`)
	})

	fmt.Printf("Kehilla REST API listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
