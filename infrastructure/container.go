package infrastructure

import (
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/application/queries"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/domain/mongo/repositories"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/grpc/interceptors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/projectors"
	"github.com/vulpes-ferrilata/catan-service/presentation"
	v1 "github.com/vulpes-ferrilata/catan-service/presentation/v1"
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
	container.Provide(NewQueryBus)
	container.Provide(NewCommandBus)
	//--Grpc interceptors
	container.Provide(interceptors.NewRecoverInterceptor)
	container.Provide(interceptors.NewErrorHandlerInterceptor)
	container.Provide(interceptors.NewLocaleInterceptor)
	container.Provide(interceptors.NewRandomSeedingInterceptor)
	//--Cqrs middlewares
	container.Provide(middlewares.NewValidationMiddleware)
	container.Provide(middlewares.NewTransactionMiddleware)

	//Domain layer
	//--Repositories
	container.Provide(repositories.NewGameRepository)

	//View layer
	//--Projectors
	container.Provide(projectors.NewGamePaginationProjector)
	container.Provide(projectors.NewGameDetailProjector)

	//Application layer
	//--Queries
	container.Provide(queries.NewFindGamePaginationByLimitByOffsetQueryHandler)
	container.Provide(queries.NewGetGameDetailByIDByUserIDQueryHandler)
	//--Commands
	container.Provide(commands.NewCreateGameCommandHandler)
	container.Provide(commands.NewJoinGameCommandHandler)
	container.Provide(commands.NewStartGameCommandHandler)
	container.Provide(commands.NewBuildSettlementAndRoadCommandHandler)
	container.Provide(commands.NewRollDicesCommandHandler)
	container.Provide(commands.NewDiscardResourceCardsCommandHandler)
	container.Provide(commands.NewMoveRobberCommandHandler)
	container.Provide(commands.NewEndTurnCommandHandler)
	container.Provide(commands.NewBuildSettlementCommandHandler)
	container.Provide(commands.NewBuildRoadCommandHandler)
	container.Provide(commands.NewUpgradeCityCommandHandler)
	container.Provide(commands.NewBuyDevelopmentCardCommandHandler)
	container.Provide(commands.NewToggleResourceCardsCommandHandler)
	container.Provide(commands.NewMaritimeTradeCommandHandler)
	container.Provide(commands.NewSendTradeOfferCommandHandler)
	container.Provide(commands.NewConfirmTradeOfferCommandHandler)
	container.Provide(commands.NewCancelTradeOfferCommandHandler)
	container.Provide(commands.NewPlayKnightCardCommandHandler)
	container.Provide(commands.NewPlayRoadBuildingCardCommandHandler)
	container.Provide(commands.NewPlayYearOfPlentyCardCommandHandler)
	container.Provide(commands.NewPlayMonopolyCardCommandHandler)
	container.Provide(commands.NewPlayVictoryPointCardCommandHandler)

	//Presentation layer
	container.Provide(presentation.NewServer)
	container.Provide(v1.NewCatanServer)

	return container
}
