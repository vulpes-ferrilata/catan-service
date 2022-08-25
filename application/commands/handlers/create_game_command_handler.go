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

func NewCreateGameCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.CreateGameCommand] {
	handler := &createGameCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.CreateGameCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper[*commands.CreateGameCommand](validate, transactionWrapper)

	return validationWrapper
}

type createGameCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (c createGameCommandHandler) Handle(ctx context.Context, createGameCommand *commands.CreateGameCommand) error {
	gameID, err := primitive.ObjectIDFromHex(createGameCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(createGameCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game := models.NewGameBuilder().
		SetID(gameID).
		SetStatus(models.Waiting).
		SetTurn(1).
		SetIsRolledDices(false).
		Create()

	if err := game.NewPlayer(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := c.gameRepository.Insert(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
