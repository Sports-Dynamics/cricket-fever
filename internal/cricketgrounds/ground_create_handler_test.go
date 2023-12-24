package cricketgrounds

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
	type fields struct {
		service TeamService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	requestBody := CreateNewGroundRequest{models.CricketTeam{TeamName: "goldy", TeamCountry: "india"}}

	requestJSON, _ := json.Marshal(requestBody)

	requestRecorder := httptest.NewRecorder()

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "create new team",
			fields: fields{service: CreateCricketGroundStub{CreateFunc: func(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketTeam, error) {
				return &models.CricketTeam{
					TeamID:      1,
					TeamName:    "golden eagles",
					TeamCountry: "india",
					TeamState:   "karnataka",
					TeamCity:    "bangalore",
				}, nil
			}}},
			args: args{w: requestRecorder, r: httptest.NewRequest("POST", "/v1/teams", bytes.NewReader(requestJSON))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := createTeam{
				service: tt.fields.service,
			}
			tr.ServeHTTP(tt.args.w, tt.args.r)

			if status := requestRecorder.Result().StatusCode; status != http.StatusCreated {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			assert.Contains(t, requestRecorder.Body.String(), "golden eagles")
		})
	}
}
