package teamplayers

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type TeamPlayersService interface {

	// join an existing cricket team
	Join(ctx context.Context, teamUUID, playerUUID string, joiningDate int) (*models.TeamPlayer, error)

	// removing a player from the team
	Remove(ctx context.Context, teamUUID, playerID string) (*models.TeamPlayer, error)

	// find all players belonging to a particular team
	GetPlayers(ctx context.Context, teamUUID string) (*[]models.CricketPlayer, error)

	// find all teams a player belongs to
	FindTeams(ctx context.Context, playerUUID string) (*[]models.CricketTeam, error)
}

type teamplayers struct {
	teamRepo TeamPlayersRepo
}

func NewTeamService(teams TeamPlayersRepo) TeamPlayersService {
	return &teamplayers{teamRepo: teams}
}

func (t *teamplayers) Join(ctx context.Context, teamUUID, playerUUID string, joiningDate int) (*models.TeamPlayer, error) {

	newTeam, err := t.teamRepo.Join(ctx, 1, 2, joiningDate)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil
}

func (t *teamplayers) Remove(ctx context.Context, teamUUID, playerID string) (*models.TeamPlayer, error) {
	return nil, nil
}

func (t *teamplayers) GetPlayers(ctx context.Context, teamUUID string) (*[]models.CricketPlayer, error) {
	return nil, nil
}

func (t *teamplayers) FindTeams(ctx context.Context, playerUUID string) (*[]models.CricketTeam, error) {
	return nil, nil
}
