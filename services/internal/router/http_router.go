package router

import (
	"net/http"

	"github.com/badrchoubai/services/internal/observability/logging/zap"
	"github.com/badrchoubai/services/internal/service"
)

var _ IRouter = (*Router)(nil)

type Router struct {
	Name       string
	Handler    *http.ServeMux
	Logger     *logging.Logger
	Service    *service.Service
	Middleware []func(http.Handler) http.Handler
}

type IRouter interface {
	ApplyMiddleware(http.Handler) http.Handler
	WithOptions(opts ...Option) *Router

	clone() *Router
}

func NewRouter(name string, opts ...Option) *Router {
	router := &Router{
		Name:    name,
		Handler: http.NewServeMux(),
	}

	return router.WithOptions(opts...)
}

func (r *Router) ApplyMiddleware(handler http.Handler) http.Handler {
	// Apply middleware in reverse order, so the first middleware added
	// is the outermost one in the chain.
	for i := len(r.Middleware) - 1; i >= 0; i-- {
		handler = r.Middleware[i](handler)
	}

	return handler
}

// WithOptions clones the current Router, applies the supplied Options, and
// returns the resulting Router. It's safe to use concurrently.
func (r *Router) WithOptions(opts ...Option) *Router {
	c := r.clone()
	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}

func (r *Router) clone() *Router {
	clone := *r
	return &clone
}

// addRoutes is where the entire API surface is mapped
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo
func addRoutes(mux *http.ServeMux) {
	mux.Handle("/*", http.NotFoundHandler())
}
