package repositories

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/google/uuid"
)

type AchievementRepository interface {
	FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*entities.Achievement, error)
	Save(ctx context.Context, achievementEntity *entities.Achievement) error
}
