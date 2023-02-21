package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToggleResourceCardsCommand struct {
	GameID          string   `validate:"required,objectid"`
	UserID          string   `validate:"required,objectid"`
	ResourceCardIDs []string `validate:"required,unique"`
}

func NewToggleResourceCardsCommandHandler(gameRepository repositories.GameRepository) *ToggleResourceCardsCommandHandler {
	return &ToggleResourceCardsCommandHandler{
		gameRepository: gameRepository,
	}
}

type ToggleResourceCardsCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (t ToggleResourceCardsCommandHandler) Handle(ctx context.Context, toggleResourceCardsCommand *ToggleResourceCardsCommand) error {
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
	}, toggleResourceCardsCommand.ResourceCardIDs...)

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
