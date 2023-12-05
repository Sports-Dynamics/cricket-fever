package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithSuccess(t *testing.T) {
	ctx := context.Background()
	responseData := map[string]string{"message": "success"}
	statusCode := http.StatusOK
	recorder := httptest.NewRecorder()

	respondWithSuccess(ctx, recorder, responseData, statusCode)

	assert.Equal(t, statusCode, recorder.Code, "Status code should match expected")
	var responseBody map[string]string
	err := json.NewDecoder(recorder.Body).Decode(&responseBody)
	assert.NoError(t, err, "Decoding response body should not error")
	assert.Equal(t, responseData, responseBody, "Response body should match expected")
}
