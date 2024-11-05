// Package encoding provides utilities for encoding and decoding JSON data in HTTP requests and responses
package encoding

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var _ EncoderDecoder = (*JSONEncoderDecoder)(nil)

type (
	// EncoderDecoder interface defines methods for encoding responses and decoding requests
	EncoderDecoder interface {
		EncodeResponse(w http.ResponseWriter, status int, v any) error
		DecodeRequest(r *http.Request, dest any) error
	}

	// JSONEncoderDecoder is our concrete implementation of the EncoderDecoder interface.
	JSONEncoderDecoder struct{}
)

// NewEncoderDecoder creates and returns a new instance of JSONEncoderDecoder
func NewEncoderDecoder() EncoderDecoder {
	return &JSONEncoderDecoder{}
}

// EncodeResponse handles encoding a JSON response
func (ed *JSONEncoderDecoder) EncodeResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encoding response: %w", err)
	}
	return nil
}

// DecodeRequest handles decoding a JSON request body
func (ed *JSONEncoderDecoder) DecodeRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decoding request: %w", err)
	}
	return nil
}
