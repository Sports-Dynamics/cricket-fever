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

type TeamRepo interface {
	Create(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error)
	Get(ctx context.Context, id int) (*models.CricketTeam, error)
	Update(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error)
	Delete(ctx context.Context, id int) (*models.CricketTeam, error)
}

type Repo struct {
	db *sql.DB
}

func NewTeamRepo(db *sql.DB) TeamRepo {
	return &Repo{db: db}
}

func (ct Repo) Create(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error) {

	new_team := models.CricketTeam{
		TeamUUID:      uuid.New().String(),
		TeamName:      req.TeamName,
		TeamCountry:   req.TeamCountry,
		TeamState:     "karnataka",
		TeamCity:      "bangalore",
		TeamCoachID:   null.IntFrom(2),
		CaptainID:     null.IntFrom(3),
		ViceCaptainID: null.IntFrom(4),
	}

	if err := new_team.Insert(context.Background(), ct.db, boil.Infer()); err != nil {
		fmt.Println(" failed to insert into db : ", err.Error())
		return &new_team, err
	}

	fmt.Println(" new team value : ", new_team)

	return &new_team, nil
}

func (ct Repo) Get(ctx context.Context, teamID int) (*models.CricketTeam, error) {

	query := models.CricketTeams(models.CricketTeamWhere.TeamID.EQ(teamID))

	team, err := query.One(ctx, ct.db)
	if err != nil {
		fmt.Println("failed to fetch particular team from team table : ", err.Error())
		return nil, err
	}

	fmt.Println(" new team value : ", team)

	return team, nil
}

func (ct Repo) Update(ctx context.Context, req *CreateTeamRequestParams) (*models.CricketTeam, error) {

	team := models.CricketTeam{TeamID: req.TeamID}

	team.Update(ctx, ct.db, boil.Infer())

	fmt.Println(" new team value : ", team)

	return &team, nil
}

func (ct Repo) Delete(ctx context.Context, id int) (*models.CricketTeam, error) {

	team := models.CricketTeam{TeamID: id}

	team.Delete(ctx, ct.db)

	fmt.Println(" new team value : ", team)

	return &team, nil
}
