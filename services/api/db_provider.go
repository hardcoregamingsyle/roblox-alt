package main

import (
	"sync"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbInstance *pgxpool.Pool
	once       sync.Once
	ready      = make(chan struct{})
)

func InitDB(connString string) {
	once.Do(func() {
		pool, err := pgxpool.New(context.Background(), connString)
		if err != nil {
			panic(err)
		}
		dbInstance = pool
		close(ready)
	})
}

func GetDB() *pgxpool.Pool {
	<-ready
	return dbInstance
}