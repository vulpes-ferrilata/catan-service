package servers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/application/queries"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCatanServer(findGamesByUserIDQueryHandler query.QueryHandler[*queries.FindGamesByUserIDQuery, []*models.Game],
	getGameByIDByUserIDQueryHandler query.QueryHandler[*queries.GetGameByIDByUserIDQuery, *models.Game],
	createGameCommandHandler command.CommandHandler[*commands.CreateGameCommand],
	joinGameCommandHandler command.CommandHandler[*commands.JoinGameCommand],
	startGameCommandHandler command.CommandHandler[*commands.StartGameCommand],
	buildSettlementAndRoadCommandHandler command.CommandHandler[*commands.BuildSettlementAndRoadCommand],
	rollDicesCommandHandler command.CommandHandler[*commands.RollDicesCommand],
	moveRobberCommandHandler command.CommandHandler[*commands.MoveRobberCommand],
	endTurnCommandHandler command.CommandHandler[*commands.EndTurnCommand],
	buildSettlementCommandHandler command.CommandHandler[*commands.BuildSettlementCommand],
	buildRoadCommand command.CommandHandler[*commands.BuildRoadCommand],
	upgradeCityCommandHandler command.CommandHandler[*commands.UpgradeCityCommand],
	buyDevelopmentCardCommandHandler command.CommandHandler[*commands.BuyDevelopmentCardCommand],
	toggleResourceCardsCommandHandler command.CommandHandler[*commands.ToggleResourceCardsCommand],
	maritimeTradeCommandHandler command.CommandHandler[*commands.MaritimeTradeCommand],
	offerTradingCommandHandler command.CommandHandler[*commands.OfferTradingCommand],
	confirmTradingCommandHandler command.CommandHandler[*commands.ConfirmTradingCommand],
	cancelTradingCommandHandler command.CommandHandler[*commands.CancelTradingCommand],
	playKnightCardCommandHandler command.CommandHandler[*commands.PlayKnightCardCommand],
	playRoadBuildingCardCommandHandler command.CommandHandler[*commands.PlayRoadBuildingCardCommand],
	playYearOfPlentyCardCommandHandler command.CommandHandler[*commands.PlayYearOfPlentyCardCommand],
	playMonopolyCardCommandHandler command.CommandHandler[*commands.PlayMonopolyCardCommand]) catan.CatanServer {
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
	catan.UnimplementedCatanServer
	findGamesByUserIDQueryHandler        query.QueryHandler[*queries.FindGamesByUserIDQuery, []*models.Game]
	getGameByIDByUserIDQueryHandler      query.QueryHandler[*queries.GetGameByIDByUserIDQuery, *models.Game]
	createGameCommandHandler             command.CommandHandler[*commands.CreateGameCommand]
	joinGameCommandHandler               command.CommandHandler[*commands.JoinGameCommand]
	startGameCommandHandler              command.CommandHandler[*commands.StartGameCommand]
	buildSettlementAndRoadCommandHandler command.CommandHandler[*commands.BuildSettlementAndRoadCommand]
	rollDicesCommandHandler              command.CommandHandler[*commands.RollDicesCommand]
	moveRobberCommandHandler             command.CommandHandler[*commands.MoveRobberCommand]
	endTurnCommandHandler                command.CommandHandler[*commands.EndTurnCommand]
	buildSettlementCommandHandler        command.CommandHandler[*commands.BuildSettlementCommand]
	buildRoadCommandHandler              command.CommandHandler[*commands.BuildRoadCommand]
	upgradeCityCommandHandler            command.CommandHandler[*commands.UpgradeCityCommand]
	buyDevelopmentCardCommandHandler     command.CommandHandler[*commands.BuyDevelopmentCardCommand]
	toggleResourceCardsCommandHandler    command.CommandHandler[*commands.ToggleResourceCardsCommand]
	maritimeTradeCommandHandler          command.CommandHandler[*commands.MaritimeTradeCommand]
	offerTradingCommandHandler           command.CommandHandler[*commands.OfferTradingCommand]
	confirmTradingCommandHandler         command.CommandHandler[*commands.ConfirmTradingCommand]
	cancelTradingCommandHandler          command.CommandHandler[*commands.CancelTradingCommand]
	playKnightCardCommandHandler         command.CommandHandler[*commands.PlayKnightCardCommand]
	playRoadBuildingCardCommandHandler   command.CommandHandler[*commands.PlayRoadBuildingCardCommand]
	playYearOfPlentyCardCommandHandler   command.CommandHandler[*commands.PlayYearOfPlentyCardCommand]
	playMonopolyCardCommandHandler       command.CommandHandler[*commands.PlayMonopolyCardCommand]
}

