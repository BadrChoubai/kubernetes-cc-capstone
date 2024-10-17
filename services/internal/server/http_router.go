package server

import (
	"net/http"

	"github.com/badrchoubai/services/internal/middleware"
)

// addRoutes is where the entire API surface is mapped
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo
func addRoutes(mux *http.ServeMux) {
	mux.Handle("/*", http.NotFoundHandler())
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux)

	var handler http.Handler = mux
	handler = middleware.Heartbeat(handler, "/health")

	return handler
}
