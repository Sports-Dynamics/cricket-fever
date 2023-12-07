package players

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CreatePlayerRepo interface {
	Create(ctx context.Context, req NewPlayerRequestParams) (models.CricketPlayer, error)
}

type Repo struct {
	db *sql.DB
}

func NewPlayerRepo(db *sql.DB) CreatePlayerRepo {
	return &Repo{db: db}
}

func (ct Repo) Create(ctx context.Context, req NewPlayerRequestParams) (models.CricketPlayer, error) {

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
		return new_team, err
	}

	fmt.Println(" new team value : ", new_team)

	return new_team, nil
}
