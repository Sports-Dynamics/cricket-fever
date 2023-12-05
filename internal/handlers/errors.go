package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/validation"

	"go.uber.org/zap"
)

type ErrorResponse struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Target  string        `json:"target,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Target  string `json:"target,omitempty"`
	Message string `json:"message"`
}

func RespondWithError(ctx context.Context, w http.ResponseWriter, err error, fallbackStatusCode int) {
	var errResponse ErrorResponse
	var appError validation.AppError
	if errors.As(err, &appError) {
		errResponse.Code = appError.ErrorCode()
		errResponse.Message = appError.Message()
		errResponse.Target = appError.Target()
		errResponse.Details = makeErrorDetails(appError.Details())
	} else {
		errResponse.Code = validation.ErrCodeFallback
		errResponse.Message = err.Error()

	}
	w.WriteHeader(mapErrorToHttpStatus(err, fallbackStatusCode))
	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		zap.Error(errors.New("Could not encode error response" + err.Error()))
	}
}

func makeErrorDetails(errorArray []error) []ErrorDetail {
	var details []ErrorDetail
	for _, err := range errorArray {
		var appError validation.AppError
		var detail ErrorDetail
		if errors.As(err, &appError) {
			detail.Code = appError.ErrorCode()
			detail.Target = appError.Target()
			detail.Message = appError.Message()
		} else {
			detail.Code = validation.ErrCodeFallback
			detail.Message = err.Error()
		}
		details = append(details, detail)
	}
	return details
}

func mapErrorToHttpStatus(err error, fallbackCode int) int {
	var appError validation.AppError
	if errors.As(err, &appError) {
		if statusCode, ok := errorCodeToStatusCode[appError.ErrorCode()]; ok {
			return statusCode
		}
	}

	return fallbackCode
}

var errorCodeToStatusCode = map[string]int{
	validation.ErrCodeNotFound:          http.StatusNotFound,
	validation.ErrCodeInvalidInput:      http.StatusBadRequest,
	validation.ErrCodeRequireField:      http.StatusBadRequest,
	validation.ErrCodeDuplicate:         http.StatusConflict,
	validation.ErrCodeForbidden:         http.StatusForbidden,
	validation.ErrCodeInsufficientFunds: http.StatusUnprocessableEntity,
}
