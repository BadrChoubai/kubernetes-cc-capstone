package service

import (
	"github.com/badrchoubai/services/internal/observability/logging"
	"net/http"
	"sync"
)

var _ IService = (*Service)(nil)

type Service struct {
	handler      http.Handler
	logger       *logging.Logger
	mux          *http.ServeMux
	name         string
	serviceMutex *sync.Mutex
}

// IService interface
type IService interface {
	Name() string
	WithOptions(opts ...Option) *Service

	clone() *Service
}
