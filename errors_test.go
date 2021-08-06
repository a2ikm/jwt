package jwt_test

import (
	"errors"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

func TestValidationErrorUnwrap(t *testing.T) {
	original := errors.New("inner error")

	err := &jwt.ValidationError{Inner: original, Errors: jwt.ValidationErrorUnverifiable}
	unwrapped := errors.Unwrap(err)

	if unwrapped != original {
		t.Errorf("Unwrap() returned unexpected result: %v", unwrapped)
	}
}
