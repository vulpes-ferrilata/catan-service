package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildSettlementAndRoadCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
	LandID string `validate:"required,objectid"`
	PathID string `validate:"required,objectid"`
}

func NewBuildSettlementAndRoadCommandHandler(gameRepository repositories.GameRepository) *BuildSettlementAndRoadCommandHandler {
	return &BuildSettlementAndRoadCommandHandler{
		gameRepository: gameRepository,
	}
}

type BuildSettlementAndRoadCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (b BuildSettlementAndRoadCommandHandler) Handle(ctx context.Context, buildSettlementAndRoadCommand *BuildSettlementAndRoadCommand) error {
	gameID, err := primitive.ObjectIDFromHex(buildSettlementAndRoadCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(buildSettlementAndRoadCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	landID, err := primitive.ObjectIDFromHex(buildSettlementAndRoadCommand.LandID)
	if err != nil {
		return errors.WithStack(err)
	}

	pathID, err := primitive.ObjectIDFromHex(buildSettlementAndRoadCommand.PathID)
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
