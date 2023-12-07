package players

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type PlayerService interface {
	Create(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error)
	Get(ctx context.Context, playerUUID string) (*models.CricketPlayer, error)
	Update(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error)
	Delete(ctx context.Context, playerUUID string) (*models.CricketPlayer, error)
}

type newPlayer struct {
	repo PlayerRepo
}

func NewPlayerService(repo PlayerRepo) PlayerService {
	return &newPlayer{repo: repo}
}

func (t *newPlayer) Create(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error) {

	newPlayer, err := t.repo.Create(ctx, req)
	if err != nil {
		return newPlayer, err
	}

	return newPlayer, nil

}
func (t *newPlayer) Get(ctx context.Context, playerUUID string) (*models.CricketPlayer, error) {

	// TO-DO fetch UUID for the player
	newTeam, err := t.repo.Get(ctx, 1)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *newPlayer) Update(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error) {

	newTeam, err := t.repo.Update(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *newPlayer) Delete(ctx context.Context, playerUUID string) (*models.CricketPlayer, error) {

	// TO-DO fetch UUID for the player
	newTeam, err := t.repo.Delete(ctx, 1)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}
