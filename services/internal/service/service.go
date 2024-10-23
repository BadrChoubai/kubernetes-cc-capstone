package service

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
)

var _ IService = (*Service)(nil)

type Service struct {
	Name           string
	ServiceMutex   *sync.Mutex
	Logger         *logging.Logger
	EncoderDecoder *encoding.ServerEncoderDecoder
	DB             *database.Database
}

// IService interface
type IService interface {
	WithOptions(opts ...Option) *Service
}

func NewService(name string, opts ...Option) *Service {
	svc := &Service{
		Name:         name,
		ServiceMutex: &sync.Mutex{},
	}

	return svc.WithOptions(opts...)
}

// WithOptions clones the current Service, applies the supplied Options, and
// returns the resulting Service. It's safe to use concurrently.
func (svc *Service) WithOptions(opts ...Option) *Service {
	s := svc.clone()
	for _, opt := range opts {
		opt.apply(s)
	}
	return s
}

func (svc *Service) clone() *Service {
	clone := *svc
	return &clone
}

func (svc *Service) HandleError(w http.ResponseWriter, status int, whatWasHappening string, error error) {
	type ErrorResponse struct {
		Error []string `json:"errors"`
		Count int      `json:"count"`
	}

	errResponse := &ErrorResponse{
		Error: []string{error.Error()},
		Count: 1,
	}

	svc.Logger.Error(whatWasHappening, error)
	encodeErr := svc.EncoderDecoder.Encode(w, status, errResponse)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (svc *Service) HandleValidationErrors(w http.ResponseWriter, status int, whatWasHappening string, errors map[string]string) {
	type ErrorsResponse struct {
		Errors map[string]string `json:"errors"`
		Count  int               `json:"count"`
	}

	parsedErrors := map[string]string{}
	for k, ev := range errors {
		svc.Logger.Error(whatWasHappening, fmt.Errorf("%s", k))
		parsedErrors[k] = ev
	}

	errResponse := &ErrorsResponse{
		Errors: parsedErrors,
		Count:  len(errors),
	}

	encodeErr := svc.EncoderDecoder.Encode(w, status, errResponse)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
