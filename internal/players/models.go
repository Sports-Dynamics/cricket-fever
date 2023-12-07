package players

import (
	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type NewPlayerRequestParams struct {
	models.CricketPlayer
}

func (t *NewPlayerRequestParams) Validate() error {
	return nil
}
