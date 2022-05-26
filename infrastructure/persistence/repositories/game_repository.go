package repositories

import (
	"context"

	infrastructure_context "github.com/VulpesFerrilata/catan-service/infrastructure/context"
	"github.com/VulpesFerrilata/catan-service/infrastructure/persistence"
	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/VulpesFerrilata/catan-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewGameRepository() repositories.GameRepository {
	return &gameRepository{}
}

type gameRepository struct{}

func (g gameRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.Game, error) {
	gameEntity := new(entities.Game)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx = tx.First(gameEntity, id)
	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(persistence.ErrRecordNotFound)
	}
	if err := tx.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return gameEntity, nil
}

func (g gameRepository) IsExists(ctx context.Context, id uuid.UUID) (bool, error) {
	count := int64(0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	tx = tx.Model(&entities.Game{})
	tx = tx.Where("id = ?", id)
	tx = tx.Count(&count)
	if err := tx.Error; err != nil {
		return false, errors.WithStack(err)
	}

	return count > 0, nil
}

func (g gameRepository) Insert(ctx context.Context, gameEntity *entities.Game) error {
	gameEntity.Version = 1

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Create(gameEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (g gameRepository) Update(ctx context.Context, gameEntity *entities.Game) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Where("id = ?", gameEntity.ID)
	tx = tx.Where("version = ?", gameEntity.Version)

	gameEntity.Version++

	tx = tx.Updates(gameEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}
	if rowsAffected := tx.RowsAffected; rowsAffected == 0 {
		return errors.WithStack(persistence.ErrStaleObject)
	}

	return nil
}

func (g gameRepository) Save(ctx context.Context, gameEntity *entities.Game) error {
	isExists, err := g.IsExists(ctx, gameEntity.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if isExists {
		if err := g.Update(ctx, gameEntity); err != nil {
			return errors.WithStack(err)
		}
	} else {
		if err := g.Insert(ctx, gameEntity); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
