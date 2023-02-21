package infrastructure

import (
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/middlewares"
	"github.com/vulpes-ferrilata/cqrs"
)

func NewCommandBus(validationMiddleware *middlewares.ValidationMiddleware,
	transactionMiddleware *middlewares.TransactionMiddleware,
	buildRoadCommandHandler *commands.BuildRoadCommandHandler,
	buildSettlementAndRoadCommandHandler *commands.BuildSettlementAndRoadCommandHandler,
	buildSettlementCommandHandler *commands.BuildSettlementCommandHandler,
	buyDevelopmentCardCommandHandler *commands.BuyDevelopmentCardCommandHandler,
	cancelTradeOfferCommandHandler *commands.CancelTradeOfferCommandHandler,
	confirmTradeOfferCommandHandler *commands.ConfirmTradeOfferCommandHandler,
	createGameCommandHandler *commands.CreateGameCommandHandler,
	discardResourceCardsCommandHandler *commands.DiscardResourceCardsCommandHandler,
	endTurnCommandHandler *commands.EndTurnCommandHandler,
	joinGameCommandHandler *commands.JoinGameCommandHandler,
	maritimeTradeCommandHandler *commands.MaritimeTradeCommandHandler,
	moveRobberCommandHandler *commands.MoveRobberCommandHandler,
	playKnightCardCommandHandler *commands.PlayKnightCardCommandHandler,
	playMonopolyCardCommandHandler *commands.PlayMonopolyCardCommandHandler,
	playRoadBuildingCardCommandHandler *commands.PlayRoadBuildingCardCommandHandler,
	playVictoryPointCardCommandHandler *commands.PlayVictoryPointCardCommandHandler,
	playYearOfPlentyCardCommandHandler *commands.PlayYearOfPlentyCardCommandHandler,
	rollDicesCommandHandler *commands.RollDicesCommandHandler,
	sendTradeOfferCommandHandler *commands.SendTradeOfferCommandHandler,
	startGameCommandHandler *commands.StartGameCommandHandler,
	toggleResourceCardsCommandHandler *commands.ToggleResourceCardsCommandHandler,
	upgradeCityCommandHandler *commands.UpgradeCityCommandHandler) (*cqrs.CommandBus, error) {
	commandBus := &cqrs.CommandBus{}

	commandBus.Use(
		validationMiddleware.CommandHandlerMiddleware(),
		transactionMiddleware.CommandHandlerMiddleware(),
	)

	commandBus.Register(&commands.BuildRoadCommand{}, cqrs.WrapCommandHandlerFunc(buildRoadCommandHandler.Handle))
	commandBus.Register(&commands.BuildSettlementAndRoadCommand{}, cqrs.WrapCommandHandlerFunc(buildSettlementAndRoadCommandHandler.Handle))
	commandBus.Register(&commands.BuildSettlementCommand{}, cqrs.WrapCommandHandlerFunc(buildSettlementCommandHandler.Handle))
	commandBus.Register(&commands.BuyDevelopmentCardCommand{}, cqrs.WrapCommandHandlerFunc(buyDevelopmentCardCommandHandler.Handle))
	commandBus.Register(&commands.CancelTradeOfferCommand{}, cqrs.WrapCommandHandlerFunc(cancelTradeOfferCommandHandler.Handle))
	commandBus.Register(&commands.ConfirmTradeOfferCommand{}, cqrs.WrapCommandHandlerFunc(confirmTradeOfferCommandHandler.Handle))
	commandBus.Register(&commands.CreateGameCommand{}, cqrs.WrapCommandHandlerFunc(createGameCommandHandler.Handle))
	commandBus.Register(&commands.DiscardResourceCardsCommand{}, cqrs.WrapCommandHandlerFunc(discardResourceCardsCommandHandler.Handle))
	commandBus.Register(&commands.EndTurnCommand{}, cqrs.WrapCommandHandlerFunc(endTurnCommandHandler.Handle))
	commandBus.Register(&commands.JoinGameCommand{}, cqrs.WrapCommandHandlerFunc(joinGameCommandHandler.Handle))
	commandBus.Register(&commands.MaritimeTradeCommand{}, cqrs.WrapCommandHandlerFunc(maritimeTradeCommandHandler.Handle))
	commandBus.Register(&commands.MoveRobberCommand{}, cqrs.WrapCommandHandlerFunc(moveRobberCommandHandler.Handle))
	commandBus.Register(&commands.PlayKnightCardCommand{}, cqrs.WrapCommandHandlerFunc(playKnightCardCommandHandler.Handle))
	commandBus.Register(&commands.PlayMonopolyCardCommand{}, cqrs.WrapCommandHandlerFunc(playMonopolyCardCommandHandler.Handle))
	commandBus.Register(&commands.PlayRoadBuildingCardCommand{}, cqrs.WrapCommandHandlerFunc(playRoadBuildingCardCommandHandler.Handle))
	commandBus.Register(&commands.PlayVictoryPointCardCommand{}, cqrs.WrapCommandHandlerFunc(playVictoryPointCardCommandHandler.Handle))
	commandBus.Register(&commands.PlayYearOfPlentyCardCommand{}, cqrs.WrapCommandHandlerFunc(playYearOfPlentyCardCommandHandler.Handle))
	commandBus.Register(&commands.RollDicesCommand{}, cqrs.WrapCommandHandlerFunc(rollDicesCommandHandler.Handle))
	commandBus.Register(&commands.SendTradeOfferCommand{}, cqrs.WrapCommandHandlerFunc(sendTradeOfferCommandHandler.Handle))
	commandBus.Register(&commands.StartGameCommand{}, cqrs.WrapCommandHandlerFunc(startGameCommandHandler.Handle))
	commandBus.Register(&commands.ToggleResourceCardsCommand{}, cqrs.WrapCommandHandlerFunc(toggleResourceCardsCommandHandler.Handle))
	commandBus.Register(&commands.UpgradeCityCommand{}, cqrs.WrapCommandHandlerFunc(upgradeCityCommandHandler.Handle))

	return commandBus, nil
}
