package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EndTurnCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewEndTurnCommandHandler(gameRepository repositories.GameRepository) *EndTurnCommandHandler {
	return &EndTurnCommandHandler{
		gameRepository: gameRepository,
	}
}

type EndTurnCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (e EndTurnCommandHandler) Handle(ctx context.Context, endTurnCommand *EndTurnCommand) error {
	gameID, err := primitive.ObjectIDFromHex(endTurnCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(endTurnCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := e.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.EndTurn(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := e.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
