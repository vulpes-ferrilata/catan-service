package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service-proto/pb"
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/application/queries"
	"github.com/vulpes-ferrilata/catan-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/cqrs"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewCatanServer(queryBus *cqrs.QueryBus,
	commandBus *cqrs.CommandBus) pb.CatanServer {
	return &catanServer{
		queryBus:   queryBus,
		commandBus: commandBus,
	}
}

type catanServer struct {
	pb.UnimplementedCatanServer
	queryBus   *cqrs.QueryBus
	commandBus *cqrs.CommandBus
}

func (c catanServer) FindGamePaginationByLimitByOffset(ctx context.Context, findGamePaginationByLimitByOffsetRequest *pb_models.FindGamePaginationByLimitByOffsetRequest) (*pb_models.GamePagination, error) {
	findGamePaginationByLimitByOffsetQuery := &queries.FindGamePaginationByLimitByOffsetQuery{
		Limit:  int(findGamePaginationByLimitByOffsetRequest.GetLimit()),
		Offset: int(findGamePaginationByLimitByOffsetRequest.GetOffset()),
	}

	gamePagination, err := cqrs.ParseQueryHandlerFunc[*queries.FindGamePaginationByLimitByOffsetQuery, *models.Pagination[*models.Game]](c.queryBus.Execute)(ctx, findGamePaginationByLimitByOffsetQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gamePaginationResponse, err := mappers.GamePaginationMapper{}.ToResponse(gamePagination)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gamePaginationResponse, nil
}

func (c catanServer) GetGameDetailByIDByUserID(ctx context.Context, getGameDetailByIDByUserIDRequest *pb_models.GetGameDetailByIDByUserIDRequest) (*pb_models.GameDetail, error) {
	getGameDetailByIDByUserIDQuery := &queries.GetGameDetailByIDByUserIDQuery{
		GameID: getGameDetailByIDByUserIDRequest.GetGameID(),
		UserID: getGameDetailByIDByUserIDRequest.GetUserID(),
	}

	gameDetail, err := cqrs.ParseQueryHandlerFunc[*queries.GetGameDetailByIDByUserIDQuery, *models.GameDetail](c.queryBus.Execute)(ctx, getGameDetailByIDByUserIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gameDetailResponse, err := mappers.GameDetailMapper{}.ToResponse(gameDetail)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gameDetailResponse, nil
}

func (c catanServer) CreateGame(ctx context.Context, createGameRequest *pb_models.CreateGameRequest) (*emptypb.Empty, error) {
	createGameCommand := &commands.CreateGameCommand{
		GameID: createGameRequest.GetGameID(),
		UserID: createGameRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, createGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) JoinGame(ctx context.Context, joinGameRequest *pb_models.JoinGameRequest) (*emptypb.Empty, error) {
	joinGameCommand := &commands.JoinGameCommand{
		GameID: joinGameRequest.GetGameID(),
		UserID: joinGameRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, joinGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) StartGame(ctx context.Context, startGameRequest *pb_models.StartGameRequest) (*emptypb.Empty, error) {
	startGameCommand := &commands.StartGameCommand{
		GameID: startGameRequest.GetGameID(),
		UserID: startGameRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, startGameCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildSettlementAndRoad(ctx context.Context, buildSettlementAndRoadRequest *pb_models.BuildSettlementAndRoadRequest) (*emptypb.Empty, error) {
	buildSettlementAndRoadCommand := &commands.BuildSettlementAndRoadCommand{
		GameID: buildSettlementAndRoadRequest.GetGameID(),
		UserID: buildSettlementAndRoadRequest.GetUserID(),
		LandID: buildSettlementAndRoadRequest.GetLandID(),
		PathID: buildSettlementAndRoadRequest.GetPathID(),
	}

	if err := c.commandBus.Execute(ctx, buildSettlementAndRoadCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) RollDices(ctx context.Context, rollDicesRequest *pb_models.RollDicesRequest) (*emptypb.Empty, error) {
	rollDicesCommand := &commands.RollDicesCommand{
		GameID: rollDicesRequest.GetGameID(),
		UserID: rollDicesRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, rollDicesCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) DiscardResourceCards(ctx context.Context, discardResourceCardsRequest *pb_models.DiscardResourceCardsRequest) (*emptypb.Empty, error) {
	discardResourceCardsCommand := &commands.DiscardResourceCardsCommand{
		GameID:          discardResourceCardsRequest.GetGameID(),
		UserID:          discardResourceCardsRequest.GetUserID(),
		ResourceCardIDs: discardResourceCardsRequest.GetResourceCardIDs(),
	}

	if err := c.commandBus.Execute(ctx, discardResourceCardsCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) MoveRobber(ctx context.Context, moveRobberRequest *pb_models.MoveRobberRequest) (*emptypb.Empty, error) {
	moveRobberCommand := &commands.MoveRobberCommand{
		GameID:    moveRobberRequest.GetGameID(),
		UserID:    moveRobberRequest.GetUserID(),
		TerrainID: moveRobberRequest.GetTerrainID(),
		PlayerID:  moveRobberRequest.GetPlayerID(),
	}

	if err := c.commandBus.Execute(ctx, moveRobberCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) EndTurn(ctx context.Context, endTurnRequest *pb_models.EndTurnRequest) (*emptypb.Empty, error) {
	endTurnCommand := &commands.EndTurnCommand{
		GameID: endTurnRequest.GetGameID(),
		UserID: endTurnRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, endTurnCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildSettlement(ctx context.Context, buildSettlementRequest *pb_models.BuildSettlementRequest) (*emptypb.Empty, error) {
	buildSettlementCommand := &commands.BuildSettlementCommand{
		GameID: buildSettlementRequest.GetGameID(),
		UserID: buildSettlementRequest.GetUserID(),
		LandID: buildSettlementRequest.GetLandID(),
	}

	if err := c.commandBus.Execute(ctx, buildSettlementCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuildRoad(ctx context.Context, buildRoadRequest *pb_models.BuildRoadRequest) (*emptypb.Empty, error) {
	buildRoadCommand := &commands.BuildRoadCommand{
		GameID: buildRoadRequest.GetGameID(),
		UserID: buildRoadRequest.GetUserID(),
		PathID: buildRoadRequest.GetPathID(),
	}

	if err := c.commandBus.Execute(ctx, buildRoadCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) UpgradeCity(ctx context.Context, upgradeCityRequest *pb_models.UpgradeCityRequest) (*emptypb.Empty, error) {
	upgradeCityCommand := &commands.UpgradeCityCommand{
		GameID:         upgradeCityRequest.GetGameID(),
		UserID:         upgradeCityRequest.GetUserID(),
		ConstructionID: upgradeCityRequest.GetConstructionID(),
	}

	if err := c.commandBus.Execute(ctx, upgradeCityCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) BuyDevelopmentCard(ctx context.Context, buyDevelopmentCardRequest *pb_models.BuyDevelopmentCardRequest) (*emptypb.Empty, error) {
	buyDevelopmentCardCommand := &commands.BuyDevelopmentCardCommand{
		GameID: buyDevelopmentCardRequest.GetGameID(),
		UserID: buyDevelopmentCardRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, buyDevelopmentCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) ToggleResourceCards(ctx context.Context, toggleResourceCardsRequest *pb_models.ToggleResourceCardsRequest) (*emptypb.Empty, error) {
	toggleResourceCardCommand := &commands.ToggleResourceCardsCommand{
		GameID:          toggleResourceCardsRequest.GetGameID(),
		UserID:          toggleResourceCardsRequest.GetUserID(),
		ResourceCardIDs: toggleResourceCardsRequest.GetResourceCardIDs(),
	}

	if err := c.commandBus.Execute(ctx, toggleResourceCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) MaritimeTrade(ctx context.Context, maritimeTradeRequest *pb_models.MaritimeTradeRequest) (*emptypb.Empty, error) {
	maritimeTradeCommand := &commands.MaritimeTradeCommand{
		GameID:                    maritimeTradeRequest.GetGameID(),
		UserID:                    maritimeTradeRequest.GetUserID(),
		ResourceCardType:          maritimeTradeRequest.GetResourceCardType(),
		DemandingResourceCardType: maritimeTradeRequest.GetDemandingResourceCardType(),
	}

	if err := c.commandBus.Execute(ctx, maritimeTradeCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) SendTradeOffer(ctx context.Context, sendTradeOfferRequest *pb_models.SendTradeOfferRequest) (*emptypb.Empty, error) {
	sendTradeOfferCommand := &commands.SendTradeOfferCommand{
		GameID:   sendTradeOfferRequest.GetGameID(),
		UserID:   sendTradeOfferRequest.GetUserID(),
		PlayerID: sendTradeOfferRequest.GetPlayerID(),
	}

	if err := c.commandBus.Execute(ctx, sendTradeOfferCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) ConfirmTradeOffer(ctx context.Context, confirmTradeOfferRequest *pb_models.ConfirmTradeOfferRequest) (*emptypb.Empty, error) {
	confirmTradeOfferCommand := &commands.ConfirmTradeOfferCommand{
		GameID: confirmTradeOfferRequest.GetGameID(),
		UserID: confirmTradeOfferRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, confirmTradeOfferCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) CancelTradeOffer(ctx context.Context, cancelTradeOfferRequest *pb_models.CancelTradeOfferRequest) (*emptypb.Empty, error) {
	cancelTradeOfferCommand := &commands.CancelTradeOfferCommand{
		GameID: cancelTradeOfferRequest.GetGameID(),
		UserID: cancelTradeOfferRequest.GetUserID(),
	}

	if err := c.commandBus.Execute(ctx, cancelTradeOfferCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayKnightCard(ctx context.Context, playKnightCardRequest *pb_models.PlayKnightCardRequest) (*emptypb.Empty, error) {
	playKnightCardCommand := &commands.PlayKnightCardCommand{
		GameID:            playKnightCardRequest.GetGameID(),
		UserID:            playKnightCardRequest.GetUserID(),
		DevelopmentCardID: playKnightCardRequest.GetDevelopmentCardID(),
		TerrainID:         playKnightCardRequest.GetTerrainID(),
		PlayerID:          playKnightCardRequest.GetPlayerID(),
	}

	if err := c.commandBus.Execute(ctx, playKnightCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayRoadBuildingCard(ctx context.Context, playRoadBuildingCardRequest *pb_models.PlayRoadBuildingCardRequest) (*emptypb.Empty, error) {
	playRoadBuildingCardCommand := &commands.PlayRoadBuildingCardCommand{
		GameID:            playRoadBuildingCardRequest.GetGameID(),
		UserID:            playRoadBuildingCardRequest.GetUserID(),
		DevelopmentCardID: playRoadBuildingCardRequest.GetDevelopmentCardID(),
		PathIDs:           playRoadBuildingCardRequest.GetPathIDs(),
	}

	if err := c.commandBus.Execute(ctx, playRoadBuildingCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayYearOfPlentyCard(ctx context.Context, playYearOfPlentyCardRequest *pb_models.PlayYearOfPlentyCardRequest) (*emptypb.Empty, error) {
	playYearOfPlentyCardCommand := &commands.PlayYearOfPlentyCardCommand{
		GameID:                     playYearOfPlentyCardRequest.GetGameID(),
		UserID:                     playYearOfPlentyCardRequest.GetUserID(),
		DevelopmentCardID:          playYearOfPlentyCardRequest.GetDevelopmentCardID(),
		DemandingResourceCardTypes: playYearOfPlentyCardRequest.GetDemandingResourceCardTypes(),
	}

	if err := c.commandBus.Execute(ctx, playYearOfPlentyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayMonopolyCard(ctx context.Context, playMonopolyCardRequest *pb_models.PlayMonopolyCardRequest) (*emptypb.Empty, error) {
	playMonopolyCardCommand := &commands.PlayMonopolyCardCommand{
		GameID:                    playMonopolyCardRequest.GetGameID(),
		UserID:                    playMonopolyCardRequest.GetUserID(),
		DevelopmentCardID:         playMonopolyCardRequest.GetDevelopmentCardID(),
		DemandingResourceCardType: playMonopolyCardRequest.GetDemandingResourceCardType(),
	}

	if err := c.commandBus.Execute(ctx, playMonopolyCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (c catanServer) PlayVictoryPointCard(ctx context.Context, playVictoryPointCardRequest *pb_models.PlayVictoryPointCardRequest) (*emptypb.Empty, error) {
	playVictoryPointCardCommand := &commands.PlayVictoryPointCardCommand{
		GameID:            playVictoryPointCardRequest.GetGameID(),
		UserID:            playVictoryPointCardRequest.GetUserID(),
		DevelopmentCardID: playVictoryPointCardRequest.GetDevelopmentCardID(),
	}

	if err := c.commandBus.Execute(ctx, playVictoryPointCardCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
