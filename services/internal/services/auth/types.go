package services

import (
	"time"

	"github.com/badrchoubai/services/internal/validator"
)

type (
	Token struct {
		Plaintext string    `json:"plaintext"`
		Hash      []byte    `json:"-"`
		UserID    int64     `json:"-"`
		Expiry    time.Time `json:"expiry"`
		Scope     string    `json:"-"`
	}

	IToken interface {
		New(userID int64, ttl time.Duration, scope string) (*Token, error)
		Insert(token *Token) error
		DeleteAllForUser(scope string, userID int64) error
	}
)

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}
