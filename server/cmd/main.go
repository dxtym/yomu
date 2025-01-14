package main

import (
	"log"
	"sync"

	"github.com/dxtym/yomu/server/api"
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
	"github.com/redis/go-redis/v9"
)

func main() {
	var wg sync.WaitGroup
	config, err := internal.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewStore(config.PostgresAddr)
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})

	scrape := internal.NewScrape()
	server := api.NewServer(db, rdb, scrape, config)

	wg.Add(1)
	go func() {
		if err := server.Start(config.Address); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
