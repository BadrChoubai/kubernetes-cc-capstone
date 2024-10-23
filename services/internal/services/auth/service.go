package services

import "github.com/badrchoubai/services/internal/service"

func NewAuthService(name string, opts ...service.Option) *service.Service {
	return service.NewService(name, opts...)
}
