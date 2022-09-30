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

func NewOfferTradingCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.OfferTrading] {
	handler := &offerTradingCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.OfferTrading](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type offerTradingCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (o offerTradingCommandHandler) Handle(ctx context.Context, offerTradingCommand *commands.OfferTrading) error {
	gameID, err := primitive.ObjectIDFromHex(offerTradingCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(offerTradingCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	playerID, err := primitive.ObjectIDFromHex(offerTradingCommand.PlayerID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := o.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.OfferTrading(userID, playerID); err != nil {
		return errors.WithStack(err)
	}

	if err := o.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
