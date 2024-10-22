package validator

import (
	"regexp"
)

var _ IValidator = (*Validator)(nil)

type (
	Validator struct {
		Errors map[string]string
	}

	IValidator interface {
		Valid() bool
		AddError(key, message string)
		Check(ok bool, key, message string)
		Matches(value string, rx *regexp.Regexp) bool
	}
)

func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}

	return false
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)
	for _, v := range values {
		uniqueValues[v] = true
	}

	return len(values) == len(uniqueValues)
}
