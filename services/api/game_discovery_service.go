package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type GameMetadata struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Genre      string `json:"genre"`
	Popularity int    `json:"popularity"`
}

type DiscoveryService struct {
	rdb *redis.Client
}

func NewDiscoveryService(rdb *redis.Client) *DiscoveryService {
	return &DiscoveryService{rdb: rdb}
}

func (s *DiscoveryService) GetTrendingGames(ctx context.Context) ([]GameMetadata, error) {
	val, err := s.rdb.Get(ctx, "games:trending").Result()
	if err == nil {
		var games []GameMetadata
		if json.Unmarshal([]byte(val), &games) == nil {
			return games, nil
		}
	}
	return []GameMetadata{}, nil
}