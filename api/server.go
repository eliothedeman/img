package api

import "net/http"

// Server serves api requests via http
type Server struct {
}

// Server serve the api
func (s *Server) Serve(port string) {
	http.HandleFunc("/", Router)
	http.ListenAndServe(":"+port, nil)
}
