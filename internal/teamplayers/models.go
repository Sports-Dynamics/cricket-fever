package teamplayers

import (
	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type AddPlayerRequestParams struct {
	models.TeamPlayer
}

func (t *AddPlayerRequestParams) Validate() error {
	return nil
}

type RemovePlayerRequestParams struct {
	TeamID   int `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`
	PlayerID int `boil:"player_id" json:"player_id" toml:"player_id" yaml:"player_id"`
}

func (t *RemovePlayerRequestParams) Validate() error {
	return nil
}

type ListPlayersRequestParams struct {
	TeamID int `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`
}

func (lp *ListPlayersRequestParams) Validate() error {
	return nil
}

type FindTeamRequestParams struct {
	PlayerID int `boil:"player_id" json:"player_id" toml:"player_id" yaml:"player_id"`
}

func (ft *FindTeamRequestParams) Validate() error {
	return nil
}
