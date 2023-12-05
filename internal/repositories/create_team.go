package repositories

import "context"

type Team interface {
	Create(ctx context.Context)
	AddPlayers(ctx context.Context)
}

type TeamRepo struct {
}
