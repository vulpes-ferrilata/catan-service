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

func NewAchievementRepository() repositories.AchievementRepository {
	return &achievementRepository{}
}

type achievementRepository struct{}

func (a achievementRepository) FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*entities.Achievement, error) {
	achievementEntities := make([]*entities.Achievement, 0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx = tx.Where("game_id = ?", gameID)
	tx = tx.Where("player_id = ?", playerID)
	tx = tx.Find(achievementEntities)
	if err := tx.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return achievementEntities, nil
}

func (a achievementRepository) IsExists(ctx context.Context, id uuid.UUID) (bool, error) {
	count := int64(0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	tx = tx.Model(&entities.Achievement{})
	tx = tx.Where("id = ?", id)
	tx = tx.Count(&count)
	if err := tx.Error; err != nil {
		return false, errors.WithStack(err)
	}

	return count > 0, nil
}

func (a achievementRepository) Insert(ctx context.Context, achievementEntity *entities.Achievement) error {
	achievementEntity.Version = 1

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Create(achievementEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (a achievementRepository) Update(ctx context.Context, achievementEntity *entities.Achievement) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Where("id = ?", achievementEntity.ID)
	tx = tx.Where("version = ?", achievementEntity.Version)

	achievementEntity.Version++

	tx = tx.Updates(achievementEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}
	if rowsAffected := tx.RowsAffected; rowsAffected == 0 {
		return errors.WithStack(persistence.ErrStaleObject)
	}

	return nil
}

func (a achievementRepository) Save(ctx context.Context, achievementEntity *entities.Achievement) error {
	isExists, err := a.IsExists(ctx, achievementEntity.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if isExists {
		if err := a.Update(ctx, achievementEntity); err != nil {
			return errors.WithStack(err)
		}
	} else {
		if err := a.Insert(ctx, achievementEntity); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
