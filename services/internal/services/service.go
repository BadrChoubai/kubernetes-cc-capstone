package services

import "net/http"

// Service interface
type Service interface {
	RegisterRouter(router *http.ServeMux)
}
