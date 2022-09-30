package servers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service-proto/pb"
	"github.com/vulpes-ferrilata/catan-service-proto/pb/requests"
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/application/queries"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCatanServer(findGamesByUserIDQueryHandler query.QueryHandler[*queries.FindGamesByUserID, []*models.Game],
	getGameByIDByUserIDQueryHandler query.QueryHandler[*queries.GetGameByIDByUserID, *models.Game],
	createGameCommandHandler command.CommandHandler[*commands.CreateGame],
	joinGameCommandHandler command.CommandHandler[*commands.JoinGame],
	startGameCommandHandler command.CommandHandler[*commands.StartGame],
	buildSettlementAndRoadCommandHandler command.CommandHandler[*commands.BuildSettlementAndRoad],
	rollDicesCommandHandler command.CommandHandler[*commands.RollDices],
	moveRobberCommandHandler command.CommandHandler[*commands.MoveRobber],
	endTurnCommandHandler command.CommandHandler[*commands.EndTurn],
	buildSettlementCommandHandler command.CommandHandler[*commands.BuildSettlement],
	buildRoadCommand command.CommandHandler[*commands.BuildRoad],
	upgradeCityCommandHandler command.CommandHandler[*commands.UpgradeCity],
	buyDevelopmentCardCommandHandler command.CommandHandler[*commands.BuyDevelopmentCard],
	toggleResourceCardsCommandHandler command.CommandHandler[*commands.ToggleResourceCards],
	maritimeTradeCommandHandler command.CommandHandler[*commands.MaritimeTrade],
	offerTradingCommandHandler command.CommandHandler[*commands.OfferTrading],
	confirmTradingCommandHandler command.CommandHandler[*commands.ConfirmTrading],
	cancelTradingCommandHandler command.CommandHandler[*commands.CancelTrading],
	playKnightCardCommandHandler command.CommandHandler[*commands.PlayKnightCard],
	playRoadBuildingCardCommandHandler command.CommandHandler[*commands.PlayRoadBuildingCard],
	playYearOfPlentyCardCommandHandler command.CommandHandler[*commands.PlayYearOfPlentyCard],
	playMonopolyCardCommandHandler command.CommandHandler[*commands.PlayMonopolyCard]) pb.CatanServer {
	return &catanServer{
		findGamesByUserIDQueryHandler:        findGamesByUserIDQueryHandler,
		getGameByIDByUserIDQueryHandler:      getGameByIDByUserIDQueryHandler,
		createGameCommandHandler:             createGameCommandHandler,
		joinGameCommandHandler:               joinGameCommandHandler,
		startGameCommandHandler:              startGameCommandHandler,
		buildSettlementAndRoadCommandHandler: buildSettlementAndRoadCommandHandler,
		rollDicesCommandHandler:              rollDicesCommandHandler,
		moveRobberCommandHandler:             moveRobberCommandHandler,
		endTurnCommandHandler:                endTurnCommandHandler,
		buildSettlementCommandHandler:        buildSettlementCommandHandler,
		buildRoadCommandHandler:              buildRoadCommand,
		upgradeCityCommandHandler:            upgradeCityCommandHandler,
		buyDevelopmentCardCommandHandler:     buyDevelopmentCardCommandHandler,
		toggleResourceCardsCommandHandler:    toggleResourceCardsCommandHandler,
		maritimeTradeCommandHandler:          maritimeTradeCommandHandler,
		offerTradingCommandHandler:           offerTradingCommandHandler,
		confirmTradingCommandHandler:         confirmTradingCommandHandler,
		cancelTradingCommandHandler:          cancelTradingCommandHandler,
		playKnightCardCommandHandler:         playKnightCardCommandHandler,
		playRoadBuildingCardCommandHandler:   playRoadBuildingCardCommandHandler,
		playYearOfPlentyCardCommandHandler:   playYearOfPlentyCardCommandHandler,
		playMonopolyCardCommandHandler:       playMonopolyCardCommandHandler,
	}
}

