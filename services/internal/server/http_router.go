package server

import (
	"github.com/badrchoubai/services/internal/observability"
	"github.com/badrchoubai/services/internal/observability/logging/zap"
	"net/http"

	"github.com/badrchoubai/services/internal/middleware"
	"github.com/badrchoubai/services/internal/services"
)

func NewRouter(logger *logging.Logger, service services.ServiceInterface) http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux, service)

	var handler http.Handler = mux
	handler = middleware.Heartbeat(handler, "/health")
	handler = observability.RequestLoggingMiddleware(handler, logger)

	return handler
}

// addRoutes is where the entire API surface is mapped
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo
func addRoutes(mux *http.ServeMux, service services.ServiceInterface) {
	if service != nil {
		addServiceRoutes(mux, service)
	}

	mux.Handle("/*", http.NotFoundHandler())
}

func addServiceRoutes(mux *http.ServeMux, service services.ServiceInterface) {
	service.RegisterRouter(mux)
}
