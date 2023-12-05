package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeartbeatHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/heartbeat", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HeartbeatHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")
	expected := http.StatusText(http.StatusOK)
	assert.Equal(t, expected, rr.Body.String(), "handler returned unexpected body")
}

func TestHeartbeatHandler_MethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/heartbeat", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HeartbeatHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code, "handler returned wrong status code")
}