type catanServer struct {
	pb.UnimplementedCatanServer
	findGamesByUserIDQueryHandler        query.QueryHandler[*queries.FindGamesByUserID, []*models.Game]
	getGameByIDByUserIDQueryHandler      query.QueryHandler[*queries.GetGameByIDByUserID, *models.Game]
	createGameCommandHandler             command.CommandHandler[*commands.CreateGame]
	joinGameCommandHandler               command.CommandHandler[*commands.JoinGame]
	startGameCommandHandler              command.CommandHandler[*commands.StartGame]
	buildSettlementAndRoadCommandHandler command.CommandHandler[*commands.BuildSettlementAndRoad]
	rollDicesCommandHandler              command.CommandHandler[*commands.RollDices]
	moveRobberCommandHandler             command.CommandHandler[*commands.MoveRobber]
	endTurnCommandHandler                command.CommandHandler[*commands.EndTurn]
	buildSettlementCommandHandler        command.CommandHandler[*commands.BuildSettlement]
	buildRoadCommandHandler              command.CommandHandler[*commands.BuildRoad]
	upgradeCityCommandHandler            command.CommandHandler[*commands.UpgradeCity]
	buyDevelopmentCardCommandHandler     command.CommandHandler[*commands.BuyDevelopmentCard]
	toggleResourceCardsCommandHandler    command.CommandHandler[*commands.ToggleResourceCards]
	maritimeTradeCommandHandler          command.CommandHandler[*commands.MaritimeTrade]
	offerTradingCommandHandler           command.CommandHandler[*commands.OfferTrading]
	confirmTradingCommandHandler         command.CommandHandler[*commands.ConfirmTrading]
	cancelTradingCommandHandler          command.CommandHandler[*commands.CancelTrading]
	playKnightCardCommandHandler         command.CommandHandler[*commands.PlayKnightCard]
	playRoadBuildingCardCommandHandler   command.CommandHandler[*commands.PlayRoadBuildingCard]
	playYearOfPlentyCardCommandHandler   command.CommandHandler[*commands.PlayYearOfPlentyCard]
	playMonopolyCardCommandHandler       command.CommandHandler[*commands.PlayMonopolyCard]
}

