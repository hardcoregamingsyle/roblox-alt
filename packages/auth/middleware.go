package auth

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// Fix 4: Atomic Redis Session Logic
var sessionLuaScript = `
local exists = redis.call("EXISTS", KEYS[1])
if exists == 0 then
    redis.call("SET", KEYS[1], ARGV[1], "EX", ARGV[2])
    return 1
else
    return 0
end
`

func ValidateSession(ctx context.Context, rdb *redis.Client, jti string, ttl int) (bool, error) {
	res, err := rdb.Eval(ctx, sessionLuaScript, []string{jti}, "active", ttl).Int()
	return res == 1, err
}