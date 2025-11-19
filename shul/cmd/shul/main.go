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
		fmt.Fprintf(w, `{"status":"ok","service":"shul"}`)
	})

	// TODO: Add REST API endpoints (Story 1.3 - API Contract Design)

	fmt.Println("Shul Service (REST) listening on port 8002...")
	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
