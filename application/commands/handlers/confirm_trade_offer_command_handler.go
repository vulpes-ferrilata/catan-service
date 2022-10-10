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

func NewConfirmTradeOfferCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.ConfirmTradeOffer] {
	handler := &confirmTradeOfferCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.ConfirmTradeOffer](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type confirmTradeOfferCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (c confirmTradeOfferCommandHandler) Handle(ctx context.Context, confirmTradeOfferCommand *commands.ConfirmTradeOffer) error {
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
