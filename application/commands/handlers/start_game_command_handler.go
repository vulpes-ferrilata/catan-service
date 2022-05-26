package handlers

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/application/commands"
	"github.com/VulpesFerrilata/catan-service/domain/services"
	"github.com/VulpesFerrilata/catan-service/infrastructure/dig/results"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewStartGameCommandHandler(gameService services.GameService,
	playerService services.PlayerService) results.CommandHandlerResult {
	commandHandler := &startGameCommandHandler{
		gameService: gameService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type startGameCommandHandler struct {
	gameService services.GameService
}

func (c startGameCommandHandler) GetCommand() interface{} {
	return new(commands.StartGameCommand)
}

func (c startGameCommandHandler) Handle(ctx context.Context, command interface{}) error {
	startGameCommand := command.(*commands.StartGameCommand)

	gameID, err := uuid.Parse(startGameCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := uuid.Parse(startGameCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := c.gameService.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.Start(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := c.gameService.Save(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
