package cricketgrounds

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TeamRepo interface {
	Create(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error)
	Get(ctx context.Context, id int) (*models.CricketGround, error)
	GetByUUID(ctx context.Context, uuid string) (*models.CricketGround, error)
	Update(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error)
	Delete(ctx context.Context, uuid string) (*models.CricketGround, error)
	GetIDFromUUID(ctx context.Context, uuid string) (*int, error)
}

type Repo struct {
	db *sql.DB
}

func NewTeamRepo(db *sql.DB) TeamRepo {
	return &Repo{db: db}
}

func (ct Repo) Create(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error) {

	new_team := models.CricketGround{
		GroundName:  req.GroundName,
		TeamCountry: req.TeamCountry,
		TeamState:   req.TeamState,
		TeamCity:    req.TeamCity,
	}

	if err := new_team.Insert(context.Background(), ct.db, boil.Infer()); err != nil {
		fmt.Println(" failed to insert into db : ", err.Error())
		return &new_team, err
	}

	fmt.Println(" new team value : ", new_team)

	return &new_team, nil
}

func (ct Repo) Get(ctx context.Context, teamID int) (*models.CricketGround, error) {

	query := models.CricketGrounds(models.CricketTeamWhere.TeamID.EQ(teamID))

	team, err := query.One(ctx, ct.db)
	if err != nil {
		fmt.Println("failed to fetch particular team from team table : ", err.Error())
		return nil, err
	}

	fmt.Println(" new team value : ", team)

	return team, nil
}

func (ct Repo) GetByUUID(ctx context.Context, teamUUID string) (*models.CricketGround, error) {

	query := models.CricketGrounds(models.CricketTeamWhere.TeamUUID.EQ(teamUUID))

	team, err := query.One(ctx, ct.db)
	if err != nil {
		fmt.Println("failed to fetch particular team from team table : ", err.Error())
		return nil, err
	}

	fmt.Println(" new team value : ", team)

	return team, nil
}

func (ct Repo) Update(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error) {

	_, err := req.Update(ctx, ct.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	fmt.Println(" new team value : ", req)

	return &req.CricketGround, nil
}

func (ct Repo) Delete(ctx context.Context, uuid string) (*models.CricketGround, error) {

	query := models.CricketGrounds(models.CricketTeamWhere.TeamUUID.EQ(uuid))

	team, err := query.One(ctx, ct.db)
	if err != nil {
		return nil, err
	}

	if _, err := team.Delete(ctx, ct.db); err != nil {
		return nil, err
	}

	return team, nil
}

func (ct Repo) GetIDFromUUID(ctx context.Context, uuid string) (*int, error) {

	query := models.CricketGrounds(models.CricketTeamWhere.TeamUUID.EQ(uuid))

	team, err := query.One(ctx, ct.db)
	if err != nil {
		fmt.Println("failed to fetch particular team from team table : ", err.Error())
		return nil, err
	}
	return &team.GroundID, nil
}
