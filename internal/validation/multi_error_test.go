package validation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiErrorErrorCode(t *testing.T) {
	err := MultiError{Code: "123"}
	assert.Equal(t, "123", err.ErrorCode())
}

func TestMultiErrorHasErrors(t *testing.T) {
	errWithoutErrors := MultiError{}
	assert.False(t, errWithoutErrors.HasErrors())

	errWithErrors := MultiError{Errors: []error{errors.New("error 1"), errors.New("error 2")}}
	assert.True(t, errWithErrors.HasErrors())
}

func TestMultiErrorMessage(t *testing.T) {
	err := MultiError{Reason: "Reason"}
	assert.Equal(t, "Reason", err.Message())
}

func TestMultiErrorDetails(t *testing.T) {
	err := MultiError{Errors: []error{errors.New("error 1"), errors.New("error 2")}}
	details := err.Details()
	assert.Len(t, details, 2)
}

func TestMultiErrorTarget(t *testing.T) {
	err := MultiError{}
	assert.Equal(t, "", err.Target())
}

func TestMultiErrorError(t *testing.T) {
	err := MultiError{
		Errors: []error{errors.New("error 1"), errors.New("error 2")},
		Reason: "Main error",
	}
	expectedError := "Main error: error 1; error 2"
	assert.EqualError(t, err, expectedError)

	errWithoutMessage := MultiError{
		Errors: []error{errors.New("error 1"), errors.New("error 2")},
	}
	expectedErrorWithoutMessage := "error 1; error 2"
	assert.EqualError(t, errWithoutMessage, expectedErrorWithoutMessage)
}
