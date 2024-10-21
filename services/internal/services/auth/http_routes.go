package auth

import (
	"github.com/badrchoubai/services/internal/services"
	"net/http"
)

var _ services.ServiceInterface = (*Service)(nil)

func (a *Service) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("/auth/generateToken", a.GenerateTokenHandler)
}

func (a *Service) GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
