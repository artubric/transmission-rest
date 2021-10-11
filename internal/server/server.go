package server

import (
	"artubric/transmission-rest/config"
	tserv "artubric/transmission-rest/internal/transmissionService"
	"fmt"
	"net/http"
)

type Server struct {
	conf      *config.RestServer
	transServ *tserv.TransmissionService
}

func New(conf *config.RestServer, transServ *tserv.TransmissionService) *Server {
	return &Server{conf: conf, transServ: transServ}
}

func (s *Server) Run() {
	s.setupRoutes()
	address := fmt.Sprintf(":%d", s.conf.Port)
	http.ListenAndServe(address, nil)
}

func (s *Server) setupRoutes() {
	currentAPIVersion := "v1"

	urlPath := fmt.Sprintf("/%s/%s/%s",
		s.conf.ApiBasePath,
		currentAPIVersion,
		s.conf.MainEntity,
	)

	fmt.Println("Setting up handlers")
	handleTorrents := http.HandlerFunc(s.handleTorrentsV1)
	fmt.Println("Done")

	fmt.Println("Starting up server on ", urlPath)
	http.Handle(urlPath, handleTorrents)
	fmt.Println("Done")

}
