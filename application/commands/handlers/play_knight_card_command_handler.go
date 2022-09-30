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

func NewPlayKnightCardCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.PlayKnightCard] {
	handler := &playKnightCardCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.PlayKnightCard](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type playKnightCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p playKnightCardCommandHandler) Handle(ctx context.Context, playKnightCardCommand *commands.PlayKnightCard) error {
	gameID, err := primitive.ObjectIDFromHex(playKnightCardCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(playKnightCardCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	terrainID, err := primitive.ObjectIDFromHex(playKnightCardCommand.TerrainID)
	if err != nil {
		return errors.WithStack(err)
	}

	var playerID primitive.ObjectID
	if playKnightCardCommand.PlayerID != "" {
		playerID, err = primitive.ObjectIDFromHex(playKnightCardCommand.PlayerID)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	game, err := p.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.PlayKnightCard(userID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	if err := p.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
