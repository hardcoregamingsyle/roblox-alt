package auth

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
	"golang.org/x/text/unicode/norm"
)

var allowedChars = regexp.MustCompile(`^[a-zA-Z0-9_]{3,50}$`)
var controlChars = regexp.MustCompile(`[\p{C}]`)

func ValidateUsername(u string) (string, error) {
	if !utf8.ValidString(u) {
		return "", fmt.Errorf("invalid UTF-8 sequence")
	}
	
	u = strings.ReplaceAll(u, "\x00", "")
	normalized := norm.NFKC.String(u)
	
	if controlChars.MatchString(normalized) {
		return "", fmt.Errorf("invalid control characters")
	}
	
	if !allowedChars.MatchString(normalized) {
		return "", fmt.Errorf("invalid characters or length")
	}
	return normalized, nil
}