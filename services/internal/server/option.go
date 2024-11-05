package server

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/badrchoubai/services/internal/service"
)

// Option represents a configuration option for a Service.
// Each Option modifies a Service instance when applied.
type Option interface {
	apply(*Server)
}

type optionFunc func(*Server)

func (f optionFunc) apply(server *Server) {
	f(server)
}

// WithLogger returns an Option that sets the logger for a Server instance.
// It allows customization of the Server's logging behavior during initialization.
func WithLogger(logger *zap.Logger) Option {
	return optionFunc(func(s *Server) {
		s.logger = logger
	})
}

// WithMiddleware returns an Option that adds one or more middleware functions
// to a Server instance. The middleware functions are applied in the order
// they are provided, allowing for flexible customization of request handling.
func WithMiddleware(middleware ...func(http.Handler) http.Handler) Option {
	return optionFunc(func(server *Server) {
		server.middlewares = append(server.middlewares, middleware...)
	})
}

// WithService returns an Option that adds a Service instance to the Server.
// This allows the Server to register and manage the provided Service,
// enabling it to handle requests associated with that Service.
func WithService(srv *service.Service) Option {
	return optionFunc(func(server *Server) {
		server.services = append(server.services, srv)
	})
}
