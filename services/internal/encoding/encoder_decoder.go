package encoding

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/badrchoubai/services/internal/observability/logging/zap"
)

type (
	EncoderDecoder interface {
		EncodeResponse(w http.ResponseWriter, status int, v any) error
		DecodeRequest(r *http.Request) []byte
	}

	// ServerEncoderDecoder is our concrete implementation of EncoderDecoder.
	ServerEncoderDecoder struct {
		logger *logging.Logger
	}

	encoder interface {
		Encode(w http.ResponseWriter, status int, v any) error
	}

	decoder interface {
		Decode(r *http.Request) (any, error)
	}
)

func NewEncoderDecoder(logger *logging.Logger) *ServerEncoderDecoder {
	return &ServerEncoderDecoder{
		logger: logger,
	}
}

func (sec *ServerEncoderDecoder) Encode(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		sec.logger.Error("encoding JSON", err)
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func (sec *ServerEncoderDecoder) Decode(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(&dest); err != nil {
		sec.logger.Error("decoding JSON", err)
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}
