package service

import (
	"net/http"

	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging"
)

var _ IService = (*Service)(nil)

// Service struct
type Service struct {
	name           string
	url            string
	encoderDecoder encoding.EncoderDecoder

	// These values are applied by WithOptions
	database *database.Database
	logger   *logging.Logger
	mux      *http.ServeMux
}

// IService interface
type IService interface {
	Name() string
	WithOptions(opts ...Option) *Service

	clone() *Service
}
