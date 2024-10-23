package services

import (
	"github.com/badrchoubai/services/internal/services/users"
	"github.com/badrchoubai/services/internal/validator"
	"net/http"
)

func (as *AuthService) RegisterRouter(mux *http.ServeMux) {
	mux.Handle("POST /auth/v1/generateToken", as.GenerateTokenHandler())
}

// GenerateTokenHandler handle request for authorization token generation
func (as *AuthService) GenerateTokenHandler() http.HandlerFunc {
	type input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type token struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var input = &input{}

		err := as.Service().EncoderDecoder.Decode(r, input)
		if err != nil {
			as.Service().HandleError(w, http.StatusBadRequest, "decoding request body", err)
			return
		}

		v := validator.NewValidator()
		users.ValidateEmail(v, input.Email)
		users.ValidatePasswordPlaintext(v, input.Password)

		if !v.Valid() {
			as.Service().HandleValidationErrors(w, http.StatusUnprocessableEntity, "decoding request body", v.Errors)
			return
		}

		var token = &token{}
		token.Token = "token"

		err = as.Service().EncoderDecoder.Encode(w, http.StatusOK, token.Token)
		if err != nil {
			as.Service().HandleError(w, http.StatusBadRequest, "encoding response body", err)
			return
		}
	}
}
