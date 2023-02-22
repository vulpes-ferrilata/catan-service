package events

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/events"
	"github.com/vulpes-ferrilata/catan-service/view/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/repositories"
)

func NewGameOverviewProjector() *GameOverviewProjector {
	return &GameOverviewProjector{}
}

type GameOverviewProjector struct {
	gameOverviewRepository repositories.GameOverviewRepository
}

func (g GameOverviewProjector) Handle(ctx context.Context, gameUpdatedEvent *events.GameUpdatedEvent) error {
	gameOverview, err := mappers.GameOverviewMapper{}.ToView(gameUpdatedEvent.Game)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := g.gameOverviewRepository.InsertOrUpdate(ctx, gameOverview); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
