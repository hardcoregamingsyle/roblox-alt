package main

import (
	"net/http"
	"regexp"
	"github.com/google/uuid"
)

var cidRegex = regexp.MustCompile(`^[a-fA-F0-9-]{36}$`)

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := r.Header.Get("X-Correlation-ID")
		if cid == "" || !cidRegex.MatchString(cid) {
			cid = uuid.New().String()
		}
		
		w.Header().Set("X-Correlation-ID", cid)
		next.ServeHTTP(w, r)
	})
}