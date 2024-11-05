package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging"
)

var _ IService = (*Service)(nil)

// Service struct
type Service struct {
	ctx            context.Context
	name           string
	path           string
	encoderDecoder encoding.EncoderDecoder

	// These values are applied by WithOptions
	database    *database.Database
	logger      *logging.Logger
	middlewares []func(http.Handler) http.Handler
	mux         *http.ServeMux
}

var (
	InvalidNameError  = errors.New("invalid service name format, expected <resource>-service-v<version>")
	NameCheckingError = errors.New("error checking name format")
)

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
func NewService(ctx context.Context, name string, options ...Option) (*Service, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	path, err := makePathFromName(name)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		ctx:            ctx,
		encoderDecoder: encoding.NewEncoderDecoder(),
		mux:            http.NewServeMux(),
		name:           name,
		path:           path,
	}
	svc.WithOptions(options...)

	return svc, nil
}

func validateName(name string) error {
	var (
		pattern = `^[a-z]+-service-v[0-9]+$`
	)

	matched, err := regexp.MatchString(pattern, name)
	if err != nil {
		return NameCheckingError
	}

	if !matched {
		return InvalidNameError
	}

	return nil
}

func makePathFromName(name string) (string, error) {
	if err := validateName(name); err != nil {
		return "", err
	}

	path := "/api/%s/%s"
	parts := strings.Split(name, "-")
	path = fmt.Sprintf(path, parts[2], parts[0])

	return path, nil
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

// Path returns the service url
func (svc *Service) Path() string {
	return svc.path
}
