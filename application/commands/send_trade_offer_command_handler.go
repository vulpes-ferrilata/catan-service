package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SendTradeOfferCommand struct {
	GameID   string `validate:"required,objectid"`
	UserID   string `validate:"required,objectid"`
	PlayerID string `validate:"required,objectid"`
}

func NewSendTradeOfferCommandHandler(gameRepository repositories.GameRepository) *SendTradeOfferCommandHandler {
	return &SendTradeOfferCommandHandler{
		gameRepository: gameRepository,
	}
}

type SendTradeOfferCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (s SendTradeOfferCommandHandler) Handle(ctx context.Context, sendTradeOfferCommand *SendTradeOfferCommand) error {
	gameID, err := primitive.ObjectIDFromHex(sendTradeOfferCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(sendTradeOfferCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	playerID, err := primitive.ObjectIDFromHex(sendTradeOfferCommand.PlayerID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := s.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.SendTradeOffer(userID, playerID); err != nil {
		return errors.WithStack(err)
	}

	if err := s.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
