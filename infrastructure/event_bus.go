package infrastructure

import (
	event_handlers "github.com/vulpes-ferrilata/catan-service/application/events"
	"github.com/vulpes-ferrilata/catan-service/domain/events"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewEventBus(gameOverviewProjector *event_handlers.GameOverviewProjector,
	gameDetailProjector *event_handlers.GameDetailProjector) (*cqrs.EventBus, error) {
	queryBus := &cqrs.EventBus{}

	queryBus.Register(&events.GameUpdatedEvent{}, cqrs.WrapEventHandlerFunc(gameOverviewProjector.Handle))
	queryBus.Register(&events.GameUpdatedEvent{}, cqrs.WrapEventHandlerFunc(gameDetailProjector.Handle))

	return queryBus, nil
}
