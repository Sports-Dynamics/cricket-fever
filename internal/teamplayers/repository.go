package teamplayers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TeamPlayersRepo interface {
	Join(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error)
	Remove(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error)
	GetPlayers(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error)
	FindTeams(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error)
}

type Repo struct {
	db *sql.DB
}

func NewTeamPlayersREpo(db *sql.DB) TeamPlayersRepo {
	return &Repo{db: db}
}

func (ct Repo) Join(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error) {

	if err := req.Insert(context.Background(), ct.db, boil.Infer()); err != nil {
		fmt.Println(" failed to insert into db : ", err.Error())
		return &req.TeamPlayer, err
	}

	return &req.TeamPlayer, nil
}

func (ct Repo) Remove(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayer, error) {

	_, err := models.TeamPlayers(
		qm.Where("team_id = ?", req.TeamID),
		qm.And("player_id = ?", req.PlayerID),
	).DeleteAll(ctx, ct.db)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (ct Repo) GetPlayers(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error) {

	teamPlayers, err := models.TeamPlayers(
		qm.Where("team_id = ?", req.TeamID),
		qm.And("player_id = ?", req.PlayerID),
	).All(ctx, ct.db)
	if err != nil {
		return nil, err
	}

	return &teamPlayers, nil
}

func (ct Repo) FindTeams(ctx context.Context, req *AddPlayerRequestParams) (*models.TeamPlayerSlice, error) {

	teamPlayers, err := models.TeamPlayers(
		qm.And("player_id = ?", req.PlayerID),
	).All(ctx, ct.db)
	if err != nil {
		return nil, err
	}

	return &teamPlayers, nil
}
