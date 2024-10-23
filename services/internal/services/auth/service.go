package services

import "github.com/badrchoubai/services/internal/service"

var _ IAuthService = (*AuthService)(nil)

type AuthService struct {
	service *service.Service
}

type IAuthService interface {
	Service() *service.Service
}

func NewAuthService(name string, opts ...service.Option) *AuthService {
	svc := service.NewService(name, opts...)

	authSvc := &AuthService{
		service: svc,
	}

	return authSvc
}

func (as *AuthService) Service() *service.Service {
	return as.service
}
