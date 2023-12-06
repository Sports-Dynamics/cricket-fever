package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

func RespondWithSuccess(ctx context.Context, w http.ResponseWriter, response any, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		zap.Error(errors.New("encoding failed , err = " + err.Error()))
	}

}
