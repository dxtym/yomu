package main

import (
	"log"

	"github.com/dxtym/yomu/server/api"
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
)

func main() {
	config, err := internal.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	store, err := db.NewStore(config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	token, err := internal.NewToken(config.SecretKey)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(store, token, config)
	if err := server.Start(config.Address); err != nil {
		log.Fatal(err)
	}
}
