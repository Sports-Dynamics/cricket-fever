package players

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type CreatePlayerService interface {
	CreateNewPlayer(ctx context.Context, req NewPlayerRequestParams) (models.CricketPlayer, error)
}

type newPlayer struct {
	newPlayerRepo CreatePlayerRepo
}

func NewPlayerService(repo CreatePlayerRepo) CreatePlayerService {
	return &newPlayer{newPlayerRepo: repo}
}

func (t *newPlayer) CreateNewPlayer(ctx context.Context, req NewPlayerRequestParams) (models.CricketPlayer, error) {

	newTeam, err := t.newPlayerRepo.Create(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *newPlayer) JoinTeam(ctx context.Context, teamID, playerID int) {

}
