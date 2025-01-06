package main

import (
	"log"

	"github.com/dxtym/yomu/server/api/server"
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := internal.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewStore(config.PostgresAddr)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: give password
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: "",
		DB:       0,
	})

	scrape := internal.NewScrape()
	server := server.NewServer(db, rdb, scrape, config)
	if err := server.Start(config.Address); err != nil {
		log.Fatal(err)
	}
}
