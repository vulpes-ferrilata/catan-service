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

func NewFindGamesByUserIDQueryHandler(validate *validator.Validate, gameProjector projectors.GameProjector) query.QueryHandler[*queries.FindGamesByUserID, []*models.Game] {
	handler := &findGamesByUserIDQueryHandler{
		gameProjector: gameProjector,
	}
	validationWrapper := wrappers.NewValidationWrapper[*queries.FindGamesByUserID, []*models.Game](validate, handler)

	return validationWrapper
}

type findGamesByUserIDQueryHandler struct {
	gameProjector projectors.GameProjector
}

func (f findGamesByUserIDQueryHandler) Handle(ctx context.Context, findGamesByUserIDQuery *queries.FindGamesByUserID) ([]*models.Game, error) {
	userID, err := primitive.ObjectIDFromHex(findGamesByUserIDQuery.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	games, err := f.gameProjector.FindByUserID(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return games, nil
}
