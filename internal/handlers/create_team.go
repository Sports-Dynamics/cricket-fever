package handlers

import (
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/services"
)

type createTeam struct {
	service services.Team
}

func NewCreateTeamHandler(service services.Team) http.HandlerFunc {
	return createTeam{service: service}.ServeHTTP
}

type CreateTeam struct {
}

func (h createTeam) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// var request CreateAccountRequest

	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&request); err != nil {
	// 	RespondWithError(r.Context(), w, err, http.StatusBadRequest)
	// 	return
	// }
	// defer r.Body.Close()

	// if err := request.Validate(); err != nil {
	// 	RespondWithError(r.Context(), w, err, http.StatusUnprocessableEntity)
	// 	return
	// }

	// account, err := h.sov.CreateAccount(r.Context(), request.NewAccount)
	// if err != nil {
	// 	modo.Logger(r.Context()).Error("Could not create account.", zap.Error(err))
	// 	RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
	// 	return
	// }

	// respondWithSuccess(r.Context(), w, account, http.StatusCreated)
}
