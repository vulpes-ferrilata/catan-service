package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StartGameCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewStartGameCommandHandler(gameRepository repositories.GameRepository) *StartGameCommandHandler {
	return &StartGameCommandHandler{
		gameRepository: gameRepository,
	}
}

type StartGameCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (s StartGameCommandHandler) Handle(ctx context.Context, startGameCommand *StartGameCommand) error {
	gameID, err := primitive.ObjectIDFromHex(startGameCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(startGameCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := s.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.Start(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := s.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
