package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/application/commands"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/command/wrappers"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPlayRoadBuildingCardCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.PlayRoadBuildingCardCommand] {
	handler := &playRoadBuildingCardCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.PlayRoadBuildingCardCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type playRoadBuildingCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p playRoadBuildingCardCommandHandler) Handle(ctx context.Context, playRoadBuildingCard *commands.PlayRoadBuildingCardCommand) error {
	gameID, err := primitive.ObjectIDFromHex(playRoadBuildingCard.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(playRoadBuildingCard.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	pathIDs, err := slices.Map(func(pathID string) (primitive.ObjectID, error) {
		return primitive.ObjectIDFromHex(pathID)
	}, playRoadBuildingCard.PathIDs)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := p.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.PlayRoadBuildingCard(userID, pathIDs); err != nil {
		return errors.WithStack(err)
	}

	if err := p.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
