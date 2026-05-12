package auth

import (
	"regexp"
	"strings"
	"golang.org/x/text/unicode/norm"
)

// Use anchored, non-backtracking regex pattern to prevent ReDoS
var usernameRegex = regexp.MustCompile(`^[a-z0-9]{3,32}$`)

func ValidateUsername(u string) bool {
	normalized := strings.ToLower(norm.NFKC.String(u))
	return usernameRegex.MatchString(normalized)
}