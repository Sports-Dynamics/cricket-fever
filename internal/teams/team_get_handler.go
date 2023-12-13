package teams

import (
	"context"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"github.com/sports-dynamics/cricket-fever/internal/utils"
	"go.uber.org/zap"
)

type getTeam struct {
	repo TeamRepo
}

func GetTeamHandler(repo TeamRepo) http.HandlerFunc {
	return getTeam{repo: repo}.ServeHTTP
}

func (t getTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	uuid, err := utils.GetUUIDFromRequest(r, TeamUUID)
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
