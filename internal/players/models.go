package players

import (
	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type PlayerRequestParams struct {
	models.CricketPlayer
}

func (t *PlayerRequestParams) Validate() error {
	return nil
}

const (
	PlayerUUID string = "PlayerUUID"
)
