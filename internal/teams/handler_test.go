package teams

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

func Test_createTeam_ServeHTTP(t *testing.T) {

	createTeam := func(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error) {
		return models.CricketTeam{
			TeamID:      1,
			TeamName:    req.Name,
			TeamUUID:    uuid.New().String(),
			TeamCountry: req.Country,
			TeamState:   "karnataka",
			TeamCity:    "bangalore",
			TeamCoachID: null.IntFrom(1),
			CaptainID:   null.IntFrom(2),
		}, nil
	}

	handler := NewCreateTeamHandler(CreateTeamStub{CreateTeamFunc: createTeam})

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
