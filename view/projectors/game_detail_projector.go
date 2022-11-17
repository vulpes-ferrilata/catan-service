package projectors

import (
	"context"

	"github.com/vulpes-ferrilata/catan-service/view/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameDetailProjector interface {
	GetByIDByUserID(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*models.GameDetail, error)
}
