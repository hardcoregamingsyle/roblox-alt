package main

import (
	"net/http"
	"github.com/jackc/pgx/v5"
)

func UpdateGameState(w http.ResponseWriter, r *http.Request, db *pgx.Conn, userID string, gameID string) {
	// IDOR FIX: Verify ownership
	var exists bool
	err := db.QueryRow(r.Context(), "SELECT EXISTS(SELECT 1 FROM games WHERE id = $1 AND owner_id = $2)", gameID, userID).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "FORBIDDEN", http.StatusForbidden)
		return
	}
	// ... proceed with update
}