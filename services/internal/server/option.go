package server

import (
	"net/http"

	"github.com/badrchoubai/services/internal/observability/logging"
	"github.com/badrchoubai/services/internal/service"
)

type Option interface {
	apply(*Server)
}

type optionFunc func(*Server)

func (f optionFunc) apply(server *Server) {
	f(server)
}

func WithLogger(logger *logging.Logger) Option {
	return optionFunc(func(s *Server) {
		s.logger = logger
	})
}

func WithMiddleware(middleware func(http.Handler) http.Handler) Option {
	return optionFunc(func(server *Server) {
		server.middlewares = append(server.middlewares, middleware)
	})
}

func WithService(service *service.Service) Option {
	return optionFunc(func(server *Server) {
		server.services = append(server.services, service)
	})
}
