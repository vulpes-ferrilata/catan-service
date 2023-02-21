package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JoinGameCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewJoinGameCommandHandler(gameRepository repositories.GameRepository) *JoinGameCommandHandler {
	return &JoinGameCommandHandler{
		gameRepository: gameRepository,
	}
}

type JoinGameCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (j JoinGameCommandHandler) Handle(ctx context.Context, joinGameCommand *JoinGameCommand) error {
	gameID, err := primitive.ObjectIDFromHex(joinGameCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(joinGameCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := j.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.NewPlayer(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := j.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
