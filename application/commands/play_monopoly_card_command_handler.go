package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayMonopolyCardCommand struct {
	GameID                    string `validate:"required,objectid"`
	UserID                    string `validate:"required,objectid"`
	DevelopmentCardID         string `validate:"required,objectid"`
	DemandingResourceCardType string `validate:"required"`
}

func NewPlayMonopolyCardCommandHandler(gameRepository repositories.GameRepository) *PlayMonopolyCardCommandHandler {
	return &PlayMonopolyCardCommandHandler{
		gameRepository: gameRepository,
	}
}

type PlayMonopolyCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p PlayMonopolyCardCommandHandler) Handle(ctx context.Context, playMonopolyCardCommand *PlayMonopolyCardCommand) error {
	gameID, err := primitive.ObjectIDFromHex(playMonopolyCardCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(playMonopolyCardCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	developmentCardID, err := primitive.ObjectIDFromHex(playMonopolyCardCommand.DevelopmentCardID)
	if err != nil {
		return errors.WithStack(err)
	}

	demandingResourceCardType, err := models.NewResourceCardType(playMonopolyCardCommand.DemandingResourceCardType)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := p.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.PlayMonopolyCard(userID, developmentCardID, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	if err := p.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
