package params

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	"go.uber.org/dig"
)

type SagaBusParams struct {
	dig.In

	SagaHandlers []bus.SagaHandler `group:"sagaBus"`
}
