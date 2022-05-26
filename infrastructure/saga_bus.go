package infrastructure

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	"github.com/VulpesFerrilata/catan-service/infrastructure/dig/params"
)

func NewSagaBus(params params.SagaBusParams) bus.SagaBus {
	sagaBus := bus.NewSagaBus()

	sagaBus.Register(params.SagaHandlers...)

	return sagaBus
}
