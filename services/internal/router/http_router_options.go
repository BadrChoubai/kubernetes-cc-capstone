package router

import (
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
	"github.com/badrchoubai/services/internal/service"
	"net/http"
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
		router.logger = logger
	})
}

func WithMiddleware(middleware func(http.Handler) http.Handler) Option {
	return optionFunc(func(router *Router) {
		router.middleware = append(router.middleware, middleware)
	})
}

func WithService(service *service.Service) Option {
	return optionFunc(func(router *Router) {
		router.service = service
	})
}
