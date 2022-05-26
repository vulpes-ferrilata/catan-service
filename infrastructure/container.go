package infrastructure

import (
	command_handlers "github.com/VulpesFerrilata/catan-service/application/commands/handlers"
	query_handlers "github.com/VulpesFerrilata/catan-service/application/queries/handlers"
	"github.com/VulpesFerrilata/catan-service/domain/mappers"
	"github.com/VulpesFerrilata/catan-service/domain/services"
	"github.com/VulpesFerrilata/catan-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/catan-service/infrastructure/persistence/repositories"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//3rd party libraries
	container.Provide(NewConfig)
	container.Provide(NewGorm)
	container.Provide(NewValidate)
	container.Provide(NewCommandBus)
	container.Provide(NewQueryBus)
	container.Provide(NewSagaBus)
	container.Provide(NewUniversalTranslator)

	//middlewares
	container.Provide(middlewares.NewErrorHandlerMiddleware)
	container.Provide(middlewares.NewValidationMiddleware)
	container.Provide(middlewares.NewTransactionMiddleware)
	container.Provide(middlewares.NewAuthenticationMiddleware)
	container.Provide(middlewares.NewTranslatorMiddleware)

	//Persistence layer
	container.Provide(repositories.NewGameRepository)
	container.Provide(repositories.NewPlayerRepository)
	container.Provide(repositories.NewDiceRepository)

	//Domain layer
	//--Services
	container.Provide(services.NewAuthenticationService)
	container.Provide(services.NewGameService)
	container.Provide(services.NewPlayerService)
	container.Provide(services.NewDiceService)
	//--Mappers
	container.Provide(mappers.NewGameMapper)
	container.Provide(mappers.NewPlayerMapper)
	container.Provide(mappers.NewDiceMapper)

	//Application layer
	container.Provide(query_handlers.NewGetClaimByAccessTokenQueryHandler)
	container.Provide(command_handlers.NewCreateGameCommandHandler)
	container.Provide(command_handlers.NewJoinGameCommandHandler)
	container.Provide(command_handlers.NewStartGameCommandHandler)

	return container
}
