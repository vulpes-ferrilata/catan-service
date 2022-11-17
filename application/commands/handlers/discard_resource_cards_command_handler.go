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

func NewDiscardResourceCardsCommandHandler(validate *validator.Validate, db *mongo.Database, gameRepository repositories.GameRepository) command.CommandHandler[*commands.DiscardResourceCards] {
	handler := &discardResourceCardsCommandHandler{
		gameRepository: gameRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.DiscardResourceCards](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type discardResourceCardsCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (d discardResourceCardsCommandHandler) Handle(ctx context.Context, discardResourceCardsCommand *commands.DiscardResourceCards) error {
	gameID, err := primitive.ObjectIDFromHex(discardResourceCardsCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(discardResourceCardsCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	resourceCardIDs, err := slices.Map(func(resoureCardID string) (primitive.ObjectID, error) {
		return primitive.ObjectIDFromHex(resoureCardID)
	}, discardResourceCardsCommand.ResourceCardIDs)

	game, err := d.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.DiscardResourceCards(userID, resourceCardIDs); err != nil {
		return errors.WithStack(err)
	}

	if err := d.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
