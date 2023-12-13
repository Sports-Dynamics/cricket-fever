package teamplayers

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type CreateTeamStub struct {
	JoinFunc       func(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error)
	RemoveFunc     func(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error)
	GetPlayersFunc func(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error)
	FindTeamsFunc  func(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error)
}

var _ TeamPlayersService = (*CreateTeamStub)(nil)

func (s CreateTeamStub) Join(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error) {
	if s.JoinFunc != nil {
		return s.JoinFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}
func (s CreateTeamStub) Remove(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error) {
	if s.RemoveFunc != nil {
		return s.RemoveFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s CreateTeamStub) GetPlayers(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error) {
	if s.GetPlayersFunc != nil {
		return s.GetPlayersFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s CreateTeamStub) FindTeams(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error) {
	if s.FindTeamsFunc != nil {
		return s.FindTeamsFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}
