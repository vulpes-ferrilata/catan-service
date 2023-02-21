package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MoveRobberCommand struct {
	GameID    string `validate:"required,objectid"`
	UserID    string `validate:"required,objectid"`
	TerrainID string `validate:"required,objectid"`
	PlayerID  string `validate:"omitempty,objectid"`
}

func NewMoveRobberCommandHandler(gameRepository repositories.GameRepository) *MoveRobberCommandHandler {
	return &MoveRobberCommandHandler{
		gameRepository: gameRepository,
	}
}

type MoveRobberCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (m MoveRobberCommandHandler) Handle(ctx context.Context, moveRobberCommand *MoveRobberCommand) error {
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
