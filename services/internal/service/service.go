package service

import (
	"net/http"

	"github.com/badrchoubai/services/internal/encoding"
)

// NewService creates a new Service instance with the specified name and applies any
// provided options, such as a Logger or Database, to configure the Service.
func NewService(name string, options ...Option) *Service {
	mux := http.NewServeMux()
	encoderDecoder := encoding.NewEncoderDecoder()
	svc := &Service{
		name:           name,
		mux:            mux,
		encoderDecoder: encoderDecoder,
	}

	return svc.WithOptions(options...)
}

// RegisterRoute adds a new http.Handler to the Service's mux for the specified path.
// If the provided path is empty, it defaults to the root path ("/").
func (svc *Service) RegisterRoute(path string, handler http.Handler) {
	if path == "" {
		path = "/"
	}

	svc.mux.Handle(svc.URL()+path, handler)
}

// Name returns the service name
func (svc *Service) Name() string {
	return svc.name
}

// URL returns the service url
func (svc *Service) URL() string {
	return svc.url
}

// Mux returns the service HTTP Multiplexer
func (svc *Service) Mux() *http.ServeMux {
	return svc.mux
}
