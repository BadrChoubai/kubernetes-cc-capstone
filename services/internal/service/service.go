package service

import (
	"github.com/badrchoubai/services/internal/observability/logging"
	"net/http"
	"sync"
)

var _ IService = (*Service)(nil)

type Service struct {
	name         string
	serviceMutex *sync.Mutex
	serviceMux   *http.ServeMux
	logger       *logging.Logger
}

// IService interface
type IService interface {
	WithOptions(opts ...Option) *Service
	Mux() *http.ServeMux
	Name() string
	RegisterRoutes(handlers []Handler)

	clone() *Service
}

type Handler struct {
	Path    string
	Handler http.HandlerFunc
}

func NewService(name string, opts ...Option) *Service {
	svc := &Service{
		name:       name,
		serviceMux: http.NewServeMux(),
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

func (svc *Service) Mux() *http.ServeMux {
	return svc.serviceMux
}

func (svc *Service) Name() string {
	return svc.name
}

func (svc *Service) RegisterRoutes(handlers []Handler) {
	for _, handler := range handlers {
		svc.serviceMux.Handle(
			handler.Path,
			handler.Handler,
		)
	}
}
