package main

import (
	"artubric/transmission-rest/config"

	"artubric/transmission-rest/internal/server"
	tserv "artubric/transmission-rest/internal/transmissionService"
)

func main() {
	conf := config.Load()
	transService := tserv.New(&conf.TransmissionServer)
	server := server.New(&conf.RestServer, transService)
	server.Run()
}
