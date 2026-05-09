package main

import (
	"context"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/redis/go-redis/v9"
	"time"
)

type ACLKey struct {
	UserID     string
	InstanceID string
}

var aclCache = expirable.New[ACLKey, bool](1000, nil, 5*time.Minute)

func IsUserInInstance(ctx context.Context, rdb *redis.Client, userID string, instanceID string) (bool, error) {
	key := ACLKey{UserID: userID, InstanceID: instanceID}
	if val, ok := aclCache.Get(key); ok {
		return val, nil
	}

	isMember, err := rdb.SIsMember(ctx, "acl:"+instanceID, userID).Result()
	if err != nil {
		return false, err
	}

	aclCache.Add(key, isMember)
	return isMember, nil
}