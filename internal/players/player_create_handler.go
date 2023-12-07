package players

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sports-dynamics/cricket-fever/internal/handlers"
	"github.com/sports-dynamics/cricket-fever/internal/modo"
	"go.uber.org/zap"
)

type createPlayer struct {
	service PlayerService
}

func CreatePlayerHandler(service PlayerService) http.HandlerFunc {

	return createPlayer{service: service}.ServeHTTP
}

type CreateNewPlayer struct {
}

func (t createPlayer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request *PlayerRequestParams

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		handlers.RespondWithError(r.Context(), w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println(" dumping value at handler level : ", request)

	if err := request.Validate(); err != nil {
		handlers.RespondWithError(r.Context(), w, err, http.StatusUnprocessableEntity)
		return
	}

	team, err := t.service.Create(context.Background(), request)

	if err != nil {
		modo.Logger(r.Context()).Error("Could not create account.", zap.Error(err))
		handlers.RespondWithError(r.Context(), w, err, http.StatusInternalServerError)
		return
	}

	handlers.RespondWithSuccess(r.Context(), w, team, http.StatusCreated)
}
