package infrastructure

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	"github.com/VulpesFerrilata/catan-service/infrastructure/dig/params"
	"github.com/VulpesFerrilata/catan-service/infrastructure/middlewares"
)

func NewCommandBus(params params.CommandBusParams,
	validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware) bus.CommandBus {
	commandBus := bus.NewCommandBus()

	commandBus.Register(params.CommandHandlers...)
	commandBus.Use(
		validationMiddleware.WrapHandler,
		transactionMiddleware.WrapHandler,
	)

	return commandBus
}
