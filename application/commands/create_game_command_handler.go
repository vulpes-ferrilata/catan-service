package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateGameCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewCreateGameCommandHandler(gameRepository repositories.GameRepository) *CreateGameCommandHandler {
	return &CreateGameCommandHandler{
		gameRepository: gameRepository,
	}
}

type CreateGameCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (c CreateGameCommandHandler) Handle(ctx context.Context, createGameCommand *CreateGameCommand) error {
	gameID, err := primitive.ObjectIDFromHex(createGameCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(createGameCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	game := models.GameBuilder{}.
		SetID(gameID).
		SetStatus(models.Waiting).
		SetPhase(models.ResourceConsumption).
		SetTurn(1).
		Create()

	if err := game.NewPlayer(userID); err != nil {
		return errors.WithStack(err)
	}

	if err := c.gameRepository.Insert(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
