package users

import (
	"net/http"

	"github.com/badrchoubai/services/internal/services"
)

var _ services.IService = (*Service)(nil)

// RegisterRouter register service specific routes
func (u *Service) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /users/v1/register", u.RegisterUserHandler)
	router.HandleFunc("POST /users/v1/login", u.LoginUserHandler)
}

// RegisterUserHandler handle request user registration
func (u *Service) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

// LoginUserHandler handle request for user login
func (u *Service) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
