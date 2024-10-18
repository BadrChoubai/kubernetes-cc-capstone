package auth

import (
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging"
	"sync"

	"github.com/badrchoubai/services/internal/services"
)

var _ services.Service = (*AuthService)(nil)

// AuthService implements Service
type AuthService struct {
	ServiceMutex   sync.Mutex
	encoderDecoder *encoding.ServerEncoderDecoder
}

func NewAuthService(logger logging.Logger) services.Service {
	ser := &AuthService{
		ServiceMutex:   sync.Mutex{},
		encoderDecoder: encoding.NewEncoderDecoder(logger),
	}

	return ser
}
