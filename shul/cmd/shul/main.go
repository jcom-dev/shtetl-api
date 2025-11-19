package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// Get ports from environment or use defaults
	restPort := getEnv("SHUL_REST_PORT", "8103")
	grpcPort := getEnv("SHUL_GRPC_PORT", "8104")

	// Start REST API server in a goroutine
	go func() {
		mux := http.NewServeMux()

		// Health check endpoint
		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"status":"ok","service":"shul","version":"0.1.0"}`)
		})

		// Placeholder REST endpoints (to be implemented in Story 1.4)
		mux.HandleFunc("/api/v1/shuls", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprintf(w, `{"error":"not implemented yet"}`)
		})

		mux.HandleFunc("/api/v1/minyanim", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprintf(w, `{"error":"not implemented yet"}`)
		})

		mux.HandleFunc("/api/v1/schedules", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotImplemented)
			fmt.Fprintf(w, `{"error":"not implemented yet"}`)
		})

		fmt.Printf("Shul REST API listening on port %s...\n", restPort)
		if err := http.ListenAndServe(":"+restPort, mux); err != nil {
			log.Fatalf("REST server failed: %v", err)
		}
	}()

	// Create gRPC server
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port: %v", err)
	}

	grpcServer := grpc.NewServer()
	// TODO: Register gRPC services here (Story 1.4 - API Contract Design)

	fmt.Printf("Shul gRPC API listening on port %s...\n", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
