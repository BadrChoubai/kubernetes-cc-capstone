package auth

import (
	"net/http"

	"github.com/badrchoubai/services/internal/service"
)

func addRoutes(svc *service.Service) {
	svc.Mux().Handle("/", http.NotFoundHandler())
	svc.Mux().Handle("POST /token", generateTokenHandler())
}

func generateTokenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	}
}
