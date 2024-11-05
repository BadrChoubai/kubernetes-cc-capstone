/*
Package service provides an abstraction for managing microservices within the application.

It defines the Service struct, which encapsulates the core functionality for handling
service-specific logic, including HTTP request routing, encoding and decoding of
messages, and middleware management. The IService interface outlines the methods
that must be implemented by any service, ensuring a consistent contract for
service behavior.

Key features of the package include:

  - Creation of service instances through the NewService function, which supports
    flexible configuration via options.
  - Name validation for services to enforce a specific naming convention.
  - Support for encoding/decoding messages.

This package serves as a foundational component for building RESTful APIs and microservices
within the broader application architecture.
*/
package service

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strings"

	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
)

var _ IService = (*Service)(nil)

// Service struct
type Service struct {
	ctx            context.Context
	name           string
	path           string
	encoderDecoder encoding.EncoderDecoder

	// These values are applied by WithOptions
	database *database.Database
	logger   *zap.Logger
	mux      *http.ServeMux
}

var (
	errInvalidName  = errors.New("invalid service name format, expected <resource>-service-v<version>")
	errCheckingName = errors.New("error checking name format")
)

// IService interface
type IService interface {
	Name() string
	WithOptions(opts ...Option) *Service

	EncoderDecoder() encoding.EncoderDecoder
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
		return errCheckingName
	}

	if !matched {
		return errInvalidName
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

// EncoderDecoder returns the service encoding.EncoderDecoder
func (svc *Service) EncoderDecoder() encoding.EncoderDecoder {
	return svc.encoderDecoder
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
