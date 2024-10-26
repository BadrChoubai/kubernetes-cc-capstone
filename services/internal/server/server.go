package server

import (
	"fmt"
	"net/http"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/observability/logging"
	"github.com/badrchoubai/services/internal/service"
)

type Server struct {
	config      *config.AppConfig
	handler     *http.ServeMux
	logger      *logging.Logger
	middlewares []func(http.Handler) http.Handler
	service     *service.Service
}

type IServer interface {
	ApplyMiddleware(http.Handler) http.Handler
	Handler() *http.ServeMux
	WithOptions(opts ...Option) *http.ServeMux

	clone() *http.ServeMux
}

func NewServer(cfg *config.AppConfig, service *service.Service, opts ...Option) *Server {
	mainMux := http.NewServeMux()

	var servicePath string
	if service != nil {
		servicePath = fmt.Sprintf("/%s/", service.Name())
		mainMux.Handle(servicePath, service.Mux())
	}

	mainMux.Handle("/*", http.NotFoundHandler())

	server := &Server{
		config:  cfg,
		handler: mainMux,
	}

	return server.WithOptions(opts...)
}

func (s *Server) ApplyMiddleware(handler http.Handler) http.Handler {
	// Apply middleware in reverse order, so the first middleware added
	// is the outermost one in the chain.
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		handler = s.middlewares[i](handler)
	}

	return handler
}

// WithOptions clones the current Router, applies the supplied Options, and
// returns the resulting Router. It's safe to use concurrently.
func (s *Server) WithOptions(opts ...Option) *Server {
	c := s.clone()
	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}

func (s *Server) clone() *Server {
	clone := *s
	return &clone
}

func (s *Server) Handler() *http.ServeMux {
	return s.handler
}

func (s *Server) Service() *service.Service {
	return s.service
}
