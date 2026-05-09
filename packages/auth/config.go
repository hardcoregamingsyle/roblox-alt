package auth

import (
	"log"
	"os"
)

func CheckSecret() {
	secret := os.Getenv("JWT_SIGNING_KEY")
	if len(secret) < 32 {
		log.Fatal("SECURITY CRITICAL: JWT_SIGNING_KEY must be at least 32 characters long")
	}
}