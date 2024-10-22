package players

import (
	"context"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"github.com/sports-dynamics/cricket-fever/internal/utils"
	"go.uber.org/zap"
)

type deletePlayer struct {
	repo PlayerRepo
}

func DeletePlayerHandler(repo PlayerRepo) http.HandlerFunc {

	return deletePlayer{repo: repo}.ServeHTTP
}

func (t deletePlayer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	uuid, err := utils.GetUUIDFromRequest(r, PlayerUUID)
	if err != nil {
		handlers.RespondWithError(r.Context(), w, err, http.StatusBadRequest)
		return
	}

	team, err := t.repo.Delete(context.Background(), uuid)

	if err != nil {
		modo.Logger(r.Context()).Error("Could not get player information.", zap.Error(err))
		handlers.RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	handlers.RespondWithSuccess(r.Context(), w, team, http.StatusOK)
}
