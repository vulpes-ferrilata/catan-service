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

func NewMoveRobberCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.MoveRobber] {
	handler := &moveRobberCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.MoveRobber](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type moveRobberCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (m moveRobberCommandHandler) Handle(ctx context.Context, moveRobberCommand *commands.MoveRobber) error {
	gameID, err := primitive.ObjectIDFromHex(moveRobberCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(moveRobberCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	terrainID, err := primitive.ObjectIDFromHex(moveRobberCommand.TerrainID)
	if err != nil {
		return errors.WithStack(err)
	}

	var playerID primitive.ObjectID
	if moveRobberCommand.PlayerID != "" {
		playerID, err = primitive.ObjectIDFromHex(moveRobberCommand.PlayerID)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	game, err := m.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.MoveRobber(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	if err := m.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
