package repositories

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/google/uuid"
)

type GameRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Game, error)
	Save(ctx context.Context, gameEntity *entities.Game) error
}
