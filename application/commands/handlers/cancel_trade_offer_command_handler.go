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

func NewCancelTradeOfferCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.CancelTradeOffer] {
	handler := &cancelTradeOfferCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.CancelTradeOffer](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type cancelTradeOfferCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (c cancelTradeOfferCommandHandler) Handle(ctx context.Context, cancelTradeOfferCommand *commands.CancelTradeOffer) error {
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
