package events

import (
	"context"

	"github.com/vulpes-ferrilata/catan-service/domain/events"
)

func NewGameDetailProjector() *GameDetailProjector {
	return &GameDetailProjector{}
}

type GameDetailProjector struct {
}

func (g GameDetailProjector) Handle(ctx context.Context, gameUpdatedEvent *events.GameUpdatedEvent) error {
	return nil
}
