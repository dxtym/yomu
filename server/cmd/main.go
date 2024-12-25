package main

import (
	"log"

	"github.com/dxtym/yomu/server/api"
	"github.com/dxtym/yomu/server/db"
	"github.com/dxtym/yomu/server/internal"
	"github.com/machinebox/graphql"
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

	client := graphql.NewClient(config.ApiUrl)
	server := api.NewServer(store, token, client)
	if err := server.Start(config.Address); err != nil {
		log.Fatal(err)
	}
}
