package teams

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type TeamService interface {
	Create(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error)
	Get(ctx context.Context, teamUUID string) (*models.CricketTeam, error)
	Update(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error)
	Delete(ctx context.Context, teamUUID string) (*models.CricketTeam, error)
}

type team struct {
	teamRepo TeamRepo
}

func NewTeamService(teams TeamRepo) TeamService {
	return &team{teamRepo: teams}
}

func (t *team) Create(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error) {

	newTeam, err := t.teamRepo.Create(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *team) Get(ctx context.Context, uuid string) (*models.CricketTeam, error) {

	// TO-DO
	newTeam, err := t.teamRepo.Get(ctx, 1)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *team) Update(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error) {

	// TO-DO
	newTeam, err := t.teamRepo.Update(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *team) Delete(ctx context.Context, uuid string) (*models.CricketTeam, error) {

	// TO-DO
	newTeam, err := t.teamRepo.Delete(ctx, 1)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}
