package services

import (
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
	"net/http"
	"sync"

	"github.com/badrchoubai/services/internal/encoding"
)

type Service struct {
	Name           string
	ServiceMutex   *sync.Mutex
	Logger         *logging.Logger
	EncoderDecoder *encoding.ServerEncoderDecoder
}

// IService interface
type IService interface {
	RegisterRouter(router *http.ServeMux)
}
