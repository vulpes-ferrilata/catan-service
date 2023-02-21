package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayVictoryPointCardCommand struct {
	GameID            string `validate:"required,objectid"`
	UserID            string `validate:"required,objectid"`
	DevelopmentCardID string `validate:"required,objectid"`
}

func NewPlayVictoryPointCardCommandHandler(gameRepository repositories.GameRepository) *PlayVictoryPointCardCommandHandler {
	return &PlayVictoryPointCardCommandHandler{
		gameRepository: gameRepository,
	}
}

type PlayVictoryPointCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p PlayVictoryPointCardCommandHandler) Handle(ctx context.Context, playVictoryPointCardCommand *PlayVictoryPointCardCommand) error {
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
