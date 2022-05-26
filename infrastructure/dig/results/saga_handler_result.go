package results

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	"go.uber.org/dig"
)

type SagaHandlerResult struct {
	dig.Out

	SagaHandler bus.SagaHandler `group:"sagaBus"`
}
