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

func NewGetGameByIDByUserIDQueryHandler(validate *validator.Validate, gameProjector projectors.GameProjector) query.QueryHandler[*queries.GetGameByIDByUserID, *models.Game] {
	handler := &getGameQueryHandler{
		gameProjector: gameProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.GetGameByIDByUserID, *models.Game](validate, handler)

	return validationWrapper
}

type getGameQueryHandler struct {
	gameProjector projectors.GameProjector
}

func (g getGameQueryHandler) Handle(ctx context.Context, getGameQuery *queries.GetGameByIDByUserID) (*models.Game, error) {
	gameID, err := primitive.ObjectIDFromHex(getGameQuery.GameID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(getGameQuery.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	game, err := g.gameProjector.GetByIDByUserID(ctx, gameID, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return game, nil
}
