package auth

import (
	"net/http"

	"github.com/badrchoubai/services/internal/services"
)

var _ services.ServiceInterface = (*Service)(nil)

// RegisterRouter godoc
//
//	@title			Auth API
//	@version		1.0
//	@description	This authorization API is used to generate tokens for a given user
//	@host			0.0.0.0:8080
//	@BasePath		/auth/v1
func (a *Service) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /auth/v1/generateToken", a.GenerateTokenHandler)
}

// GenerateTokenHandler godoc
//
//	@Summary	Generate token
//	@Accept		json
//	@Produce	json
//	@Success	200	{string}	string
//	@Router		/auth/v1/generateToken [post]
func (a *Service) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
