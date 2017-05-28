package authapi

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents an API server
type Server struct {
	*mux.Router
}

// New returns a new Server
func New() *Server {
	s := Server{mux.NewRouter()}
	s.setupRoutes()
	return &s
}

// Run API Server
func Run(listen string) error {
	s := New()
	log.Printf("Server listening on %s", listen)

	return http.ListenAndServe(listen, s.Router)
}

func (s *Server) setupRoutes() {
	log.Printf("Initialize Routings...")
	r := s.Router
	r.HandleFunc(`/`, healthCheckHandler)

	// /user/...
	user := r.PathPrefix(`/user`).Subrouter()
	user.HandleFunc(`/`, postUserHandler).
		Methods("POST")
	user.HandleFunc(`/`, deleteUserHandler).
		Methods("DELETE")
	user.HandleFunc(`/list`, getUserListHandler).
		Methods("GET")

	r.HandleFunc(`/auth`, postAuthHandler).
		Methods("POST")
	r.HandleFunc(`/algorithm`, getAlgorithmHandler).
		Methods("GET")
	r.HandleFunc(`/alg`, getAlgorithmHandler). // alias to /algorithm
		Methods("GET")
	r.HandleFunc(`/verify`, postVerifyHandler).
		Methods("POST")
	r.HandleFunc(`/key`, getKeyHandler).
		Methods("GET")
}
