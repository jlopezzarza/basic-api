package server

import (
	"fmt"
	"net/http"
)

// Server holds a basic webserver
type Server struct {
	router *http.ServeMux
}

func (s *Server) loadRoutes() {
	s.router = http.NewServeMux()
	s.router.HandleFunc("/ping", s.ping)
}

func (s *Server) ping(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "pong")
}

// New creates a new instance of the server
func New() (s *Server) {
	s = &Server{}
	s.loadRoutes()
	return
}

// Run starts the server
func (s *Server) Run() (err error) {
	return http.ListenAndServe(":8080", s.router)
}
