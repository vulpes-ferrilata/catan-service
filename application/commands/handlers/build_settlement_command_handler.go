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

func NewBuildSettlementCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.BuildSettlementCommand] {
	handler := &buildSettlementCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.BuildSettlementCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type buildSettlementCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (b buildSettlementCommandHandler) Handle(ctx context.Context, buildSettlement *commands.BuildSettlementCommand) error {
	gameID, err := primitive.ObjectIDFromHex(buildSettlement.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(buildSettlement.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	landID, err := primitive.ObjectIDFromHex(buildSettlement.LandID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := b.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.BuildSettlement(userID, landID); err != nil {
		return errors.WithStack(err)
	}

	if err := b.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
