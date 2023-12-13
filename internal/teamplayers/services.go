package teamplayers

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type TeamPlayersService interface {

	// join an existing cricket team
	Join(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error)

	// removing a player from the team
	Remove(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error)

	// find all players belonging to a particular team
	GetPlayers(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error)

	// find all teams a player belongs to
	FindTeams(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error)
}

type teamplayers struct {
	teamRepo TeamPlayersRepo
}

func NewTeamPlayersService(teams TeamPlayersRepo) TeamPlayersService {
	return &teamplayers{teamRepo: teams}
}

func (t *teamplayers) Join(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error) {

	result, err := t.teamRepo.Join(ctx, req)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (t *teamplayers) Remove(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error) {

	result, err := t.teamRepo.Remove(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *teamplayers) GetPlayers(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error) {

	result, err := t.teamRepo.GetPlayers(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *teamplayers) FindTeams(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error) {

	result, err := t.teamRepo.FindTeams(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
