package teamplayers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"go.uber.org/zap"
)

type getTeam struct {
	repo TeamPlayersRepo
}

func GetTeamHandler(repo TeamPlayersRepo) http.HandlerFunc {

	return getTeam{repo: repo}.ServeHTTP
}

func (t getTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var request *AddPlayerRequestParams

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		handlers.RespondWithError(r.Context(), w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := request.Validate(); err != nil {
		handlers.RespondWithError(r.Context(), w, err, http.StatusUnprocessableEntity)
		return
	}

	players, err := t.repo.GetPlayers(context.Background(), request)

	if err != nil {
		modo.Logger(r.Context()).Error("Could not get team information.", zap.Error(err))
		handlers.RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	handlers.RespondWithSuccess(r.Context(), w, players, http.StatusOK)
}
