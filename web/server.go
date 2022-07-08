package web

import (
	"net/http"
	_"encoding/json"
	"wb_test/pkg/cache"
	"log"
)

type Server struct {
	host string
	port string
	cache *cache.Cache
}

func NewServer(host, port string) *Server {
	return &Server{
		host: host,
		port: port,
		cache: cache.NewCache(),
	}	
}

func (s *Server) Run() error {
	
	serv := &http.Server {
		Addr: s.host + ":" + s.port,
		Handler: s.NewRouter(),
	}
	
	log.Println("Run server, host:", s.host, "port:", s.port)
	return serv.ListenAndServe()
}