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

func NewDiceRepository() repositories.DiceRepository {
	return &diceRepository{}
}

type diceRepository struct{}

func (d diceRepository) FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*entities.Dice, error) {
	diceEntities := make([]*entities.Dice, 0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx = tx.Where("game_id = ?", gameID)
	tx = tx.Find(diceEntities)
	if err := tx.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return diceEntities, nil
}

func (d diceRepository) IsExists(ctx context.Context, id uuid.UUID) (bool, error) {
	count := int64(0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	tx = tx.Model(&entities.Dice{})
	tx = tx.Where("id = ?", id)
	tx = tx.Count(&count)
	if err := tx.Error; err != nil {
		return false, errors.WithStack(err)
	}

	return count > 0, nil
}

func (d diceRepository) Insert(ctx context.Context, diceEntity *entities.Dice) error {
	diceEntity.Version = 1

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Create(diceEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (d diceRepository) Update(ctx context.Context, diceEntity *entities.Dice) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Where("id = ?", diceEntity.ID)
	tx = tx.Where("version = ?", diceEntity.Version)

	diceEntity.Version++

	tx = tx.Updates(diceEntity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}
	if rowsAffected := tx.RowsAffected; rowsAffected == 0 {
		return errors.WithStack(persistence.ErrStaleObject)
	}

	return nil
}

func (d diceRepository) Save(ctx context.Context, diceEntity *entities.Dice) error {
	isExists, err := d.IsExists(ctx, diceEntity.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	if isExists {
		if err := d.Update(ctx, diceEntity); err != nil {
			return errors.WithStack(err)
		}
	} else {
		if err := d.Insert(ctx, diceEntity); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
