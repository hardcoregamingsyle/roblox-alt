package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	dbReady = make(chan struct{})
	dbErr   error
)

func GetDB(timeout time.Duration) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-dbReady:
		if dbErr != nil {
			return nil, dbErr
		}
		return "db_instance", nil
	case <-ctx.Done():
		return nil, errors.New("database initialization timed out")
	}
}