package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiscardResourceCardsCommand struct {
	GameID          string   `validate:"required,objectid"`
	UserID          string   `validate:"required,objectid"`
	ResourceCardIDs []string `validate:"required,unique"`
}

func NewDiscardResourceCardsCommandHandler(gameRepository repositories.GameRepository) *DiscardResourceCardsCommandHandler {
	return &DiscardResourceCardsCommandHandler{
		gameRepository: gameRepository,
	}
}

type DiscardResourceCardsCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (d DiscardResourceCardsCommandHandler) Handle(ctx context.Context, discardResourceCardsCommand *DiscardResourceCardsCommand) error {
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
	}, discardResourceCardsCommand.ResourceCardIDs...)

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
