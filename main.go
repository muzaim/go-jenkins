package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status:  "success",
		Message: "API is healthy and running with Jenkins CI!",
	})
}

func main() {
	http.HandleFunc("/health", HealthCheckHandler)
	fmt.Println("Test aja bro")
	fmt.Println("Server running on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
