package teams

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"go.uber.org/zap"
)

type createTeam struct {
	service Team
}

func NewCreateTeamHandler(service Team) http.HandlerFunc {

	return createTeam{service: service}.ServeHTTP
}

type CreateTeam struct {
}

func (t createTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request CreateTeamRequestParams

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

	team, err := t.service.CreateTeam(context.Background(), request)

	if err != nil {
		modo.Logger(r.Context()).Error("Could not create account.", zap.Error(err))
		handlers.RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	handlers.RespondWithSuccess(r.Context(), w, team, http.StatusCreated)
}
