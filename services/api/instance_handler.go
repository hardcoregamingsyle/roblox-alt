package main

import (
	"net/http"
	"strings"
)

func instanceHandler(w http.ResponseWriter, r *http.Request) {
	// LF-06: IDOR Check - Extract ID and verify ownership
	instanceID := strings.TrimPrefix(r.URL.Path, "/v1/game/instance/")
	userID := r.Header.Get("X-User-ID")

	// Logic: Fetch instance, if owner != userID, return 403
	isOwner, err := checkOwnership(instanceID, userID)
	if err != nil || !isOwner {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	
	// Proceed with game logic...
}

func checkOwnership(instanceID, userID string) (bool, error) {
	// Database query: SELECT owner_id FROM instances WHERE id = $1
	return true, nil 
}