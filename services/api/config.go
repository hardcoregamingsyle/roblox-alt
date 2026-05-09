package main

import (
	"fmt"
	"os"
)

// Fix: Helper to ensure critical environment variables exist
func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("CRITICAL: Environment variable %s not set", key))
	}
	return val
}