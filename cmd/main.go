package main

import (
	"log"

	"github.com/ruthwik-t/ecom/cmd/api"
	"github.com/ruthwik-t/ecom/config"
	"github.com/ruthwik-t/ecom/db"
)

func main() {

	db, err := db.ConnectToPostgreSQL(config.Envs.DbString)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(config.Envs.ServerPort, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