func (c catanServer) FindGamesByUserID(ctx context.Context, findGamesByUserIDRequest *requests.FindGamesByUserID) (*responses.GameList, error) {
	findGamesByUserIDQuery := &queries.FindGamesByUserID{
		UserID: findGamesByUserIDRequest.GetUserID(),
	}

	games, err := c.findGamesByUserIDQueryHandler.Handle(ctx, findGamesByUserIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gameResponses, _ := slices.Map(func(game *models.Game) (*responses.Game, error) {
		return mappers.ToGameResponse(game), nil
	}, games)

	gameListResponse := &responses.GameList{
		Games: gameResponses,
	}

	return gameListResponse, nil
}

func (c catanServer) GetGameByIDByUserID(ctx context.Context, getGameByIDByUserIDRequest *requests.GetGameByIDByUserID) (*responses.Game, error) {
	getGameByIDByUserIDQuery := &queries.GetGameByIDByUserID{
		GameID: getGameByIDByUserIDRequest.GetGameID(),
		UserID: getGameByIDByUserIDRequest.GetUserID(),
	}

	game, err := c.getGameByIDByUserIDQueryHandler.Handle(ctx, getGameByIDByUserIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gameResponse := mappers.ToGameResponse(game)

	return gameResponse, nil
}

func (c catanServer) CreateGame(ctx context.Context, createGameRequest *requests.CreateGame) (*emptypb.Empty, error) {
	createGameCommand := &commands.CreateGame{
		GameID: createGameRequest.GetGameID(),
		UserID: createGameRequest.GetUserID(),
	}

	if err := c.createGameCommandHandler.Handle(ctx, createGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) JoinGame(ctx context.Context, joinGameRequest *requests.JoinGame) (*emptypb.Empty, error) {
	joinGameCommand := &commands.JoinGame{
		GameID: joinGameRequest.GetGameID(),
		UserID: joinGameRequest.GetUserID(),
	}

	if err := c.joinGameCommandHandler.Handle(ctx, joinGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) StartGame(ctx context.Context, startGameRequest *requests.StartGame) (*emptypb.Empty, error) {
	startGameCommand := &commands.StartGame{
		GameID: startGameRequest.GetGameID(),
		UserID: startGameRequest.GetUserID(),
	}

	if err := c.startGameCommandHandler.Handle(ctx, startGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildSettlementAndRoad(ctx context.Context, buildSettlementAndRoadRequest *requests.BuildSettlementAndRoad) (*emptypb.Empty, error) {
	buildSettlementAndRoadCommand := &commands.BuildSettlementAndRoad{
		GameID: buildSettlementAndRoadRequest.GetGameID(),
		UserID: buildSettlementAndRoadRequest.GetUserID(),
		LandID: buildSettlementAndRoadRequest.GetLandID(),
		PathID: buildSettlementAndRoadRequest.GetPathID(),
	}

	if err := c.buildSettlementAndRoadCommandHandler.Handle(ctx, buildSettlementAndRoadCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) RollDices(ctx context.Context, rollDicesRequest *requests.RollDices) (*emptypb.Empty, error) {
	rollDicesCommand := &commands.RollDices{
		GameID: rollDicesRequest.GetGameID(),
		UserID: rollDicesRequest.GetUserID(),
	}

	if err := c.rollDicesCommandHandler.Handle(ctx, rollDicesCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) MoveRobber(ctx context.Context, moveRobberRequest *requests.MoveRobber) (*emptypb.Empty, error) {
	moveRobberCommand := &commands.MoveRobber{
		GameID:    moveRobberRequest.GetGameID(),
		UserID:    moveRobberRequest.GetUserID(),
		TerrainID: moveRobberRequest.GetTerrainID(),
		PlayerID:  moveRobberRequest.GetPlayerID(),
	}

	if err := c.moveRobberCommandHandler.Handle(ctx, moveRobberCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) EndTurn(ctx context.Context, endTurnRequest *requests.EndTurn) (*emptypb.Empty, error) {
	endTurnCommand := &commands.EndTurn{
		GameID: endTurnRequest.GetGameID(),
		UserID: endTurnRequest.GetUserID(),
	}

	if err := c.endTurnCommandHandler.Handle(ctx, endTurnCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildSettlement(ctx context.Context, buildSettlementRequest *requests.BuildSettlement) (*emptypb.Empty, error) {
	buildSettlementCommand := &commands.BuildSettlement{
		GameID: buildSettlementRequest.GetGameID(),
		UserID: buildSettlementRequest.GetUserID(),
		LandID: buildSettlementRequest.GetLandID(),
	}

	if err := c.buildSettlementCommandHandler.Handle(ctx, buildSettlementCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildRoad(ctx context.Context, buildRoadRequest *requests.BuildRoad) (*emptypb.Empty, error) {
	buildRoadCommand := &commands.BuildRoad{
		GameID: buildRoadRequest.GetGameID(),
		UserID: buildRoadRequest.GetUserID(),
		PathID: buildRoadRequest.GetPathID(),
	}

	if err := c.buildRoadCommandHandler.Handle(ctx, buildRoadCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) UpgradeCity(ctx context.Context, upgradeCityRequest *requests.UpgradeCity) (*emptypb.Empty, error) {
	upgradeCityCommand := &commands.UpgradeCity{
		GameID:         upgradeCityRequest.GetGameID(),
		UserID:         upgradeCityRequest.GetUserID(),
		ConstructionID: upgradeCityRequest.GetConstructionID(),
	}

	if err := c.upgradeCityCommandHandler.Handle(ctx, upgradeCityCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuyDevelopmentCard(ctx context.Context, buyDevelopmentCardRequest *requests.BuyDevelopmentCard) (*emptypb.Empty, error) {
	buyDevelopmentCardCommand := &commands.BuyDevelopmentCard{
		GameID: buyDevelopmentCardRequest.GetGameID(),
		UserID: buyDevelopmentCardRequest.GetUserID(),
	}

	if err := c.buyDevelopmentCardCommandHandler.Handle(ctx, buyDevelopmentCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) ToggleResourceCards(ctx context.Context, toggleResourceCardsRequest *requests.ToggleResourceCards) (*emptypb.Empty, error) {
	toggleResourceCardCommand := &commands.ToggleResourceCards{
		GameID:          toggleResourceCardsRequest.GetGameID(),
		UserID:          toggleResourceCardsRequest.GetUserID(),
		ResourceCardIDs: toggleResourceCardsRequest.GetResourceCardIDs(),
	}

	if err := c.toggleResourceCardsCommandHandler.Handle(ctx, toggleResourceCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) MaritimeTrade(ctx context.Context, maritimeTradeRequest *requests.MaritimeTrade) (*emptypb.Empty, error) {
	maritimeTradeCommand := &commands.MaritimeTrade{
		GameID:           maritimeTradeRequest.GetGameID(),
		UserID:           maritimeTradeRequest.GetUserID(),
		ResourceCardType: maritimeTradeRequest.GetResourceCardType(),
	}

	if err := c.maritimeTradeCommandHandler.Handle(ctx, maritimeTradeCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) OfferTrading(ctx context.Context, offerTradingRequest *requests.OfferTrading) (*emptypb.Empty, error) {
	offerTradingCommand := &commands.OfferTrading{
		GameID:   offerTradingRequest.GetGameID(),
		UserID:   offerTradingRequest.GetUserID(),
		PlayerID: offerTradingRequest.GetPlayerID(),
	}

	if err := c.offerTradingCommandHandler.Handle(ctx, offerTradingCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) ConfirmTrading(ctx context.Context, confirmTradingRequest *requests.ConfirmTrading) (*emptypb.Empty, error) {
	confirmTradingCommand := &commands.ConfirmTrading{
		GameID: confirmTradingRequest.GetGameID(),
		UserID: confirmTradingRequest.GetUserID(),
	}

	if err := c.confirmTradingCommandHandler.Handle(ctx, confirmTradingCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) CancelTrading(ctx context.Context, cancelTradingRequest *requests.CancelTrading) (*emptypb.Empty, error) {
	cancelTradingCommand := &commands.CancelTrading{
		GameID: cancelTradingRequest.GetGameID(),
		UserID: cancelTradingRequest.GetUserID(),
	}

	if err := c.cancelTradingCommandHandler.Handle(ctx, cancelTradingCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayKnightCard(ctx context.Context, playKnightCardRequest *requests.PlayKnightCard) (*emptypb.Empty, error) {
	playKnightCardCommand := &commands.PlayKnightCard{
		GameID:    playKnightCardRequest.GetGameID(),
		UserID:    playKnightCardRequest.GetUserID(),
		TerrainID: playKnightCardRequest.GetTerrainID(),
		PlayerID:  playKnightCardRequest.GetPlayerID(),
	}

	if err := c.playKnightCardCommandHandler.Handle(ctx, playKnightCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayRoadBuildingCard(ctx context.Context, playRoadBuildingCardRequest *requests.PlayRoadBuildingCard) (*emptypb.Empty, error) {
	playRoadBuildingCardCommand := &commands.PlayRoadBuildingCard{
		GameID:  playRoadBuildingCardRequest.GetGameID(),
		UserID:  playRoadBuildingCardRequest.GetUserID(),
		PathIDs: playRoadBuildingCardRequest.GetPathIDs(),
	}

	if err := c.playRoadBuildingCardCommandHandler.Handle(ctx, playRoadBuildingCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayYearOfPlentyCard(ctx context.Context, playYearOfPlentyCardRequest *requests.PlayYearOfPlentyCard) (*emptypb.Empty, error) {
	playYearOfPlentyCardCommand := &commands.PlayYearOfPlentyCard{
		GameID:            playYearOfPlentyCardRequest.GetGameID(),
		UserID:            playYearOfPlentyCardRequest.GetUserID(),
		ResourceCardTypes: playYearOfPlentyCardRequest.GetResourceCardTypes(),
	}

	if err := c.playYearOfPlentyCardCommandHandler.Handle(ctx, playYearOfPlentyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayMonopolyCard(ctx context.Context, playMonopolyCardRequest *requests.PlayMonopolyCard) (*emptypb.Empty, error) {
	playMonopolyCardCommand := &commands.PlayMonopolyCard{
		GameID:           playMonopolyCardRequest.GetGameID(),
		UserID:           playMonopolyCardRequest.GetUserID(),
		ResourceCardType: playMonopolyCardRequest.GetResourceCardType(),
	}

	if err := c.playMonopolyCardCommandHandler.Handle(ctx, playMonopolyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
