package auth

import "time"

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
