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
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPlayYearOfPlentyCardCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.PlayYearOfPlentyCard] {
	handler := &playYearOfPlentyCardCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.PlayYearOfPlentyCard](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type playYearOfPlentyCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p playYearOfPlentyCardCommandHandler) Handle(ctx context.Context, playYearOfPlentyCardCommand *commands.PlayYearOfPlentyCard) error {
	gameID, err := primitive.ObjectIDFromHex(playYearOfPlentyCardCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(playYearOfPlentyCardCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	developmentCardID, err := primitive.ObjectIDFromHex(playYearOfPlentyCardCommand.DevelopmentCardID)
	if err != nil {
		return errors.WithStack(err)
	}

	demandingResourceCardTypes, err := slices.Map(func(resourceCardType string) (models.ResourceCardType, error) {
		return models.NewResourceCardType(resourceCardType)
	}, playYearOfPlentyCardCommand.DemandingResourceCardTypes)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := p.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.PlayYearOfPlentyCard(userID, developmentCardID, demandingResourceCardTypes); err != nil {
		return errors.WithStack(err)
	}

	if err := p.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
