package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAppError(t *testing.T) {
	appErr := NewAppError("code", "message", "target")

	assert.NotNil(t, appErr)
	assert.Equal(t, "code", appErr.ErrorCode())
	assert.Equal(t, "message", appErr.Message())
	assert.Equal(t, []error(nil), appErr.Details())
	assert.Equal(t, "target", appErr.Target())
	assert.Equal(t, "message", appErr.Error())
}

func TestNewNotFoundError(t *testing.T) {
	appErr := NewNotFoundError("not found")

	assert.NotNil(t, appErr)
	assert.Equal(t, ErrCodeNotFound, appErr.ErrorCode())
	assert.Equal(t, "not found", appErr.Message())
	assert.Equal(t, "", appErr.Target())
	assert.Equal(t, "not found", appErr.Error())
}

func TestNewRequiredFieldError(t *testing.T) {
	appErr := NewRequiredFieldError("field")

	assert.NotNil(t, appErr)
	assert.Equal(t, ErrCodeRequireField, appErr.ErrorCode())
	assert.Equal(t, "'field' is a required field", appErr.Message())
	assert.Equal(t, "field", appErr.Target())
	assert.Equal(t, "'field' is a required field", appErr.Error())
}
