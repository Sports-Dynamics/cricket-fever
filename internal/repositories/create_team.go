package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/sports-dynamics/cricket-fever/internal/requests"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CreateTeamRepo interface {
	Create(ctx context.Context, req requests.CreateTeamRequestParams) (models.CricketTeam, error)
	// AddPlayers(ctx context.Context)
}

type CreateTeam struct {
	db *sql.DB
}

func NewCreateTeamRepo(db *sql.DB) CreateTeamRepo {
	return &CreateTeam{db: db}
}

func (ct CreateTeam) Create(ctx context.Context, req requests.CreateTeamRequestParams) (models.CricketTeam, error) {

	new_team := models.CricketTeam{
		TeamName: req.Name,
		Country:  req.Country,
	}

	if err := new_team.Insert(context.Background(), ct.db, boil.Infer()); err != nil {
		fmt.Println(" failed to insert into db : ", err.Error())
		return new_team, err
	}

	fmt.Println(" new team value : ", new_team)

	return new_team, nil
}
