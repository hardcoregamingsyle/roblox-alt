package main

import (
	"encoding/json"
	"net/http"
)

// DTO prevents mass assignment
type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Proceed with registration using only validated DTO fields...
}