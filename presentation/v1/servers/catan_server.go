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
	"github.com/vulpes-ferrilata/catan-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCatanServer(findGamePaginationByLimitByOffsetQueryHandler query.QueryHandler[*queries.FindGamePaginationByLimitByOffset, *models.Pagination[*models.Game]],
	getGameDetailByIDByUserIDQueryHandler query.QueryHandler[*queries.GetGameDetailByIDByUserID, *models.GameDetail],
	createGameCommandHandler command.CommandHandler[*commands.CreateGame],
	joinGameCommandHandler command.CommandHandler[*commands.JoinGame],
	startGameCommandHandler command.CommandHandler[*commands.StartGame],
	buildSettlementAndRoadCommandHandler command.CommandHandler[*commands.BuildSettlementAndRoad],
	rollDicesCommandHandler command.CommandHandler[*commands.RollDices],
	discardResourceCardsCommandHandler command.CommandHandler[*commands.DiscardResourceCards],
	moveRobberCommandHandler command.CommandHandler[*commands.MoveRobber],
	endTurnCommandHandler command.CommandHandler[*commands.EndTurn],
	buildSettlementCommandHandler command.CommandHandler[*commands.BuildSettlement],
	buildRoadCommand command.CommandHandler[*commands.BuildRoad],
	upgradeCityCommandHandler command.CommandHandler[*commands.UpgradeCity],
	buyDevelopmentCardCommandHandler command.CommandHandler[*commands.BuyDevelopmentCard],
	toggleResourceCardsCommandHandler command.CommandHandler[*commands.ToggleResourceCards],
	maritimeTradeCommandHandler command.CommandHandler[*commands.MaritimeTrade],
	sendTradeOfferCommandHandler command.CommandHandler[*commands.SendTradeOffer],
	confirmTradeOfferCommandHandler command.CommandHandler[*commands.ConfirmTradeOffer],
	cancelTradeOfferCommandHandler command.CommandHandler[*commands.CancelTradeOffer],
	playKnightCardCommandHandler command.CommandHandler[*commands.PlayKnightCard],
	playRoadBuildingCardCommandHandler command.CommandHandler[*commands.PlayRoadBuildingCard],
	playYearOfPlentyCardCommandHandler command.CommandHandler[*commands.PlayYearOfPlentyCard],
	playMonopolyCardCommandHandler command.CommandHandler[*commands.PlayMonopolyCard],
	playVictoryPointCardCommandHandler command.CommandHandler[*commands.PlayVictoryPointCard]) pb.CatanServer {
	return &catanServer{
		findGamePaginationByLimitByOffsetQueryHandler: findGamePaginationByLimitByOffsetQueryHandler,
		getGameDetailByIDByUserIDQueryHandler:         getGameDetailByIDByUserIDQueryHandler,
		createGameCommandHandler:                      createGameCommandHandler,
		joinGameCommandHandler:                        joinGameCommandHandler,
		startGameCommandHandler:                       startGameCommandHandler,
		buildSettlementAndRoadCommandHandler:          buildSettlementAndRoadCommandHandler,
		rollDicesCommandHandler:                       rollDicesCommandHandler,
		discardResourceCardsCommandHandler:            discardResourceCardsCommandHandler,
		moveRobberCommandHandler:                      moveRobberCommandHandler,
		endTurnCommandHandler:                         endTurnCommandHandler,
		buildSettlementCommandHandler:                 buildSettlementCommandHandler,
		buildRoadCommandHandler:                       buildRoadCommand,
		upgradeCityCommandHandler:                     upgradeCityCommandHandler,
		buyDevelopmentCardCommandHandler:              buyDevelopmentCardCommandHandler,
		toggleResourceCardsCommandHandler:             toggleResourceCardsCommandHandler,
		maritimeTradeCommandHandler:                   maritimeTradeCommandHandler,
		sendTradeOfferCommandHandler:                  sendTradeOfferCommandHandler,
		confirmTradeOfferCommandHandler:               confirmTradeOfferCommandHandler,
		cancelTradeOfferCommandHandler:                cancelTradeOfferCommandHandler,
		playKnightCardCommandHandler:                  playKnightCardCommandHandler,
		playRoadBuildingCardCommandHandler:            playRoadBuildingCardCommandHandler,
		playYearOfPlentyCardCommandHandler:            playYearOfPlentyCardCommandHandler,
		playMonopolyCardCommandHandler:                playMonopolyCardCommandHandler,
		playVictoryPointCardCommandHandler:            playVictoryPointCardCommandHandler,
	}
}

