package teams

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type CreateTeamStub struct {
	CreateFunc func(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error)
	GetFunc    func(ctx context.Context, uuid string) (*models.CricketTeam, error)
	UpdateFunc func(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error)
	DeleteFunc func(ctx context.Context, uuid string) (*models.CricketTeam, error)
}

var _ TeamService = (*CreateTeamStub)(nil)

func (s CreateTeamStub) Create(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error) {
	if s.CreateFunc != nil {
		return s.CreateFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}
func (s CreateTeamStub) Get(ctx context.Context, uuid string) (*models.CricketTeam, error) {
	if s.CreateFunc != nil {
		return s.GetFunc(ctx, uuid)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s CreateTeamStub) Update(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error) {
	if s.CreateFunc != nil {
		return s.UpdateFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s CreateTeamStub) Delete(ctx context.Context, uuid string) (*models.CricketTeam, error) {
	if s.CreateFunc != nil {
		return s.DeleteFunc(ctx, uuid)
	}
	panic("Unexpected call to CreateTeam stub")
}
