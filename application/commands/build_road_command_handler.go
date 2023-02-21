package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildRoadCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
	PathID string `validate:"required,objectid"`
}

func NewBuildRoadCommandHandler(gameRepository repositories.GameRepository) *BuildRoadCommandHandler {
	return &BuildRoadCommandHandler{
		gameRepository: gameRepository,
	}
}

type BuildRoadCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (b BuildRoadCommandHandler) Handle(ctx context.Context, buildRoadCommand *BuildRoadCommand) error {
	gameID, err := primitive.ObjectIDFromHex(buildRoadCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(buildRoadCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	pathID, err := primitive.ObjectIDFromHex(buildRoadCommand.PathID)
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
