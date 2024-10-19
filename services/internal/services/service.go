package services

import (
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging/zap"
	"net/http"
	"sync"
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
