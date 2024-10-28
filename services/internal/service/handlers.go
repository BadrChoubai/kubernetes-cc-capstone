package service

import (
	"net/http"
)

type (
	IndexResponse struct {
		Service string `json:"service"`
	}
)

func (svc *Service) Index() http.Handler {
	response := IndexResponse{Service: svc.Name()}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			svc.respond(w, http.StatusOK, response)
		})
}

// respond handles encoding the response and managing any encoding errors.
func (svc *Service) respond(w http.ResponseWriter, statusCode int, data any) {
	if err := svc.encoderDecoder.EncodeResponse(w, statusCode, data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
