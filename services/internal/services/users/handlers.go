package users

import (
	"net/http"

	"github.com/badrchoubai/services/internal/service"
)

type (
	Response struct {
		Service string `json:"service"`
	}
)

func addRoutes(svc *service.Service) {
	svc.Mux().Handle("/health", Healthz(svc))
	svc.Mux().Handle("/", http.NotFoundHandler())
}

func Healthz(svc *service.Service) http.Handler {
	response := &Response{
		Service: svc.Name(),
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			respond(svc, w, http.StatusOK, response)
		})
}

// respond handles encoding the response and managing any encoding errors.
func respond(svc *service.Service, w http.ResponseWriter, statusCode int, data any) {
	if err := svc.EncoderDecoder().EncodeResponse(w, statusCode, data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
