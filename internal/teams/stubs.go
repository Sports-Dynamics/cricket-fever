package teams

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type CreateTeamStub struct {
	CreateTeamFunc func(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error)
}

var _ Team = (*CreateTeamStub)(nil)

func (s CreateTeamStub) CreateTeam(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error) {
	if s.CreateTeamFunc != nil {
		return s.CreateTeamFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}
