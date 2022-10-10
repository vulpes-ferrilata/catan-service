package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSendTradeOfferCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.SendTradeOffer] {
	handler := &sendTradeOfferCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.SendTradeOffer](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type sendTradeOfferCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (o sendTradeOfferCommandHandler) Handle(ctx context.Context, sendTradeOfferCommand *commands.SendTradeOffer) error {
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

	game, err := o.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.SendTradeOffer(userID, playerID); err != nil {
		return errors.WithStack(err)
	}

	if err := o.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
