package queries

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetGameDetailByIDByUserIDQuery struct {
	GameID string `validate:"required,objectid"`
	UserID string `validate:"required,objectid"`
}

func NewGetGameDetailByIDByUserIDQueryHandler(gameDetailProjector projectors.GameDetailProjector) *GetGameDetailByIDByUserIDQueryHandler {
	return &GetGameDetailByIDByUserIDQueryHandler{
		gameDetailProjector: gameDetailProjector,
	}
}

type GetGameDetailByIDByUserIDQueryHandler struct {
	gameDetailProjector projectors.GameDetailProjector
}

func (g GetGameDetailByIDByUserIDQueryHandler) Handle(ctx context.Context, getGameDetailByIDByUserIDQuery *GetGameDetailByIDByUserIDQuery) (*models.GameDetail, error) {
	gameID, err := primitive.ObjectIDFromHex(getGameDetailByIDByUserIDQuery.GameID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(getGameDetailByIDByUserIDQuery.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gameDetail, err := g.gameDetailProjector.GetByIDByUserID(ctx, gameID, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gameDetail, nil
}
