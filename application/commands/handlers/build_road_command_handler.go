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

func NewBuildRoadCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.BuildRoadCommand] {
	handler := &buildRoadCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.BuildRoadCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type buildRoadCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (b buildRoadCommandHandler) Handle(ctx context.Context, buildRoad *commands.BuildRoadCommand) error {
	gameID, err := primitive.ObjectIDFromHex(buildRoad.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(buildRoad.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	pathID, err := primitive.ObjectIDFromHex(buildRoad.PathID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := b.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.BuildRoad(userID, pathID); err != nil {
		return errors.WithStack(err)
	}

	if err := b.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
