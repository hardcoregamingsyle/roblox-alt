package auth

import (
	"strings"
	"unicode"
	"golang.org/x/text/unicode/norm"
)

func SanitizeInput(input string) string {
	s := strings.Map(func(r rune) rune {
		if unicode.IsControl(r) && r != '\n' && r != '\t' { return -1 }
		return r
	}, input)
	return norm.NFKC.String(s)
}