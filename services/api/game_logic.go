package main

import (
	"errors"
	"regexp"
)

type CreateGameRequest struct {
	GameName    string `json:"gameName"`
	PlayerCount int    `json:"playerCount"`
}

var gameNameRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]{3,64}$`)

func (g *CreateGameRequest) Validate() error {
	if !gameNameRegex.MatchString(g.GameName) {
		return errors.New("invalid game name: must be alphanumeric, 3-64 chars")
	}
	if g.PlayerCount < 1 || g.PlayerCount > 100 {
		return errors.New("invalid player count: must be between 1 and 100")
	}
	return nil
}