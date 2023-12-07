package teams

import (
	"context"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"github.com/sports-dynamics/cricket-fever/internal/utils"
	"go.uber.org/zap"
)

type deleteTeam struct {
	service TeamService
}

func DeleteTeamHandler(service TeamService) http.HandlerFunc {

	return createTeam{service: service}.ServeHTTP
}

func (t deleteTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	uuid, err := utils.GetUUIDFromRequest(r, TeamID)
	if err != nil {
		handlers.RespondWithError(r.Context(), w, err, http.StatusBadRequest)
		return
	}

	team, err := t.service.Delete(context.Background(), uuid)

	if err != nil {
		modo.Logger(r.Context()).Error("Could not create account.", zap.Error(err))
		handlers.RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	handlers.RespondWithSuccess(r.Context(), w, team, http.StatusCreated)
}
