package encoding

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var _ EncoderDecoder = (*JSONEncoderDecoder)(nil)

type (
	EncoderDecoder interface {
		EncodeResponse(w http.ResponseWriter, status int, v any) error
		DecodeRequest(r *http.Request, dest any) error
	}

	// JSONEncoderDecoder is our concrete implementation of EncoderDecoder.
	JSONEncoderDecoder struct{}
)

func NewEncoderDecoder() EncoderDecoder {
	return &JSONEncoderDecoder{}
}

func (ed *JSONEncoderDecoder) EncodeResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encoding response: %w", err)
	}
	return nil
}

func (ed *JSONEncoderDecoder) DecodeRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decoding request: %w", err)
	}
	return nil
}
