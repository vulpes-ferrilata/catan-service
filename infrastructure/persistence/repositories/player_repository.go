package repositories

import (
	"context"

	infrastructure_context "github.com/VulpesFerrilata/catan-service/infrastructure/context"
	"github.com/VulpesFerrilata/catan-service/infrastructure/persistence"
	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/VulpesFerrilata/catan-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewPlayerRepository() repositories.PlayerRepository {
	return &playerRepository{}
}

type playerRepository struct{}

func (p playerRepository) FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*entities.Player, error) {
	playerEntities := make([]*entities.Player, 0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx = tx.Where("game_id = ?", gameID)
	tx = tx.Find(playerEntities)
	if err := tx.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return playerEntities, nil
}

func (p playerRepository) IsExists(ctx context.Context, id uuid.UUID) (bool, error) {
	count := int64(0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	tx = tx.Model(&entities.Player{})
	tx = tx.Where("id = ?", id)
	tx = tx.Count(&count)
	if err := tx.Error; err != nil {
		return false, errors.WithStack(err)
	}

	return count > 0, nil
}

func (p playerRepository) Insert(ctx context.Context, playerEntity *entities.Player) error {
	playerEntity.Version = 1

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Create(playerEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p playerRepository) Update(ctx context.Context, playerEntity *entities.Player) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Where("id = ?", playerEntity.ID)
	tx = tx.Where("version = ?", playerEntity.Version)

	playerEntity.Version++

	tx = tx.Updates(playerEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}
	if rowsAffected := tx.RowsAffected; rowsAffected == 0 {
		return errors.WithStack(persistence.ErrStaleObject)
	}

	return nil
}

func (p playerRepository) Save(ctx context.Context, playerEntity *entities.Player) error {
	isExists, err := p.IsExists(ctx, playerEntity.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if isExists {
		if err := p.Update(ctx, playerEntity); err != nil {
			return errors.WithStack(err)
		}
	} else {
		if err := p.Insert(ctx, playerEntity); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
