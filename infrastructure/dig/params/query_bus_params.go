package params

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	"go.uber.org/dig"
)

type QueryBusParams struct {
	dig.In

	QueryHandlers []bus.QueryHandler `group:"queryBus"`
}
