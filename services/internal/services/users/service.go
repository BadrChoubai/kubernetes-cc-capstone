package users

import (
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/observability/logging"
	"sync"

	"github.com/badrchoubai/services/internal/services"
)

var _ services.Service = (*UserService)(nil)

// UserService implements Service
type UserService struct {
	ServiceMutex   sync.Mutex
	encoderDecoder *encoding.ServerEncoderDecoder
}

func NewUsersService(logger logging.Logger) services.Service {
	ser := &UserService{
		ServiceMutex:   sync.Mutex{},
		encoderDecoder: encoding.NewEncoderDecoder(logger),
	}

	return ser
}
