package teams

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CreateTeamRepo interface {
	Create(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error)
}

type Repo struct {
	db *sql.DB
}

func NewTeamRepo(db *sql.DB) CreateTeamRepo {
	return &Repo{db: db}
}

func (ct Repo) Create(ctx context.Context, req CreateTeamRequestParams) (models.CricketTeam, error) {

	new_team := models.CricketTeam{
		TeamUUID:      uuid.New().String(),
		TeamName:      req.Name,
		TeamCountry:   req.Country,
		TeamState:     "karnataka",
		TeamCity:      "bangalore",
		TeamCoachID:   null.IntFrom(2),
		CaptainID:     null.IntFrom(3),
		ViceCaptainID: null.IntFrom(4),
	}

	if err := new_team.Insert(context.Background(), ct.db, boil.Infer()); err != nil {
		fmt.Println(" failed to insert into db : ", err.Error())
		return new_team, err
	}

	fmt.Println(" new team value : ", new_team)

	return new_team, nil
}
