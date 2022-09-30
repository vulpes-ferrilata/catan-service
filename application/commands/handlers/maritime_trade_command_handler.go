package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMaritimeTradeCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.MaritimeTrade] {
	handler := &maritimeTradeCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.MaritimeTrade](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type maritimeTradeCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (m maritimeTradeCommandHandler) Handle(ctx context.Context, maritimeTradeCommand *commands.MaritimeTrade) error {
	gameID, err := primitive.ObjectIDFromHex(maritimeTradeCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(maritimeTradeCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	demandingResourceCardType, err := models.NewResourceCardType(maritimeTradeCommand.ResourceCardType)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := m.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.MaritimeTrade(userID, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	if err := m.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
