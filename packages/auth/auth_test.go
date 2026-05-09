package auth

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"Valid", "nexususer123", "nexususer123", false},
		{"Normalization", "NEXUSUSER", "nexususer", false},
		{"FullWidth", "ＮＥＸＵＳ", "nexus", false},
		{"TooShort", "ab", "", true},
		{"TooLong", "a123456789012345678901234567890123456789012345678901", "", true},
		{"InvalidChars", "user_name!", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateUsername(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestPasswordSecurity(t *testing.T) {
	password := "SecurePass123!"
	encoded, err := HashPassword(password)
	assert.NoError(t, err)
	assert.Greater(t, len(encoded), SaltLen)

	valid, err := VerifyPassword(password, encoded)
	assert.NoError(t, err)
	assert.True(t, valid)

	invalid, err := VerifyPassword("wrong", encoded)
	assert.NoError(t, err)
	assert.False(t, invalid)
}