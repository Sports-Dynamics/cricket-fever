package cricketgrounds

import (
	"context"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"github.com/sports-dynamics/cricket-fever/internal/utils"
	"go.uber.org/zap"
)

type getGrounds struct {
	repo TeamRepo
}

// GetGroundsHandler returns an http.HandlerFunc that handles
// the GET request for retrieving team information.
func GetGroundsHandler(repo TeamRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := getGrounds{repo: repo}
		t.ServeHTTP(w, r)
	}
}

// ServeHTTP handles the HTTP request for getting a team.
func (t *getGrounds) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uuid, err := utils.GetUUIDFromRequest(r, GroundUUID)
	if err != nil {
		handlers.RespondWithError(r.Context(), w, err, http.StatusBadRequest)
		return
	}

	team, err := t.repo.GetByUUID(context.Background(), uuid)
	if err != nil {
		modo.Logger(r.Context()).Error("Could not get team information.", zap.Error(err))
		handlers.RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	handlers.RespondWithSuccess(r.Context(), w, team, http.StatusOK)
}
