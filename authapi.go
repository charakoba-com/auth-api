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

	r.HandleFunc(`/`, HealthCheckHandler)

	// /user/...
	user := r.PathPrefix(`/user`).Subrouter()
	user.HandleFunc(``, CreateUserHandler).
		Methods("POST")
	user.HandleFunc(`/{id}`, LookupUserHandler).
		Methods("GET")
	user.HandleFunc(``, UpdateUserHandler).
		Methods("PUT")
	user.HandleFunc(`/{id}`, DeleteUserHandler).
		Methods("DELETE")
	user.HandleFunc(`/list`, ListupUserHandler)

	r.HandleFunc(`/auth`, AuthHandler)
	r.HandleFunc(`/algorithm`, GetAlgorithmHandler)
	r.HandleFunc(`/alg`, GetAlgorithmHandler) // alias to /algorithm
	r.HandleFunc(`/verify`, VerifyHandler)
	r.HandleFunc(`/key`, GetKeyHandler)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
}
