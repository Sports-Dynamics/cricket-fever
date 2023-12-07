package players

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/stretchr/testify/assert"
)

func Test_createTeam_ServeHTTP(t *testing.T) {

	createNewPlayer := func(ctx context.Context, req CreateTeamRequestParams) (models.CricketPlayer, error) {
		return models.CricketPlayer{
			PlayerID:   1,
			PlayerName: "goldy",
		}, nil
	}

	handler := NewCreateNewPlayerHandler(CreateNewPlayerStub{CreateNewPlayerFunc: createNewPlayer})

	requestBody := CreateTeamRequestParams{Name: "goldy", Country: "india"}

	requestJSON, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/team", bytes.NewReader(requestJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var responseAccount models.CricketTeam
	err = json.Unmarshal(rr.Body.Bytes(), &responseAccount)
	assert.NoError(t, err)

	assert.Equal(t, "goldy", responseAccount.TeamName)

}
