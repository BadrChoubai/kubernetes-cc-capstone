package auth

import (
	"fmt"
	"github.com/badrchoubai/services/internal/services"
	"github.com/badrchoubai/services/internal/services/users"
	"github.com/badrchoubai/services/internal/validator"
	"net/http"
)

var _ services.IService = (*Service)(nil)

// RegisterRouter register service specific routes
func (s *Service) RegisterRouter(router *http.ServeMux) {
	router.Handle("POST /auth/v1/generateToken", s.GenerateTokenHandler())
}

// GenerateTokenHandler handle request for authorization token generation
func (s *Service) GenerateTokenHandler() http.HandlerFunc {
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

		err := s.EncoderDecoder.Decode(r, input)
		if err != nil {
			s.handleError(w, http.StatusBadRequest, "decoding request body", err)
			return
		}

		v := validator.NewValidator()
		users.ValidateEmail(v, input.Email)
		users.ValidatePasswordPlaintext(v, input.Password)

		if !v.Valid() {
			s.handleValidationErrors(w, http.StatusUnprocessableEntity, "decoding request body", v.Errors)
			return
		}

		var token = &token{}
		token.Token = "token"

		err = s.EncoderDecoder.Encode(w, http.StatusOK, token.Token)
		if err != nil {
			s.handleError(w, http.StatusBadRequest, "encoding response body", err)
			return
		}
	}
}

func (s *Service) handleError(w http.ResponseWriter, status int, whatWasHappening string, error error) {
	type ErrorResponse struct {
		Error []string `json:"errors"`
		Count int      `json:"count"`
	}

	errResponse := &ErrorResponse{
		Error: []string{error.Error()},
		Count: 1,
	}

	s.Logger.Error(whatWasHappening, error)
	encodeErr := s.EncoderDecoder.Encode(w, status, errResponse)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Service) handleValidationErrors(w http.ResponseWriter, status int, whatWasHappening string, errors map[string]string) {
	type ErrorsResponse struct {
		Errors map[string]string `json:"errors"`
		Count  int               `json:"count"`
	}

	parsedErrors := map[string]string{}
	for k, ev := range errors {
		s.Logger.Error(whatWasHappening, fmt.Errorf("%s", k))
		parsedErrors[k] = ev
	}

	errResponse := &ErrorsResponse{
		Errors: parsedErrors,
		Count:  len(errors),
	}

	encodeErr := s.EncoderDecoder.Encode(w, status, errResponse)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
