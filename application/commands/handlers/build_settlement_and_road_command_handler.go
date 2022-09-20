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

func NewBuildSettlementAndRoadCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.BuildSettlementAndRoadCommand] {
	handler := &buildSettlementAndRoadCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.BuildSettlementAndRoadCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type buildSettlementAndRoadCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (b buildSettlementAndRoadCommandHandler) Handle(ctx context.Context, buildSettlementAndRoad *commands.BuildSettlementAndRoadCommand) error {
	gameID, err := primitive.ObjectIDFromHex(buildSettlementAndRoad.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(buildSettlementAndRoad.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	landID, err := primitive.ObjectIDFromHex(buildSettlementAndRoad.LandID)
	if err != nil {
		return errors.WithStack(err)
	}

	pathID, err := primitive.ObjectIDFromHex(buildSettlementAndRoad.PathID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := b.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.BuildSettlementAndRoad(userID, landID, pathID); err != nil {
		return errors.WithStack(err)
	}

	if err := b.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
