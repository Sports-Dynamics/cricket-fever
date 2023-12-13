package teamplayers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"go.uber.org/zap"
)

type joinTeam struct {
	service TeamPlayersService
}

func JoinTeamHandler(service TeamPlayersService) http.HandlerFunc {

	return joinTeam{service: service}.ServeHTTP
}

func (t joinTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	team, err := t.service.Join(context.Background(), request)

	if err != nil {
		modo.Logger(r.Context()).Error("Could not create account.", zap.Error(err))
		handlers.RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	handlers.RespondWithSuccess(r.Context(), w, team, http.StatusCreated)
}
