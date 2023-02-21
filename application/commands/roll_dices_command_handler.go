package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RollDicesCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewRollDicesCommandHandler(gameRepository repositories.GameRepository) *RollDicesCommandHandler {
	return &RollDicesCommandHandler{
		gameRepository: gameRepository,
	}
}

type RollDicesCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (r RollDicesCommandHandler) Handle(ctx context.Context, rollDicesCommand *RollDicesCommand) error {
	gameID, err := primitive.ObjectIDFromHex(rollDicesCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(rollDicesCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := r.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.RollDices(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := r.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
