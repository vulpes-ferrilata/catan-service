package repositories

import (
	"context"

	"github.com/vulpes-ferrilata/catan-service/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Game, error)
	Insert(ctx context.Context, game *models.Game) error
	Update(ctx context.Context, game *models.Game) error
}
