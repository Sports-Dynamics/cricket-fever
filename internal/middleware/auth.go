package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"github.com/sports-dynamics/cricket-fever/internal/validation"
	"go.uber.org/zap"
)

const (
	xApiToken = "X-API-TOKEN"
)

func AuthMiddleware(validApiToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sovApiToken := r.Header.Get(xApiToken)
			if err := validateApiTokens(sovApiToken, validApiToken); err != nil {
				handlers.RespondWithError(r.Context(), w, err, http.StatusUnauthorized)
				return
			}

			// Create a response writer wrapper to capture the status code
			rww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			// Call the next handler
			next.ServeHTTP(rww, r)

			// Log the response
			modo.Logger(context.Background()).Info("auth middleware", zap.Int("response status", rww.Status()))
		})
	}
}

func validateApiTokens(apiToken, validApiToken string) error {
	if len(apiToken) == 0 {
		return validation.NewAppError(validation.ErrCodeAuthFailed, "unauthorized access , empty api token", xApiToken)
	}
	if apiToken != validApiToken {
		return validation.NewAppError(validation.ErrCodeAuthFailed, "unauthorized access , invalid api token", xApiToken)
	}

	return nil
}
