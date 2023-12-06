package services

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/sports-dynamics/cricket-fever/internal/repositories"
	"github.com/sports-dynamics/cricket-fever/internal/requests"
)

type Team interface {
	CreateTeam(ctx context.Context, req requests.CreateTeamRequestParams) (models.CricketTeam, error)
	// JoinTeam(ctx context.Context, teamId int, playerID int)
}

type TeamService struct {
	teamRepo repositories.CreateTeamRepo
}

func NewTeamService(teams repositories.CreateTeamRepo) Team {
	return &TeamService{teamRepo: teams}
}

func (t *TeamService) CreateTeam(ctx context.Context, req requests.CreateTeamRequestParams) (models.CricketTeam, error) {

	newTeam, err := t.teamRepo.Create(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *TeamService) JoinTeam(ctx context.Context, teamID, playerID int) {

}
