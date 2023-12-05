package services

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/repositories"
)

type Team interface {
	CreateTeam(ctx context.Context)
	JoinTeam(ctx context.Context, teamId int, playerID int)
}

type TeamService struct {
	teamRepo repositories.TeamRepo
}

func NewTeamService(teams repositories.TeamRepo) Team {
	return &TeamService{teamRepo: teams}
}

func (t *TeamService) CreateTeam(ctx context.Context) {

}

func (t *TeamService) JoinTeam(ctx context.Context, teamID, playerID int) {

}
