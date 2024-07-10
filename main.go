package main

import (
	apí "github.com/rafaapcode/goAPi/api"
	"log"
)

func main() {
	server := apí.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
