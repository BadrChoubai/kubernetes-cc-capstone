package service

import (
	"context"
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging"
	"net/http"
)

// NewServiceMux create a new Mux instance
func NewServiceMux() *Mux {
	return &Mux{
		mux: http.NewServeMux(),
	}
}

// NewService creates a new Service instance with the specified name and applies any
// provided options, such as a Logger or Database, to configure the Service.
func NewService(ctx context.Context, options ...Option) *Service {
	encoderDecoder := encoding.NewEncoderDecoder()

	svc := &Service{
		ctx:            ctx,
		mux:            NewServiceMux(),
		encoderDecoder: encoderDecoder,
	}

	return svc.WithOptions(options...)
}

func (svc *Service) EncoderDecoder() encoding.EncoderDecoder {
	return svc.encoderDecoder
}

func (svc *Service) Logger() *logging.Logger {
	return svc.logger
}

// Mux returns the service Mux http.ServeMux
func (svc *Service) Mux() *http.ServeMux {
	return svc.mux.ServeMux()
}

// Name returns the service name
func (svc *Service) Name() string {
	return svc.name
}

// URL returns the service url
func (svc *Service) URL() string {
	return svc.url
}

func (m *Mux) ServeMux() *http.ServeMux {
	return m.mux
}
