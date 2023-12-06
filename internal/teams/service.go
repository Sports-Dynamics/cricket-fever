package teams

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type Team interface {
	CreateTeam(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error)
	// JoinTeam(ctx context.Context, teamId int, playerID int)
}

type TeamService struct {
	teamRepo CreateTeamRepo
}

func NewTeamService(teams CreateTeamRepo) Team {
	return &TeamService{teamRepo: teams}
}

func (t *TeamService) CreateTeam(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error) {

	newTeam, err := t.teamRepo.Create(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *TeamService) JoinTeam(ctx context.Context, teamID, playerID int) {

}
