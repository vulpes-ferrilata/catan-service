package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildSettlementCommand struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
	LandID string `validate:"required,objectid"`
}

func NewBuildSettlementCommandHandler(gameRepository repositories.GameRepository) *BuildSettlementCommandHandler {
	return &BuildSettlementCommandHandler{
		gameRepository: gameRepository,
	}
}

type BuildSettlementCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (b BuildSettlementCommandHandler) Handle(ctx context.Context, buildSettlementCommand *BuildSettlementCommand) error {
	gameID, err := primitive.ObjectIDFromHex(buildSettlementCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(buildSettlementCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	landID, err := primitive.ObjectIDFromHex(buildSettlementCommand.LandID)
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
