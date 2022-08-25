package projectors

import (
	"context"

	"github.com/vulpes-ferrilata/catan-service/view/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameProjector interface {
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Game, error)
	GetByIDByUserID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*models.Game, error)
}
