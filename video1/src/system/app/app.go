package app

import (
	"log"
	"net/http"
)

// Server - server
type Server struct {
	port string
}

// NewServer - new server
func NewServer() Server {
	return Server{}
}

// Init - init server
func (s *Server) Init() {
	log.Println("Initializing server...")
	s.port = ":8000"
}

// Start - start server
func (s *Server) Start() {
	log.Println("Starting server...")
	http.ListenAndServe(s.port, nil)
}
