package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Injected at build time via -ldflags in the Dockerfile
var (
	version   = "dev"
	commitSHA = "unknown"
	buildTime = "unknown"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type VersionResponse struct {
	Version   string `json:"version"`
	CommitSHA string `json:"commit_sha"`
	BuildTime string `json:"build_time"`
	Hostname  string `json:"hostname"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(VersionResponse{
		Version:   version,
		CommitSHA: commitSHA,
		BuildTime: buildTime,
		Hostname:  hostname,
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ECS Fargate Blue-Green Demo | version=%s commit=%s\n", version, commitSHA)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/version", versionHandler)
	mux.HandleFunc("/test", versionHandler)
	mux.HandleFunc("/", rootHandler)

	log.Printf("Starting server on :%s  version=%s  commit=%s  built=%s\n",
		port, version, commitSHA, buildTime)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
