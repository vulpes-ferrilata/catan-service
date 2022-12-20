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

func NewPlayVictoryPointCardCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.PlayVictoryPointCard] {
	handler := &playVictoryPointCardCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.PlayVictoryPointCard](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type playVictoryPointCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p playVictoryPointCardCommandHandler) Handle(ctx context.Context, playVictoryPointCardCommand *commands.PlayVictoryPointCard) error {
	gameID, err := primitive.ObjectIDFromHex(playVictoryPointCardCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(playVictoryPointCardCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	developmentCardID, err := primitive.ObjectIDFromHex(playVictoryPointCardCommand.DevelopmentCardID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := p.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.PlayVictoryPointCard(userID, developmentCardID); err != nil {
		return errors.WithStack(err)
	}

	if err := p.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
