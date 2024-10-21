package services

import (
	"net/http"
	"sync"

	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging/zap"
)

type Service struct {
	Name           string
	Logger         logging.Logger
	ServiceMutex   sync.Mutex
	EncoderDecoder *encoding.ServerEncoderDecoder
}

// ServiceInterface interface
type ServiceInterface interface {
	RegisterRouter(router *http.ServeMux)
}
