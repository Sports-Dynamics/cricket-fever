package players

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PlayerRepo interface {
	Create(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error)
	Get(ctx context.Context, id int) (*models.CricketPlayer, error)
	Update(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error)
	Delete(ctx context.Context, id int) (*models.CricketPlayer, error)
}

type Repo struct {
	db *sql.DB
}

func NewPlayerRepo(db *sql.DB) PlayerRepo {
	return &Repo{db: db}
}

func (ct Repo) Create(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error) {

	fmt.Println("dumping request value : ", req)

	new_team := models.CricketPlayer{
		PlayerName:        req.PlayerName,
		PlayerEmail:       req.PlayerEmail,
		PlayerMobile:      req.PlayerMobile,
		PlayerPicture:     null.BytesFrom(req.PlayerPicture.Bytes),
		PlayerRole:        req.PlayerRole,
		BattingPositions:  req.BattingPositions,
		BowlerTypes:       req.BowlerTypes,
		FieldingPositions: req.FieldingPositions,
	}

	if err := new_team.Insert(context.Background(), ct.db, boil.Infer()); err != nil {
		fmt.Println("[players][repo] failed to insert into db : ", err.Error())
		return &new_team, err
	}

	fmt.Println(" new team value : ", new_team)

	return &new_team, nil
}

func (ct Repo) Get(ctx context.Context, playerID int) (*models.CricketPlayer, error) {

	query := models.CricketPlayers(models.CricketTeamWhere.TeamID.EQ(playerID))

	player, err := query.One(ctx, ct.db)
	if err != nil {
		fmt.Println("failed to fetch particular team from team table : ", err.Error())
		return nil, err
	}

	return player, nil
}

func (ct Repo) Update(ctx context.Context, req *PlayerRequestParams) (*models.CricketPlayer, error) {

	player := models.CricketPlayer{PlayerID: req.PlayerID}

	player.Update(ctx, ct.db, boil.Infer())

	return &player, nil
}

func (ct Repo) Delete(ctx context.Context, playerID int) (*models.CricketPlayer, error) {

	player := models.CricketPlayer{PlayerID: playerID}

	player.Delete(ctx, ct.db)

	return &player, nil
}
