package validation

import (
	"fmt"
	"strings"
)

type MultiError struct {
	Errors []error
	Reason string
	Code   string
}

var _ AppError = (*MultiError)(nil)

func (e MultiError) ErrorCode() string {
	return e.Code
}

func (e MultiError) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e MultiError) Message() string {
	return e.Reason
}

func (e MultiError) Details() []error {
	return e.Errors
}

func (e MultiError) Target() string {
	return ""
}

func (e MultiError) Error() string {
	errorMessages := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		errorMessages[i] = err.Error()
	}
	allErrors := strings.Join(errorMessages, "; ")

	if e.Message() != "" {
		return fmt.Sprintf("%s: %s", e.Reason, allErrors)
	}

	return allErrors
}
