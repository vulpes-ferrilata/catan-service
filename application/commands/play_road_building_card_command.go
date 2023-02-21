package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/domain/repositories"
	"github.com/vulpes-ferrilata/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayRoadBuildingCardCommand struct {
	GameID            string   `validate:"required,objectid"`
	UserID            string   `validate:"required,objectid"`
	DevelopmentCardID string   `validate:"required,objectid"`
	PathIDs           []string `validate:"required,min=1,max=2,unique"`
}

func NewPlayRoadBuildingCardCommandHandler(gameRepository repositories.GameRepository) *PlayRoadBuildingCardCommandHandler {
	return &PlayRoadBuildingCardCommandHandler{
		gameRepository: gameRepository,
	}
}

type PlayRoadBuildingCardCommandHandler struct {
	gameRepository repositories.GameRepository
}

func (p PlayRoadBuildingCardCommandHandler) Handle(ctx context.Context, playRoadBuildingCardCommand *PlayRoadBuildingCardCommand) error {
	gameID, err := primitive.ObjectIDFromHex(playRoadBuildingCardCommand.GameID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(playRoadBuildingCardCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	developmentCardID, err := primitive.ObjectIDFromHex(playRoadBuildingCardCommand.DevelopmentCardID)
	if err != nil {
		return errors.WithStack(err)
	}

	pathIDs, err := slices.Map(func(pathID string) (primitive.ObjectID, error) {
		return primitive.ObjectIDFromHex(pathID)
	}, playRoadBuildingCardCommand.PathIDs...)
	if err != nil {
		return errors.WithStack(err)
	}

	game, err := p.gameRepository.GetByID(ctx, gameID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := game.PlayRoadBuildingCard(userID, developmentCardID, pathIDs); err != nil {
		return errors.WithStack(err)
	}

	if err := p.gameRepository.Update(ctx, game); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
