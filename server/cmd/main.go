package main

import (
	"context"
	"log"
	"sync"

	"github.com/dxtym/yomu/server/api"
	"github.com/dxtym/yomu/server/db/store"
	"github.com/dxtym/yomu/server/internal"
	"github.com/redis/go-redis/v9"
)

func main() {
	// TODO: use custom logger
	config, err := internal.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.NewStore(config.PostgresAddr)
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{Addr: config.RedisAddr})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	server := api.NewServer(db, rdb, config)

	wg.Add(1)
	go func() {
		if err := server.Start(config.Address); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
