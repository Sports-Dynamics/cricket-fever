package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("successfull"))

	})

	validApiToken := "sovApiToken"
	r := chi.NewRouter()
	authMiddlewareHandler := AuthMiddleware(validApiToken)
	r.Use(authMiddlewareHandler)

	t.Run("Valid Token", func(tt *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-API-TOKEN", validApiToken)
		recorder := httptest.NewRecorder()
		r.Get("/test", mockHandler)
		r.ServeHTTP(recorder, req)
		assert.Equal(tt, http.StatusOK, recorder.Code)
		assert.Equal(tt, "successfull", recorder.Body.String())
	})

	t.Run("Invalid Token", func(tt *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-API-TOKEN", "invalidToken")
		recorder := httptest.NewRecorder()
		r.Get("/test", mockHandler)
		r.ServeHTTP(recorder, req)
		assert.Equal(tt, http.StatusUnauthorized, recorder.Code)
	})

	t.Run("EmptyToken", func(tt *testing.T) {
		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("X-API-TOKEN", "")
		recorder := httptest.NewRecorder()
		r.Get("/test", mockHandler)
		r.ServeHTTP(recorder, req)
		assert.Equal(tt, http.StatusUnauthorized, recorder.Code)
	})
}
