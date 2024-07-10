package main

import (
	apí "github.com/rafaapcode/goAPi/api"
	"github.com/rs/zerolog/log"
)

func main() {
	server := apí.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server")
	}
}
