package infrastructure

import (
	command_handlers "github.com/vulpes-ferrilata/catan-service/application/commands/handlers"
	query_handlers "github.com/vulpes-ferrilata/catan-service/application/queries/handlers"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/repositories"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/projectors"
	"github.com/vulpes-ferrilata/catan-service/presentation"
	"github.com/vulpes-ferrilata/catan-service/presentation/v1/servers"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//Infrastructure layer
	container.Provide(NewConfig)
	container.Provide(NewMongo)
	container.Provide(NewValidator)
	container.Provide(NewLogrus)
	container.Provide(NewUniversalTranslator)
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)

	//Domain layer
	//--Repositories
	container.Provide(repositories.NewGameRepository)

	//View layer
	//--Projectors
	container.Provide(projectors.NewGameProjector)

	//Application layer
	//--Queries
	container.Provide(query_handlers.NewFindGamesByUserIDQueryHandler)
	container.Provide(query_handlers.NewGetGameByIDByUserIDQueryHandler)
	//--Commands
	container.Provide(command_handlers.NewCreateGameCommandHandler)
	container.Provide(command_handlers.NewJoinGameCommandHandler)
	container.Provide(command_handlers.NewStartGameCommandHandler)
	container.Provide(command_handlers.NewBuildSettlementAndRoadCommandHandler)
	container.Provide(command_handlers.NewRollDicesCommandHandler)
	container.Provide(command_handlers.NewMoveRobberCommandHandler)
	container.Provide(command_handlers.NewEndTurnCommandHandler)
	container.Provide(command_handlers.NewBuildSettlementCommandHandler)
	container.Provide(command_handlers.NewBuildRoadCommandHandler)
	container.Provide(command_handlers.NewUpgradeCityCommandHandler)
	container.Provide(command_handlers.NewBuyDevelopmentCardCommandHandler)
	container.Provide(command_handlers.NewToggleResourceCardsCommandHandler)
	container.Provide(command_handlers.NewMaritimeTradeCommandHandler)
	container.Provide(command_handlers.NewOfferTradingCommandHandler)
	container.Provide(command_handlers.NewConfirmTradingCommandHandler)
	container.Provide(command_handlers.NewCancelTradingCommandHandler)
	container.Provide(command_handlers.NewPlayKnightCardCommandHandler)
	container.Provide(command_handlers.NewPlayRoadBuildingCardCommandHandler)
	container.Provide(command_handlers.NewPlayYearOfPlentyCardCommandHandler)
	container.Provide(command_handlers.NewPlayMonopolyCardCommandHandler)

	//Presentation layer
	//--Server
	container.Provide(presentation.NewServer)
	//--Controllers
	container.Provide(servers.NewCatanServer)

	return container
}
