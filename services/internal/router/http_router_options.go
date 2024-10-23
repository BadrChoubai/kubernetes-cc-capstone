package router

import (
	"net/http"

	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
	"github.com/badrchoubai/services/internal/service"
)

type Option interface {
	apply(*Router)
}

type optionFunc func(*Router)

func (f optionFunc) apply(service *Router) {
	f(service)
}

func WithLogger(logger *logging.Logger) Option {
	return optionFunc(func(router *Router) {
		router.Logger = logger
	})
}

func WithService(service *service.Service) Option {
	return optionFunc(func(router *Router) {
		router.Service = service
	})
}

func WithMiddleware(middleware func(http.Handler) http.Handler) Option {
	return optionFunc(func(router *Router) {
		router.Middleware = append(router.Middleware, middleware)
	})
}
