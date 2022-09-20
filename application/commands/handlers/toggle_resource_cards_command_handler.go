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

func NewToggleResourceCardsCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.ToggleResourceCardsCommand] {
	handler := &toggleResourceCardsCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.ToggleResourceCardsCommand](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type toggleResourceCardsCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (t toggleResourceCardsCommandHandler) Handle(ctx context.Context, toggleResourceCardsCommand *commands.ToggleResourceCardsCommand) error {
	gameID, err := primitive.ObjectIDFromHex(toggleResourceCardsCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(toggleResourceCardsCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	resourceCardIDs, err := slices.Map(func(resoureCardID string) (primitive.ObjectID, error) {
		return primitive.ObjectIDFromHex(resoureCardID)
	}, toggleResourceCardsCommand.ResourceCardIDs)

	game, err := t.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.ToggleResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	if err := t.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
