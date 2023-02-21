package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MaritimeTradeCommand struct {
	GameID                    string `validate:"required,objectid"`
	UserID                    string `validate:"required,objectid"`
	ResourceCardType          string `validate:"required"`
	DemandingResourceCardType string `validate:"required"`
}

func NewMaritimeTradeCommandHandler(gameRepository repositories.GameRepository) *MaritimeTradeCommandHandler {
	return &MaritimeTradeCommandHandler{
		gameRepository: gameRepository,
	}
}

type MaritimeTradeCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (m MaritimeTradeCommandHandler) Handle(ctx context.Context, maritimeTradeCommand *MaritimeTradeCommand) error {
	gameID, err := primitive.ObjectIDFromHex(maritimeTradeCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(maritimeTradeCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	resourceCardType, err := models.NewResourceCardType(maritimeTradeCommand.ResourceCardType)
	if err != nil {
		return errors.WithStack(err)
	}

	demandingResourceCardType, err := models.NewResourceCardType(maritimeTradeCommand.DemandingResourceCardType)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := m.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.MaritimeTrade(userID, resourceCardType, demandingResourceCardType); err != nil {
		return errors.WithStack(err)
	}

	if err := m.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
