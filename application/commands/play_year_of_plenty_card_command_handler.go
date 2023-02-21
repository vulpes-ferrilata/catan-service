package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayYearOfPlentyCardCommand struct {
	GameID                     string   `validate:"required,objectid"`
	UserID                     string   `validate:"required,objectid"`
	DevelopmentCardID          string   `validate:"required,objectid"`
	DemandingResourceCardTypes []string `validate:"required,min=1,max=2"`
}

func NewPlayYearOfPlentyCardCommandHandler(gameRepository repositories.GameRepository) *PlayYearOfPlentyCardCommandHandler {
	return &PlayYearOfPlentyCardCommandHandler{
		gameRepository: gameRepository,
	}
}

type PlayYearOfPlentyCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p PlayYearOfPlentyCardCommandHandler) Handle(ctx context.Context, playYearOfPlentyCardCommand *PlayYearOfPlentyCardCommand) error {
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
	}, playYearOfPlentyCardCommand.DemandingResourceCardTypes...)
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
