package handlers

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/application/commands"
	"github.com/VulpesFerrilata/catan-service/domain/services"
	"github.com/VulpesFerrilata/catan-service/infrastructure/dig/results"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewCreateGameCommandHandler(gameService services.GameService,
	playerService services.PlayerService) results.CommandHandlerResult {
	commandHandler := &createGameCommandHandler{
		gameService: gameService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type createGameCommandHandler struct {
	gameService services.GameService
}

func (c createGameCommandHandler) GetCommand() interface{} {
	return new(commands.CreateGameCommand)
}

func (c createGameCommandHandler) Handle(ctx context.Context, command interface{}) error {
	createGameCommand := command.(*commands.CreateGameCommand)

	gameID, err := uuid.Parse(createGameCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := uuid.Parse(createGameCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := c.gameService.NewGame(ctx, gameID, userID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.gameService.Save(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
