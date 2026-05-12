package auth

import (
	"context"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var allowedKids map[string]bool

func init() {
	allowedKids = make(map[string]bool)
}

func ValidateJWT(ctx context.Context, tokenStr string) (*NexusClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &NexusClaims{}, func(t *jwt.Token) (interface{}, error) {
		kid, ok := t.Header["kid"].(string)
		if !ok || !allowedKids[kid] { return nil, errors.New("invalid kid") }
		return FetchPublicKeyFromJWKS(kid), nil
	}, jwt.WithIssuer("NexusEngine"), jwt.WithAudience("api.nexusengine.com"), jwt.WithLeeway(time.Second * 30))

	if err != nil || !token.Valid { return nil, err }
	return token.Claims.(*NexusClaims), nil
}