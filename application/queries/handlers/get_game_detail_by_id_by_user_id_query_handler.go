package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/application/queries"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/cqrs/query/wrappers"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/catan-service/view/projectors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewGetGameDetailByIDByUserIDQueryHandler(validate *validator.Validate, gameDetailProjector projectors.GameDetailProjector) query.QueryHandler[*queries.GetGameDetailByIDByUserID, *models.GameDetail] {
	handler := &getGameDetailQueryHandler{
		gameDetailProjector: gameDetailProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.GetGameDetailByIDByUserID, *models.GameDetail](validate, handler)

	return validationWrapper
}

type getGameDetailQueryHandler struct {
	gameDetailProjector projectors.GameDetailProjector
}

func (g getGameDetailQueryHandler) Handle(ctx context.Context, getGameDetailByIDByUserIDQuery *queries.GetGameDetailByIDByUserID) (*models.GameDetail, error) {
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
