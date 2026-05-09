package main

import (
	"log/slog"
	"os"
	"strings"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func SanitizeLog(msg string) string {
	// Strip control characters and newlines to prevent log injection
	return strings.Map(func(r rune) rune {
		if r < 32 || r == 127 {
			return -1
		}
		return r
	}, msg)
}

func LogError(err error, fields ...any) {
	logger.Error(SanitizeLog(err.Error()), fields...)
}