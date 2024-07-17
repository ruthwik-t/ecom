package main

import (
	"log"

	"github.com/ruthwik-t/ecom/cmd/api"
	"github.com/ruthwik-t/ecom/config"
)

func main() {

	server := api.NewAPIServer(config.Envs.ServerPort, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
