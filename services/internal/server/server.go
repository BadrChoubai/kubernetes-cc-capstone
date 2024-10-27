package server

import (
	"net"
	"net/http"
	"strconv"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/observability/logging"
	"github.com/badrchoubai/services/internal/service"
)

var _ IServer = (*Server)(nil)

type Server struct {
	config      *config.AppConfig
	httpServer  *http.Server
	logger      *logging.Logger
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
	services    []*service.Service
}

type IServer interface {
	ApplyMiddleware(http.Handler) http.Handler
	WithOptions(opts ...Option) *Server
	HttpServer() *http.Server

	clone() *Server
}

func NewServer(cfg *config.AppConfig, opts ...Option) *Server {
	server := &Server{
		config: cfg,
		mux:    http.NewServeMux(),
		httpServer: &http.Server{
			Addr:    net.JoinHostPort(cfg.HTTPHost(), strconv.Itoa(cfg.HTTPPort())),
			Handler: nil,
		},
	}
	server = server.WithOptions(opts...)

	for _, svc := range server.services {
		server.mux.Handle(svc.URL()+"/", svc.Mux()) // Register with service URL prefix
	}

	server.httpServer.Handler = server.ApplyMiddleware(server.mux)

	return server
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

func (s *Server) HttpServer() *http.Server {
	return s.httpServer
}
