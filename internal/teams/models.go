package teams

import "github.com/sports-dynamics/cricket-fever/internal/db/models"

type CreateTeamRequestParams struct {
	models.CricketTeam
}

func (t *CreateTeamRequestParams) Validate() error {
	return nil
}

const (
	TeamID string = "teamID"
)
