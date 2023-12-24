package cricketgrounds

import (
	"context"

	"github.com/sports-dynamics/cricket-fever/internal/db/models"
)

type CreateCricketGroundStub struct {
	CreateFunc func(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error)
	GetFunc    func(ctx context.Context, uuid string) (*models.CricketGround, error)
	UpdateFunc func(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error)
	DeleteFunc func(ctx context.Context, uuid string) (*models.CricketGround, error)
}

var _ TeamService = (*CreateCricketGroundStub)(nil)

func (cg CreateCricketGroundStub) Create(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error) {
	if cg.CreateFunc != nil {
		return cg.CreateFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}
func (s CreateCricketGroundStub) Get(ctx context.Context, uuid string) (*models.CricketGround, error) {
	if s.CreateFunc != nil {
		return s.GetFunc(ctx, uuid)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s CreateCricketGroundStub) Update(ctx context.Context, req *CreateNewGroundRequest) (*models.CricketGround, error) {
	if s.CreateFunc != nil {
		return s.UpdateFunc(ctx, req)
	}
	panic("Unexpected call to CreateTeam stub")
}

func (s CreateCricketGroundStub) Delete(ctx context.Context, uuid string) (*models.CricketGround, error) {
	if s.CreateFunc != nil {
		return s.DeleteFunc(ctx, uuid)
	}
	panic("Unexpected call to CreateTeam stub")
}
