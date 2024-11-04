package service

import (
	"context"
	"net/http"

	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging"
)

var _ IService = (*Service)(nil)

// Service struct
type Service struct {
	ctx            context.Context
	name           string
	url            string
	encoderDecoder encoding.EncoderDecoder

	// These values are applied by WithOptions
	database    *database.Database
	logger      *logging.Logger
	middlewares []func(http.Handler) http.Handler
	mux         *http.ServeMux
}

// IService interface
type IService interface {
	Name() string
	WithOptions(opts ...Option) *Service

	EncoderDecoder() encoding.EncoderDecoder
	Logger() *logging.Logger
	Mux() *http.ServeMux

	clone() *Service
}

// NewService creates a new Service instance with the specified name and applies any
// provided options, such as a Logger or Database, to configure the Service.
func NewService(ctx context.Context, options ...Option) *Service {
	encoderDecoder := encoding.NewEncoderDecoder()

	svc := &Service{
		ctx:            ctx,
		mux:            http.NewServeMux(),
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

// Mux returns the service http.ServeMux
func (svc *Service) Mux() *http.ServeMux {
	return svc.mux
}

// Name returns the service name
func (svc *Service) Name() string {
	return svc.name
}

// URL returns the service url
func (svc *Service) URL() string {
	return svc.url
}
