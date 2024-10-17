package server

import (
	"github.com/badrchoubai/services/internal/services"
	"net/http"

	"github.com/badrchoubai/services/internal/middleware"
)

// addRoutes is where the entire API surface is mapped
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo
func addRoutes(mux *http.ServeMux) {
	mux.Handle("/*", http.NotFoundHandler())
}

// addRoutes is where the entire API surface is mapped
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo
func addServiceRoutes(mux *http.ServeMux, service services.Service) {
	service.RegisterRouter(mux)
}

func NewRouter(service services.Service) http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux)
	addServiceRoutes(mux, service)

	var handler http.Handler = mux
	handler = middleware.Heartbeat(handler, "/health")

	return handler
}
