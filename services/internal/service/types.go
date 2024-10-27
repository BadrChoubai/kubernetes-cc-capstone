package service

import (
	"github.com/badrchoubai/services/internal/observability/logging"
	"net/http"
)

var _ IService = (*Service)(nil)

type Handler struct {
	Path    string
	Handler http.Handler
}

type Service struct {
	logger  *logging.Logger
	mux     *http.ServeMux
	handler http.Handler
	name    string

	url string
}

// IService interface
type IService interface {
	Name() string
	Handler() http.Handler
	WithOptions(opts ...Option) *Service

	clone() *Service
}
