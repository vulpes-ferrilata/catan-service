package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpgradeCityCommand struct {
	GameID         string `validate:"required,objectid"`
	UserID         string `validate:"required,objectid"`
	ConstructionID string `validate:"required,objectid"`
}

func NewUpgradeCityCommandHandler(gameRepository repositories.GameRepository) *UpgradeCityCommandHandler {
	return &UpgradeCityCommandHandler{
		gameRepository: gameRepository,
	}
}

type UpgradeCityCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (u UpgradeCityCommandHandler) Handle(ctx context.Context, upgradeCity *UpgradeCityCommand) error {
	gameID, err := primitive.ObjectIDFromHex(upgradeCity.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(upgradeCity.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	constructionID, err := primitive.ObjectIDFromHex(upgradeCity.ConstructionID)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := u.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.UpgradeCity(userID, constructionID); err != nil {
		return errors.WithStack(err)
	}

	if err := u.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
