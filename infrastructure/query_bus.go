package infrastructure

import (
	"github.com/VulpesFerrilata/catan-service/infrastructure/bus"
	"github.com/VulpesFerrilata/catan-service/infrastructure/dig/params"
	"github.com/VulpesFerrilata/catan-service/infrastructure/middlewares"
)

func NewQueryBus(params params.QueryBusParams,
	validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware) bus.QueryBus {
	queryBus := bus.NewQueryBus()

	queryBus.Register(params.QueryHandlers...)
	queryBus.Use(
		validationMiddleware.WrapHandler,
		transactionMiddleware.WrapHandler,
	)

	return queryBus
}
