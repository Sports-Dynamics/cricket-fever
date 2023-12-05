package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithError_AppError(t *testing.T) {
	recorder := httptest.NewRecorder()

	appErr := mockAppError{
		code:    "CustomErrorCode",
		message: "CustomErrorMessage",
		target:  "CustomTarget",
		details: []error{errors.New("DetailError")},
	}

	RespondWithError(nil, recorder, appErr, http.StatusInternalServerError)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	var response ErrorResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "CustomErrorCode", response.Code)
	assert.Equal(t, "CustomErrorMessage", response.Message)
	assert.Equal(t, "CustomTarget", response.Target)
	assert.Len(t, response.Details, 1)
	assert.Equal(t, "ErrCodeFallback", response.Details[0].Code)
	assert.Equal(t, "DetailError", response.Details[0].Message)
}

func TestRespondWithError_GenericError(t *testing.T) {
	recorder := httptest.NewRecorder()

	genericErr := errors.New("GenericErrorMessage")

	RespondWithError(nil, recorder, genericErr, http.StatusInternalServerError)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	var response ErrorResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "ErrCodeFallback", response.Code)
	assert.Equal(t, "GenericErrorMessage", response.Message)
	assert.Empty(t, response.Target)
	assert.Empty(t, response.Details)
}

type mockAppError struct {
	code    string
	message string
	target  string
	details []error
}

func (e mockAppError) ErrorCode() string {
	return e.code
}

func (e mockAppError) Message() string {
	return e.message
}

func (e mockAppError) Details() []error {
	return e.details
}

func (e mockAppError) Target() string {
	return e.target
}

func (e mockAppError) Error() string {
	return e.message
}
