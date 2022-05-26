package repositories

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/google/uuid"
)

type ResourceCardRepository interface {
	FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*entities.ResourceCard, error)
	Save(ctx context.Context, achievementEntity *entities.ResourceCard) error
}
