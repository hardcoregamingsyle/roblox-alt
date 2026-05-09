/**
 * @file game_discovery_handler.go
 * @description Provides high-performance game discovery endpoints with Redis-backed caching.
 */
package main

import (
	"encoding/json"
	"net/http"
	"nexus-engine/packages/auth"
	"sync"
)

// Global encoder pool to minimize heap allocation pressure during high-concurrency requests.
var encoderPool = sync.Pool{
	New: func() interface{} { return json.NewEncoder(nil) },
}

/**
 * DiscoveryHandler returns a paginated list of trending games.
 * Enforces per-IP rate limiting to protect against scraper bots.
 */
func DiscoveryHandler(svc *DiscoveryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Enforce Rate Limiting (5 req/sec)
		limiter := auth.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		games, err := svc.GetTrendingGames(r.Context())
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		enc := encoderPool.Get().(*json.Encoder)
		enc.Reset(w)
		defer encoderPool.Put(enc)
		
		enc.Encode(games)
	}
}