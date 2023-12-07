package players

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type CreateNewPlayerStub struct {
	CreateNewPlayerFunc func(ctx context.Context, req NewPlayerRequestParams) (models.CricketPlayer, error)
}

var _ CreatePlayerService = (*CreateNewPlayerStub)(nil)

func (s CreateNewPlayerStub) CreateNewPlayer(ctx context.Context, req NewPlayerRequestParams) (models.CricketPlayer, error) {
	if s.CreateNewPlayerFunc != nil {
		return s.CreateNewPlayerFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}
