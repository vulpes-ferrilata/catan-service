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
	moveRobberCommandHandler command.CommandHandler[*commands.MoveRobberCommand]) catan.CatanServer {
	return &catanServer{
		findGamesByUserIDQueryHandler:        findGamesByUserIDQueryHandler,
		getGameByIDByUserIDQueryHandler:      getGameByIDByUserIDQueryHandler,
		createGameCommandHandler:             createGameCommandHandler,
		joinGameCommandHandler:               joinGameCommandHandler,
		startGameCommandHandler:              startGameCommandHandler,
		buildSettlementAndRoadCommandHandler: buildSettlementAndRoadCommandHandler,
		rollDicesCommandHandler:              rollDicesCommandHandler,
		moveRobberCommandHandler:             moveRobberCommandHandler,
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
