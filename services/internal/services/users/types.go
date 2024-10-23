package users

import (
	"fmt"
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
		emailrx = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	)

	v.Check(email != "", "email", "must be provided")
	v.Check(v.Matches(email, emailrx), "email", "must be valid")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	var (
		minLength = 12
		maxLength = 128

		minLengthMessage = fmt.Sprintf("must be at least %d characters", minLength)
		maxLengthMessage = fmt.Sprintf("must be at least %d characters", maxLength)

		// Check for at least one uppercase letter
		simplePasswordValidations = map[string]bool{
			"hasUpper":   regexp.MustCompile(`[A-Z]`).MatchString(password),
			"hasLower":   regexp.MustCompile(`[a-z]`).MatchString(password),
			"hasSpecial": regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(password),
			"hasDigit":   regexp.MustCompile(`[0-9]`).MatchString(password),
		}
		simplePasswordValidationMessages = map[string]string{
			"hasUpper":   "password must have at least one uppercase character",
			"hasLower":   "password must have at least one lowercase character",
			"hasSpecial": "password must have at least one special character",
			"hasDigit":   "password must have at least one digit",
		}
	)

	v.Check(password != "", "password", "must be provided")
	for check, ok := range simplePasswordValidations {
		v.Check(ok, check, simplePasswordValidationMessages[check])
	}
	v.Check(len(password) >= minLength, "password", minLengthMessage)
	v.Check(len(password) <= maxLength, "password", maxLengthMessage)
}
