package repositories

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/google/uuid"
)

type DiceRepository interface {
	FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*entities.Dice, error)
	Save(ctx context.Context, diceEntity *entities.Dice) error
}
