package server

import (
	"net/http"
	"time"

	"service/server/middleware/json"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Config to apply to the created Server.
type Config struct {
	Addr string

	Logger *zap.Logger
}

// Server harbours all dependencies and service used in serving requests.
type Server struct {
	addr string

	r      *mux.Router
	logger *zap.Logger
}

// New creates a configured Server with the passed endpoints.
func New(config *Config, endpoints ...Endpoint) (*Server, error) {
	s := &Server{
		addr: config.Addr,
		r:    mux.NewRouter(),

		logger: config.Logger,
	}

	// Assign the middlewares.
	s.r.Use(json.Middleware(s.logger))
	s.r.Use(mux.CORSMethodMiddleware(s.r))

	// Attach the endpoints.
	for _, endpoint := range endpoints {
		s.r.Handle(endpoint.Path(), endpoint).Methods(endpoint.Methods()...)
	}

	return s, nil
}

// Start serving the server with responsible defaults.
func (s *Server) Start() error {
	serve := &http.Server{
		Addr:    s.addr,
		Handler: s.r,

		// Sensible defaults
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       20 * time.Second,
		MaxHeaderBytes:    32 * 1024,
	}

	return serve.ListenAndServe()
}
