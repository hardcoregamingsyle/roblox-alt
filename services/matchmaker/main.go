package main

import (
    "net/http"
    "github.com/google/uuid"
)

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if _, err := uuid.Parse(id); err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }
    // Proceed with logic...
}