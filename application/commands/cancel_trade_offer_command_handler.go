package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CancelTradeOfferCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewCancelTradeOfferCommandHandler(gameRepository repositories.GameRepository) *CancelTradeOfferCommandHandler {
	return &CancelTradeOfferCommandHandler{
		gameRepository: gameRepository,
	}
}

type CancelTradeOfferCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (c CancelTradeOfferCommandHandler) Handle(ctx context.Context, cancelTradeOfferCommand *CancelTradeOfferCommand) error {
	gameID, err := primitive.ObjectIDFromHex(cancelTradeOfferCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(cancelTradeOfferCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := c.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.CancelTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := c.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
