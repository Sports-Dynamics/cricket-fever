package players

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type PlayerStubs struct {
	CreatFunc  func(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error)
	GetFunc    func(ctx context.Context, playerUUID string) (*models.CricketPlayer, error)
	UpdateFunc func(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error)
	DeleteFunc func(ctx context.Context, playerUUID string) (*models.CricketPlayer, error)
}

var _ PlayerService = (*PlayerStubs)(nil)

func (s PlayerStubs) Create(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error) {
	if s.CreatFunc != nil {
		return s.CreatFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s PlayerStubs) Get(ctx context.Context, playerUUID string) (*models.CricketPlayer, error) {
	if s.GetFunc != nil {
		return s.GetFunc(ctx, playerUUID)
	}
	panic("Unexpected call to CreateTeam stub")
}
func (s PlayerStubs) Update(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error) {
	if s.UpdateFunc != nil {
		return s.UpdateFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s PlayerStubs) Delete(ctx context.Context, playerUUID string) (*models.CricketPlayer, error) {
	if s.DeleteFunc != nil {
		return s.DeleteFunc(ctx, playerUUID)
	}
	panic("Unexpected call to CreateTeam stub")
}
