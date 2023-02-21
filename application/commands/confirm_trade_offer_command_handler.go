package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConfirmTradeOfferCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewConfirmTradeOfferCommandHandler(gameRepository repositories.GameRepository) *ConfirmTradeOfferCommandHandler {
	return &ConfirmTradeOfferCommandHandler{
		gameRepository: gameRepository,
	}
}

type ConfirmTradeOfferCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (c ConfirmTradeOfferCommandHandler) Handle(ctx context.Context, confirmTradeOfferCommand *ConfirmTradeOfferCommand) error {
	gameID, err := primitive.ObjectIDFromHex(confirmTradeOfferCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(confirmTradeOfferCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := c.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.ConfirmTradeOffer(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := c.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