func (c catanServer) FindGamesByUserID(ctx context.Context, findGamesByUserIDRequest *catan.FindGamesByUserIDRequest) (*catan.GameListResponse, error) {
	findGamesByUserIDQuery := &queries.FindGamesByUserIDQuery{
		UserID: findGamesByUserIDRequest.GetUserID(),
	}

	games, err := c.findGamesByUserIDQueryHandler.Handle(ctx, findGamesByUserIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gameResponses, _ := slices.Map(func(game *models.Game) (*catan.GameResponse, error) {
		return mappers.ToGameResponse(game), nil
	}, games)

	gameListResponse := &catan.GameListResponse{
		Games: gameResponses,
	}

	return gameListResponse, nil
}

func (c catanServer) GetGameByIDByUserID(ctx context.Context, getGameByIDByUserIDRequest *catan.GetGameByIDByUserIDRequest) (*catan.GameResponse, error) {
	getGameByIDByUserIDQuery := &queries.GetGameByIDByUserIDQuery{
		GameID: getGameByIDByUserIDRequest.GetGameID(),
		UserID: getGameByIDByUserIDRequest.GetUserID(),
	}

	game, err := c.getGameByIDByUserIDQueryHandler.Handle(ctx, getGameByIDByUserIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gameDetailResponse := mappers.ToGameResponse(game)

	return gameDetailResponse, nil
}

func (c catanServer) CreateGame(ctx context.Context, createGameRequest *catan.CreateGameRequest) (*emptypb.Empty, error) {
	createGameCommand := &commands.CreateGameCommand{
		UserID: createGameRequest.GetUserID(),
		GameID: createGameRequest.GetGameID(),
	}

	if err := c.createGameCommandHandler.Handle(ctx, createGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) JoinGame(ctx context.Context, joinGameRequest *catan.JoinGameRequest) (*emptypb.Empty, error) {
	joinGameCommand := &commands.JoinGameCommand{
		UserID: joinGameRequest.GetUserID(),
		GameID: joinGameRequest.GetGameID(),
	}

	if err := c.joinGameCommandHandler.Handle(ctx, joinGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) StartGame(ctx context.Context, startGameRequest *catan.StartGameRequest) (*emptypb.Empty, error) {
	startGameCommand := &commands.StartGameCommand{
		UserID: startGameRequest.GetUserID(),
		GameID: startGameRequest.GetGameID(),
	}

	if err := c.startGameCommandHandler.Handle(ctx, startGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildSettlementAndRoad(ctx context.Context, buildSettlementAndRoadRequest *catan.BuildSettlementAndRoadRequest) (*emptypb.Empty, error) {
	buildSettlementAndRoadCommand := &commands.BuildSettlementAndRoadCommand{
		UserID: buildSettlementAndRoadRequest.GetUserID(),
		GameID: buildSettlementAndRoadRequest.GetGameID(),
		LandID: buildSettlementAndRoadRequest.GetLandID(),
		PathID: buildSettlementAndRoadRequest.GetPathID(),
	}

	if err := c.buildSettlementAndRoadCommandHandler.Handle(ctx, buildSettlementAndRoadCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) RollDices(ctx context.Context, rollDicesRequest *catan.RollDicesRequest) (*emptypb.Empty, error) {
	rollDicesCommand := &commands.RollDicesCommand{
		UserID: rollDicesRequest.GetUserID(),
		GameID: rollDicesRequest.GetGameID(),
	}

	if err := c.rollDicesCommandHandler.Handle(ctx, rollDicesCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) MoveRobber(ctx context.Context, moveRobberRequest *catan.MoveRobberRequest) (*emptypb.Empty, error) {
	moveRobberCommand := &commands.MoveRobberCommand{
		UserID:    moveRobberRequest.GetUserID(),
		GameID:    moveRobberRequest.GetGameID(),
		TerrainID: moveRobberRequest.GetTerrainID(),
		PlayerID:  moveRobberRequest.GetPlayerID(),
	}

	if err := c.moveRobberCommandHandler.Handle(ctx, moveRobberCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) EndTurn(ctx context.Context, endTurnRequest *catan.EndTurnRequest) (*emptypb.Empty, error) {
	endTurnCommand := &commands.EndTurnCommand{
		UserID: endTurnRequest.GetUserID(),
		GameID: endTurnRequest.GetGameID(),
	}

	if err := c.endTurnCommandHandler.Handle(ctx, endTurnCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildSettlement(ctx context.Context, buildSettlementRequest *catan.BuildSettlementRequest) (*emptypb.Empty, error) {
	buildSettlementCommand := &commands.BuildSettlementCommand{
		UserID: buildSettlementRequest.GetUserID(),
		GameID: buildSettlementRequest.GetGameID(),
		LandID: buildSettlementRequest.GetLandID(),
	}

	if err := c.buildSettlementCommandHandler.Handle(ctx, buildSettlementCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildRoad(ctx context.Context, buildRoadRequest *catan.BuildRoadRequest) (*emptypb.Empty, error) {
	buildRoadCommand := &commands.BuildRoadCommand{
		UserID: buildRoadRequest.GetUserID(),
		GameID: buildRoadRequest.GetGameID(),
		PathID: buildRoadRequest.GetPathID(),
	}

	if err := c.buildRoadCommandHandler.Handle(ctx, buildRoadCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) UpgradeCity(ctx context.Context, upgradeCityRequest *catan.UpgradeCityRequest) (*emptypb.Empty, error) {
	upgradeCityCommand := &commands.UpgradeCityCommand{
		UserID:         upgradeCityRequest.GetUserID(),
		GameID:         upgradeCityRequest.GetGameID(),
		ConstructionID: upgradeCityRequest.GetConstructionID(),
	}

	if err := c.upgradeCityCommandHandler.Handle(ctx, upgradeCityCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuyDevelopmentCard(ctx context.Context, buyDevelopmentCardRequest *catan.BuyDevelopmentCardRequest) (*emptypb.Empty, error) {
	buyDevelopmentCardCommand := &commands.BuyDevelopmentCardCommand{
		UserID: buyDevelopmentCardRequest.GetUserID(),
		GameID: buyDevelopmentCardRequest.GetGameID(),
	}

	if err := c.buyDevelopmentCardCommandHandler.Handle(ctx, buyDevelopmentCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) ToggleResourceCards(ctx context.Context, toggleResourceCardsRequest *catan.ToggleResourceCardsRequest) (*emptypb.Empty, error) {
	toggleResourceCardCommand := &commands.ToggleResourceCardsCommand{
		UserID:          toggleResourceCardsRequest.GetUserID(),
		GameID:          toggleResourceCardsRequest.GetGameID(),
		ResourceCardIDs: toggleResourceCardsRequest.GetResourceCardIDs(),
	}

	if err := c.toggleResourceCardsCommandHandler.Handle(ctx, toggleResourceCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) MaritimeTrade(ctx context.Context, maritimeTradeRequest *catan.MaritimeTradeRequest) (*emptypb.Empty, error) {
	maritimeTradeCommand := &commands.MaritimeTradeCommand{
		UserID:           maritimeTradeRequest.GetUserID(),
		GameID:           maritimeTradeRequest.GetGameID(),
		ResourceCardType: maritimeTradeRequest.GetResourceCardType(),
	}

	if err := c.maritimeTradeCommandHandler.Handle(ctx, maritimeTradeCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) OfferTrading(ctx context.Context, offerTradingRequest *catan.OfferTradingRequest) (*emptypb.Empty, error) {
	offerTradingCommand := &commands.OfferTradingCommand{
		UserID:   offerTradingRequest.GetUserID(),
		GameID:   offerTradingRequest.GetGameID(),
		PlayerID: offerTradingRequest.GetPlayerID(),
	}

	if err := c.offerTradingCommandHandler.Handle(ctx, offerTradingCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) ConfirmTrading(ctx context.Context, confirmTradingRequest *catan.ConfirmTradingRequest) (*emptypb.Empty, error) {
	confirmTradingCommand := &commands.ConfirmTradingCommand{
		UserID: confirmTradingRequest.GetUserID(),
		GameID: confirmTradingRequest.GetGameID(),
	}

	if err := c.confirmTradingCommandHandler.Handle(ctx, confirmTradingCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) CancelTrading(ctx context.Context, cancelTradingRequest *catan.CancelTradingRequest) (*emptypb.Empty, error) {
	cancelTradingCommand := &commands.CancelTradingCommand{
		UserID: cancelTradingRequest.GetUserID(),
		GameID: cancelTradingRequest.GetGameID(),
	}

	if err := c.cancelTradingCommandHandler.Handle(ctx, cancelTradingCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayKnightCard(ctx context.Context, playKnightCardRequest *catan.PlayKnightCardRequest) (*emptypb.Empty, error) {
	playKnightCardCommand := &commands.PlayKnightCardCommand{
		UserID:    playKnightCardRequest.GetUserID(),
		GameID:    playKnightCardRequest.GetGameID(),
		TerrainID: playKnightCardRequest.GetTerrainID(),
		PlayerID:  playKnightCardRequest.GetPlayerID(),
	}

	if err := c.playKnightCardCommandHandler.Handle(ctx, playKnightCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayRoadBuildingCard(ctx context.Context, playRoadBuildingCardRequest *catan.PlayRoadBuildingCardRequest) (*emptypb.Empty, error) {
	playRoadBuildingCardCommand := &commands.PlayRoadBuildingCardCommand{
		UserID:  playRoadBuildingCardRequest.GetUserID(),
		GameID:  playRoadBuildingCardRequest.GetGameID(),
		PathIDs: playRoadBuildingCardRequest.GetPathIDs(),
	}

	if err := c.playRoadBuildingCardCommandHandler.Handle(ctx, playRoadBuildingCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayYearOfPlentyCard(ctx context.Context, playYearOfPlentyCardRequest *catan.PlayYearOfPlentyCardRequest) (*emptypb.Empty, error) {
	playYearOfPlentyCardCommand := &commands.PlayYearOfPlentyCardCommand{
		UserID:            playYearOfPlentyCardRequest.GetUserID(),
		GameID:            playYearOfPlentyCardRequest.GetGameID(),
		ResourceCardTypes: playYearOfPlentyCardRequest.GetResourceCardTypes(),
	}

	if err := c.playYearOfPlentyCardCommandHandler.Handle(ctx, playYearOfPlentyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayMonopolyCard(ctx context.Context, playMonopolyCardRequest *catan.PlayMonopolyCardRequest) (*emptypb.Empty, error) {
	playMonopolyCardCommand := &commands.PlayMonopolyCardCommand{
		UserID:           playMonopolyCardRequest.GetUserID(),
		GameID:           playMonopolyCardRequest.GetGameID(),
		ResourceCardType: playMonopolyCardRequest.GetResourceCardType(),
	}

	if err := c.playMonopolyCardCommandHandler.Handle(ctx, playMonopolyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