type catanServer struct {
	pb.UnimplementedCatanServer
	findGamePaginationByLimitByOffsetQueryHandler query.QueryHandler[*queries.FindGamePaginationByLimitByOffset, *models.Pagination[*models.Game]]
	getGameDetailByIDByUserIDQueryHandler         query.QueryHandler[*queries.GetGameDetailByIDByUserID, *models.GameDetail]
	createGameCommandHandler                      command.CommandHandler[*commands.CreateGame]
	joinGameCommandHandler                        command.CommandHandler[*commands.JoinGame]
	startGameCommandHandler                       command.CommandHandler[*commands.StartGame]
	buildSettlementAndRoadCommandHandler          command.CommandHandler[*commands.BuildSettlementAndRoad]
	rollDicesCommandHandler                       command.CommandHandler[*commands.RollDices]
	discardResourceCardsCommandHandler            command.CommandHandler[*commands.DiscardResourceCards]
	moveRobberCommandHandler                      command.CommandHandler[*commands.MoveRobber]
	endTurnCommandHandler                         command.CommandHandler[*commands.EndTurn]
	buildSettlementCommandHandler                 command.CommandHandler[*commands.BuildSettlement]
	buildRoadCommandHandler                       command.CommandHandler[*commands.BuildRoad]
	upgradeCityCommandHandler                     command.CommandHandler[*commands.UpgradeCity]
	buyDevelopmentCardCommandHandler              command.CommandHandler[*commands.BuyDevelopmentCard]
	toggleResourceCardsCommandHandler             command.CommandHandler[*commands.ToggleResourceCards]
	maritimeTradeCommandHandler                   command.CommandHandler[*commands.MaritimeTrade]
	sendTradeOfferCommandHandler                  command.CommandHandler[*commands.SendTradeOffer]
	confirmTradeOfferCommandHandler               command.CommandHandler[*commands.ConfirmTradeOffer]
	cancelTradeOfferCommandHandler                command.CommandHandler[*commands.CancelTradeOffer]
	playKnightCardCommandHandler                  command.CommandHandler[*commands.PlayKnightCard]
	playRoadBuildingCardCommandHandler            command.CommandHandler[*commands.PlayRoadBuildingCard]
	playYearOfPlentyCardCommandHandler            command.CommandHandler[*commands.PlayYearOfPlentyCard]
	playMonopolyCardCommandHandler                command.CommandHandler[*commands.PlayMonopolyCard]
	playVictoryPointCardCommandHandler            command.CommandHandler[*commands.PlayVictoryPointCard]
}

