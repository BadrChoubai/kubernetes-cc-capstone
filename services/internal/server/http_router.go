package server

import (
	"net/http"

	"github.com/badrchoubai/services/internal/middleware"
	"github.com/badrchoubai/services/internal/observability"
	"github.com/badrchoubai/services/internal/observability/logging/zap"
)

func NewRouter(logger *logging.Logger) http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux)

	var handler http.Handler = mux
	handler = middleware.Heartbeat(handler, "/health")
	handler = observability.RequestLoggingMiddleware(handler, logger)

	return handler
}

// addRoutes is where the entire API surface is mapped
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo
func addRoutes(mux *http.ServeMux) {
	mux.Handle("/*", http.NotFoundHandler())
}
