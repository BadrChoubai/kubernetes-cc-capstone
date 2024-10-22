package users

import (
	"github.com/badrchoubai/services/internal/validator"
	"regexp"
	"time"
)

type (
	password struct {
		plaintext *string
		hash      []byte
	}

	User struct {
		ID        int64     `json:"id"`
		Activated bool      `json:"activated"`
		CreatedAt time.Time `json:"created_at"`
		Email     string    `json:"email"`
		Name      string    `json:"name"`
		Password  password  `json:"-"`
		Version   int       `json:"-"`
	}

	IUser interface {
		Insert(user *User) error
		GetForToken(tokenScope, tokenPlaintext string) (*User, error)
		GetByEmail(email string) (*User, error)
		Update(user *User) error
	}
)

func ValidateEmail(v *validator.Validator, email string) {
	var (
		EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	)

	v.Check(email != "", "email", "must be provided")
	v.Check(v.Matches(email, EmailRX), "email", "must be valid")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 characters long")
	v.Check(len(password) <= 72, "password", "must not exceed max length (72 characters)")
}
