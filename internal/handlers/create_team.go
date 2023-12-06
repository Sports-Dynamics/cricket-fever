package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"github.com/sports-dynamics/cricket-fever/internal/requests"
	"github.com/sports-dynamics/cricket-fever/internal/services"
	"go.uber.org/zap"
)

type createTeam struct {
	service services.Team
}

func NewCreateTeamHandler(service services.Team) http.HandlerFunc {
	return createTeam{service: service}.ServeHTTP
}

type CreateTeam struct {
}

func (t createTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateTeamRequestParams

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		RespondWithError(r.Context(), w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := request.Validate(); err != nil {
		RespondWithError(r.Context(), w, err, http.StatusUnprocessableEntity)
		return
	}

	team, err := t.service.CreateTeam(context.Background(), request)

	if err != nil {
		modo.Logger(r.Context()).Error("Could not create account.", zap.Error(err))
		RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	respondWithSuccess(r.Context(), w, team, http.StatusCreated)
}
