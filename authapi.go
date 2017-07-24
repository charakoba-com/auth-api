package authapi

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nasa9084/go-logger"
)

// Server represents an API server
type Server struct {
	*mux.Router
}

var log *logger.Logger

func init() {
	log = logger.New(os.Stdout, "", logger.InfoLevel)
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
	log.Info("Server listening on %s", listen)

	return http.ListenAndServe(listen, s.Router)
}

func (s *Server) setupRoutes() {
	log.Info("Initialize Routings...")
	r := s.Router

	r.HandleFunc(`/`, HealthCheckHandler)

	// /user/...
	user := r.PathPrefix(`/user`).Subrouter()
	user.HandleFunc(`/list`, ListupUserHandler).
		Methods("GET")
	user.HandleFunc(``, CreateUserHandler).
		Methods("POST")
	user.HandleFunc(`/{id}`, LookupUserHandler).
		Methods("GET")
	user.HandleFunc(``, UpdateUserHandler).
		Methods("PUT")
	user.HandleFunc(`/{id}`, DeleteUserHandler).
		Methods("DELETE")

	r.HandleFunc(`/auth`, AuthHandler)
	r.HandleFunc(`/algorithm`, GetAlgorithmHandler)
	r.HandleFunc(`/alg`, GetAlgorithmHandler) // alias to /algorithm
	r.HandleFunc(`/verify`, VerifyHandler)
	r.HandleFunc(`/key`, GetKeyHandler)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
}
