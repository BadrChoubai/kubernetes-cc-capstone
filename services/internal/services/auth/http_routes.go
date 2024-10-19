package auth

import (
	"github.com/badrchoubai/services/internal/services"
	"net/http"
)

var _ services.ServiceInterface = (*AuthService)(nil)

func (a *AuthService) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("/auth/generateToken", a.GenerateTokenHandler)
}

func (a *AuthService) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
