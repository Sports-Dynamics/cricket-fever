package teams

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type TeamService interface {
	CreateTeam(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error)
	// JoinTeam(ctx context.Context, teamId int, playerID int)
}

type newTeam struct {
	teamRepo CreateTeamRepo
}

func NewTeamService(teams CreateTeamRepo) TeamService {
	return &newTeam{teamRepo: teams}
}

func (t *newTeam) CreateTeam(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error) {

	newTeam, err := t.teamRepo.Create(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *newTeam) JoinTeam(ctx context.Context, teamID, playerID int) {

}
