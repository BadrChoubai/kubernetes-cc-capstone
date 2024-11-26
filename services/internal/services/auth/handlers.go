package auth

import (
	"net/http"

	"github.com/badrchoubai/services/internal/service"
)

func addRoutes(svc *service.Service) {
	svc.Mux().Handle("/", http.NotFoundHandler())
}
