package utils

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/sports-dynamics/cricket-fever/internal/validation"
)

func GetUUIDFromRequest(r *http.Request, urlParam string) (string, error) {
	paramUUID := chi.URLParam(r, urlParam)
	if _, err := uuid.Parse(paramUUID); err != nil {
		return "", validation.NewAppError(validation.ErrCodeInvalidInput, "invalid uuid, should be a valid uuid", urlParam)
	}
	return paramUUID, nil
}
