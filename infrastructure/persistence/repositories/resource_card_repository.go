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

func NewResourceCardRepository() repositories.ResourceCardRepository {
	return &resourceCardRepository{}
}

type resourceCardRepository struct{}

func (r resourceCardRepository) FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*entities.ResourceCard, error) {
	resourceCardEntities := make([]*entities.ResourceCard, 0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx = tx.Where("game_id = ?", gameID)
	tx = tx.Where("player_id = ?", playerID)
	tx = tx.Find(resourceCardEntities)
	if err := tx.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return resourceCardEntities, nil
}

func (r resourceCardRepository) IsExists(ctx context.Context, id uuid.UUID) (bool, error) {
	count := int64(0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	tx = tx.Model(&entities.ResourceCard{})
	tx = tx.Where("id = ?", id)
	tx = tx.Count(&count)
	if err := tx.Error; err != nil {
		return false, errors.WithStack(err)
	}

	return count > 0, nil
}

func (r resourceCardRepository) Insert(ctx context.Context, resourceCardEntity *entities.ResourceCard) error {
	resourceCardEntity.Version = 1

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Create(resourceCardEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r resourceCardRepository) Update(ctx context.Context, resourceCardEntity *entities.ResourceCard) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Where("id = ?", resourceCardEntity.ID)
	tx = tx.Where("version = ?", resourceCardEntity.Version)

	resourceCardEntity.Version++

	tx = tx.Updates(resourceCardEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}
	if rowsAffected := tx.RowsAffected; rowsAffected == 0 {
		return errors.WithStack(persistence.ErrStaleObject)
	}

	return nil
}

func (r resourceCardRepository) Save(ctx context.Context, resourceCardEntity *entities.ResourceCard) error {
	isExists, err := r.IsExists(ctx, resourceCardEntity.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if isExists {
		if err := r.Update(ctx, resourceCardEntity); err != nil {
			return errors.WithStack(err)
		}
	} else {
		if err := r.Insert(ctx, resourceCardEntity); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
