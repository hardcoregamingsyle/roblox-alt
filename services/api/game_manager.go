package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

const transitionScript = `
local current = redis.call("GET", KEYS[1])
if current == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[2])
    return 1
end
return 0`

func (gm *GameManager) Transition(ctx context.Context, gameID string, current, next string) error {
	res, err := gm.rdb.Eval(ctx, transitionScript, []string{gameID}, current, next).Int()
	if err != nil || res == 0 { return fmt.Errorf("invalid transition or race condition") }
	return nil
}