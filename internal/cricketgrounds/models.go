package cricketgrounds

import "github.com/sports-dynamics/cricket-fever/internal/db/models"

type CreateNewGroundRequest struct {
	models.CricketGround
}

// Validate checks if the CreateTeamRequestParams object is valid.
// It returns an error if the object is invalid, or nil if it is valid.
func (t *CreateNewGroundRequest) Validate() error {
	// Since there are no validation rules for CreateTeamRequestParams,
	// we can simply return nil to indicate that the object is valid.
	return nil
}

const (
	GroundUUID string = "groundUUID"
)
