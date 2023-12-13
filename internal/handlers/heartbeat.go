package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"go.uber.org/zap"
)

// HeartbeatHandler is used to validate the liveness/readiness of this pod
func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK) // headers must be sent before output, or output sends a default 200 OK header
	if _, err := w.Write([]byte(http.StatusText(http.StatusOK))); err != nil {
		modo.Logger(r.Context()).Error("Error writing response.", zap.Error(err))
	}
}

func YourHandlerFunction(w http.ResponseWriter, r *http.Request) {
	teamID := chi.URLParam(r, "teamID")
	// Now teamID contains the value from the URL parameter
	// Do something with teamID
	fmt.Println(" team id is ", teamID)
}
