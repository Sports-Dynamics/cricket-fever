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

	createPlayer := func(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error) {
		return &models.CricketPlayer{
			PlayerID:   1,
			PlayerName: "goldy",
		}, nil
	}

	handler := CreatePlayerHandler(PlayerStubs{CreatFunc: createPlayer})

	requestBody := PlayerRequestParams{models.CricketPlayer{PlayerName: "goldy", PlayerMobile: 123456789, PlayerEmail: "abc@gmail.com"}}

	requestJSON, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/player", bytes.NewReader(requestJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var response models.CricketPlayer
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "goldy", response.PlayerName)

}
