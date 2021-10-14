package main

import (
	"artubric/transmission-rest/config"

	"artubric/transmission-rest/internal/server"
	tserv "artubric/transmission-rest/internal/transmissionService"
)

func main() {
	conf := config.Load()
	transmissionService := tserv.New(&conf.TransmissionServer)
	server := server.New(&conf.RestServer, transmissionService)
	server.Run()
}
