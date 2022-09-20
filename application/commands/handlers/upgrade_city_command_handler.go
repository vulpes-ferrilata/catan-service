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

func NewUpgradeCityCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.UpgradeCityCommand] {
	handler := &upgradeCityCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.UpgradeCityCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type upgradeCityCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (u upgradeCityCommandHandler) Handle(ctx context.Context, upgradeCity *commands.UpgradeCityCommand) error {
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
