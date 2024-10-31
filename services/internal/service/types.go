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
	database *database.Database
	logger   *logging.Logger
	mux      *Mux
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

type Mux struct {
	mux *http.ServeMux
}

type IMux interface {
	Handle(pattern string, handler http.Handler)
	Routes() map[string]http.Handler
	ServeMux() *http.ServeMux
}
