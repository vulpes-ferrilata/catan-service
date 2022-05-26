package handlers

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/application/commands"
	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/domain/services"
	"github.com/VulpesFerrilata/catan-service/infrastructure/dig/results"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewJoinGameCommandHandler(gameService services.GameService,
	playerService services.PlayerService) results.CommandHandlerResult {
	commandHandler := &joinGameCommandHandler{
		gameService:   gameService,
		playerService: playerService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type joinGameCommandHandler struct {
	gameService   services.GameService
	playerService services.PlayerService
}

func (c joinGameCommandHandler) GetCommand() interface{} {
	return new(commands.JoinGameCommand)
}

func (c joinGameCommandHandler) Handle(ctx context.Context, command interface{}) error {
	joinGameCommand := command.(*commands.JoinGameCommand)

	gameID, err := uuid.Parse(joinGameCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := uuid.Parse(joinGameCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := c.gameService.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if game.GetStatus() != models.Waiting {
		return nil
	}

	player, err := c.playerService.NewPlayer(ctx, userID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.AddPlayer(player); err != nil {
		return errors.WithStack(err)
	}

	if err := c.gameService.Save(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
