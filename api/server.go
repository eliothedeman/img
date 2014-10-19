package api

import (
	"log"
	"net/http"
)

// Server serves api requests via http
type Server struct {
}

// Serve serve the api
func (s *Server) Serve(port string) {
	http.HandleFunc("/", Router)
	err := http.ListenAndServe(":"+port, nil)
	log.Println(err)
}
