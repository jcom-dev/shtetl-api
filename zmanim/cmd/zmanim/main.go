package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from environment or use default
	restPort := getEnv("ZMANIM_REST_PORT", "8101")

	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","service":"zmanim","version":"0.1.0"}`)
	})

	// Placeholder REST endpoints (to be implemented in Story 1.6)
	mux.HandleFunc("/api/v1/zmanim/calculate", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, `{"error":"not implemented yet"}`)
	})

	mux.HandleFunc("/api/v1/zmanim/streams", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, `{"error":"not implemented yet"}`)
	})

	fmt.Printf("Zmanim REST API listening on port %s...\n", restPort)
	if err := http.ListenAndServe(":"+restPort, mux); err != nil {
		log.Fatalf("REST server failed: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
