package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"ok","service":"kehilla"}`)
	})

	// TODO: Add REST API endpoints (Story 1.3 - API Contract Design)
	// - GET /api/v1/schedules/{shulId}/today
	// - GET /api/v1/schedules/{shulId}/week
	// - POST /api/v1/subscriptions

	fmt.Println("Kehilla Service (REST) listening on port 8003...")
	if err := http.ListenAndServe(":8003", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
