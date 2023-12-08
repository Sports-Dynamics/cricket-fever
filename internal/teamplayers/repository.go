package teamplayers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TeamPlayersRepo interface {
	Join(ctx context.Context, teamUUID, playerUUID string, joiningDate int) (*models.TeamPlayer, error)
	Remove(ctx context.Context, teamUUID, playerID string) (*models.TeamPlayer, error)
	GetPlayers(ctx context.Context, teamUUID string) (*[]models.CricketPlayer, error)
	FindTeams(ctx context.Context, playerUUID string) (*[]models.CricketTeam, error)
}

type Repo struct {
	db *sql.DB
}

func NewTeamRepo(db *sql.DB) TeamPlayersRepo {
	return &Repo{db: db}
}

func (ct Repo) Join(ctx context.Context, teamUUID, playerUUID string, joiningDate int) (*models.TeamPlayer, error) {

	new_team := models.TeamPlayer{}

	if err := new_team.Insert(context.Background(), ct.db, boil.Infer()); err != nil {
		fmt.Println(" failed to insert into db : ", err.Error())
		return &new_team, err
	}

	fmt.Println(" new team value : ", new_team)

	return &new_team, nil
}

func (ct Repo) Remove(ctx context.Context, teamUUID, playerID string) (*models.TeamPlayer, error) {

	return nil, nil
}
func (ct Repo) GetPlayers(ctx context.Context, teamUUID string) (*[]models.CricketPlayer, error) {
	return nil, nil
}
func (ct Repo) FindTeams(ctx context.Context, playerUUID string) (*[]models.CricketTeam, error) {
	return nil, nil
}