func (c catanServer) FindGamePaginationByLimitByOffset(ctx context.Context, findGamePaginationByLimitByOffsetRequest *requests.FindGamePaginationByLimitByOffset) (*responses.GamePagination, error) {
	findGamePaginationByLimitByOffsetQuery := &queries.FindGamePaginationByLimitByOffset{
		Limit:  int(findGamePaginationByLimitByOffsetRequest.GetLimit()),
		Offset: int(findGamePaginationByLimitByOffsetRequest.GetOffset()),
	}

	gamePagination, err := c.findGamePaginationByLimitByOffsetQueryHandler.Handle(ctx, findGamePaginationByLimitByOffsetQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gamePaginationResponse, err := mappers.GamePaginationMapper{}.ToResponse(gamePagination)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gamePaginationResponse, nil
}

func (c catanServer) GetGameDetailByIDByUserID(ctx context.Context, getGameDetailByIDByUserIDRequest *requests.GetGameDetailByIDByUserID) (*responses.GameDetail, error) {
	getGameDetailByIDByUserIDQuery := &queries.GetGameDetailByIDByUserID{
		GameID: getGameDetailByIDByUserIDRequest.GetGameID(),
		UserID: getGameDetailByIDByUserIDRequest.GetUserID(),
	}

	gameDetail, err := c.getGameDetailByIDByUserIDQueryHandler.Handle(ctx, getGameDetailByIDByUserIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gameDetailResponse, err := mappers.GameDetailMapper{}.ToResponse(gameDetail)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gameDetailResponse, nil
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

func (c catanServer) DiscardResourceCards(ctx context.Context, discardResourceCardsRequest *requests.DiscardResourceCards) (*emptypb.Empty, error) {
	discardResourceCardsCommand := &commands.DiscardResourceCards{
		GameID:          discardResourceCardsRequest.GetGameID(),
		UserID:          discardResourceCardsRequest.GetUserID(),
		ResourceCardIDs: discardResourceCardsRequest.GetResourceCardIDs(),
	}

	if err := c.discardResourceCardsCommandHandler.Handle(ctx, discardResourceCardsCommand); err != nil {
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
		GameID:                    maritimeTradeRequest.GetGameID(),
		UserID:                    maritimeTradeRequest.GetUserID(),
		ResourceCardType:          maritimeTradeRequest.GetResourceCardType(),
		DemandingResourceCardType: maritimeTradeRequest.GetDemandingResourceCardType(),
	}

	if err := c.maritimeTradeCommandHandler.Handle(ctx, maritimeTradeCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) SendTradeOffer(ctx context.Context, sendTradeOfferRequest *requests.SendTradeOffer) (*emptypb.Empty, error) {
	sendTradeOfferCommand := &commands.SendTradeOffer{
		GameID:   sendTradeOfferRequest.GetGameID(),
		UserID:   sendTradeOfferRequest.GetUserID(),
		PlayerID: sendTradeOfferRequest.GetPlayerID(),
	}

	if err := c.sendTradeOfferCommandHandler.Handle(ctx, sendTradeOfferCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) ConfirmTradeOffer(ctx context.Context, confirmTradeOfferRequest *requests.ConfirmTradeOffer) (*emptypb.Empty, error) {
	confirmTradeOfferCommand := &commands.ConfirmTradeOffer{
		GameID: confirmTradeOfferRequest.GetGameID(),
		UserID: confirmTradeOfferRequest.GetUserID(),
	}

	if err := c.confirmTradeOfferCommandHandler.Handle(ctx, confirmTradeOfferCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) CancelTradeOffer(ctx context.Context, cancelTradeOfferRequest *requests.CancelTradeOffer) (*emptypb.Empty, error) {
	cancelTradeOfferCommand := &commands.CancelTradeOffer{
		GameID: cancelTradeOfferRequest.GetGameID(),
		UserID: cancelTradeOfferRequest.GetUserID(),
	}

	if err := c.cancelTradeOfferCommandHandler.Handle(ctx, cancelTradeOfferCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayKnightCard(ctx context.Context, playKnightCardRequest *requests.PlayKnightCard) (*emptypb.Empty, error) {
	playKnightCardCommand := &commands.PlayKnightCard{
		GameID:            playKnightCardRequest.GetGameID(),
		UserID:            playKnightCardRequest.GetUserID(),
		DevelopmentCardID: playKnightCardRequest.GetDevelopmentCardID(),
		TerrainID:         playKnightCardRequest.GetTerrainID(),
		PlayerID:          playKnightCardRequest.GetPlayerID(),
	}

	if err := c.playKnightCardCommandHandler.Handle(ctx, playKnightCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayRoadBuildingCard(ctx context.Context, playRoadBuildingCardRequest *requests.PlayRoadBuildingCard) (*emptypb.Empty, error) {
	playRoadBuildingCardCommand := &commands.PlayRoadBuildingCard{
		GameID:            playRoadBuildingCardRequest.GetGameID(),
		UserID:            playRoadBuildingCardRequest.GetUserID(),
		DevelopmentCardID: playRoadBuildingCardRequest.GetDevelopmentCardID(),
		PathIDs:           playRoadBuildingCardRequest.GetPathIDs(),
	}

	if err := c.playRoadBuildingCardCommandHandler.Handle(ctx, playRoadBuildingCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayYearOfPlentyCard(ctx context.Context, playYearOfPlentyCardRequest *requests.PlayYearOfPlentyCard) (*emptypb.Empty, error) {
	playYearOfPlentyCardCommand := &commands.PlayYearOfPlentyCard{
		GameID:                     playYearOfPlentyCardRequest.GetGameID(),
		UserID:                     playYearOfPlentyCardRequest.GetUserID(),
		DevelopmentCardID:          playYearOfPlentyCardRequest.GetDevelopmentCardID(),
		DemandingResourceCardTypes: playYearOfPlentyCardRequest.GetDemandingResourceCardTypes(),
	}

	if err := c.playYearOfPlentyCardCommandHandler.Handle(ctx, playYearOfPlentyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayMonopolyCard(ctx context.Context, playMonopolyCardRequest *requests.PlayMonopolyCard) (*emptypb.Empty, error) {
	playMonopolyCardCommand := &commands.PlayMonopolyCard{
		GameID:                    playMonopolyCardRequest.GetGameID(),
		UserID:                    playMonopolyCardRequest.GetUserID(),
		DevelopmentCardID:         playMonopolyCardRequest.GetDevelopmentCardID(),
		DemandingResourceCardType: playMonopolyCardRequest.GetDemandingResourceCardType(),
	}

	if err := c.playMonopolyCardCommandHandler.Handle(ctx, playMonopolyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayVictoryPointCard(ctx context.Context, playVictoryPointCardRequest *requests.PlayVictoryPointCard) (*emptypb.Empty, error) {
	playVictoryPointCardCommand := &commands.PlayVictoryPointCard{
		GameID:            playVictoryPointCardRequest.GetGameID(),
		UserID:            playVictoryPointCardRequest.GetUserID(),
		DevelopmentCardID: playVictoryPointCardRequest.GetDevelopmentCardID(),
	}

	if err := c.playVictoryPointCardCommandHandler.Handle(ctx, playVictoryPointCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
