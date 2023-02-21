package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayKnightCardCommand struct {
	GameID            string `validate:"required,objectid"`
	UserID            string `validate:"required,objectid"`
	DevelopmentCardID string `validate:"required,objectid"`
	TerrainID         string `validate:"required,objectid"`
	PlayerID          string `validate:"omitempty,objectid"`
}

func NewPlayKnightCardCommandHandler(gameRepository repositories.GameRepository) *PlayKnightCardCommandHandler {
	return &PlayKnightCardCommandHandler{
		gameRepository: gameRepository,
	}
}

type PlayKnightCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p PlayKnightCardCommandHandler) Handle(ctx context.Context, playKnightCardCommand *PlayKnightCardCommand) error {
	gameID, err := primitive.ObjectIDFromHex(playKnightCardCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(playKnightCardCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	developmentCardID, err := primitive.ObjectIDFromHex(playKnightCardCommand.DevelopmentCardID)
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

	if err := game.PlayKnightCard(userID, developmentCardID, terrainID, playerID); err != nil {
		return errors.WithStack(err)
	}

	if err := p.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
