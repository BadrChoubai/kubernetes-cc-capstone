package users

import (
	"github.com/badrchoubai/services/internal/services"
	"net/http"
)

var _ services.ServiceInterface = (*Service)(nil)

func (a *Service) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("/register", a.RegisterUserHandler)
	router.HandleFunc("/login", a.LoginUserHandler)
}

func (a *Service) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (a *Service) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
