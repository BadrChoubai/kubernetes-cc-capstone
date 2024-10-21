package users

import (
	"net/http"

	"github.com/badrchoubai/services/internal/services"
)

var _ services.ServiceInterface = (*Service)(nil)

// RegisterRouter godoc
//
//	@title			User API
//	@version		1.0
//	@description	This user API is used to register and login users
//	@host			0.0.0.0:8080
//	@BasePath		/users/v1
func (a *Service) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /users/v1/register", a.RegisterUserHandler)
	router.HandleFunc("POST /users/v1/login", a.LoginUserHandler)
}

// RegisterUserHandler godoc
//
//	@Summary	register user
//	@Accept		json
//	@Produce	json
//	@Success	200	{string}	string
//	@Router		/register [post]
func (a *Service) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

// LoginUserHandler godoc
//
//	@Summary	login user
//	@Accept		json
//	@Produce	json
//	@Success	200	{string}	string
//	@Router		/login [post]
func (a *Service) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
