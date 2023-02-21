package infrastructure

import (
	"github.com/vulpes-ferrilata/catan-service/application/queries"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewQueryBus(validationMiddleware *middlewares.ValidationMiddleware,
	findGamePaginationByLimitByOffsetQueryHandler *queries.FindGamePaginationByLimitByOffsetQueryHandler,
	getGameDetailByIDByUserIDQueryHandler *queries.GetGameDetailByIDByUserIDQueryHandler) (*cqrs.QueryBus, error) {
	queryBus := &cqrs.QueryBus{}

	queryBus.Use(
		validationMiddleware.QueryHandlerMiddleware(),
	)

	queryBus.Register(&queries.FindGamePaginationByLimitByOffsetQuery{}, cqrs.WrapQueryHandlerFunc(findGamePaginationByLimitByOffsetQueryHandler.Handle))
	queryBus.Register(&queries.GetGameDetailByIDByUserIDQuery{}, cqrs.WrapQueryHandlerFunc(getGameDetailByIDByUserIDQueryHandler.Handle))

	return queryBus, nil
}
