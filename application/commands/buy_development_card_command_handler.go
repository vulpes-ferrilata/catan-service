package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuyDevelopmentCardCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewBuyDevelopmentCardCommandHandler(gameRepository repositories.GameRepository) *BuyDevelopmentCardCommandHandler {
	return &BuyDevelopmentCardCommandHandler{
		gameRepository: gameRepository,
	}
}

type BuyDevelopmentCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (b BuyDevelopmentCardCommandHandler) Handle(ctx context.Context, buyDevelopmentCardCommand *BuyDevelopmentCardCommand) error {
	gameID, err := primitive.ObjectIDFromHex(buyDevelopmentCardCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(buyDevelopmentCardCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := b.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.BuyDevelopmentCard(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := b.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
