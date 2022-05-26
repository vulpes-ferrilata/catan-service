package repositories

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/google/uuid"
)

type PlayerRepository interface {
	FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*entities.Player, error)
	Save(ctx context.Context, playerEntity *entities.Player) error
}
