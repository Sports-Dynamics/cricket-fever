package teams

import (
	"context"
	"fmt"

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

	teamID, err := t.teamRepo.GetIDFromUUID(ctx, uuid)
	if err != nil {
		fmt.Println(" failed to get teamID from UUID : err : ", err.Error())
		return nil, err
	}

	newTeam, err := t.teamRepo.Get(ctx, *teamID)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *team) Update(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error) {
	fmt.Println(" team name : ", req.TeamName)
	newTeam, err := t.teamRepo.Update(ctx, req)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}

func (t *team) Delete(ctx context.Context, uuid string) (*models.CricketTeam, error) {

	// TO-DO
	newTeam, err := t.teamRepo.Delete(ctx, uuid)
	if err != nil {
		return newTeam, err
	}

	return newTeam, nil

}
