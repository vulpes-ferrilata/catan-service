package services

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/domain/mappers"
	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type DiceService interface {
	NewDices(ctx context.Context) ([]*models.Dice, error)
	FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*models.Dice, error)
	Save(ctx context.Context, dice *models.Dice, game *models.Game) error
}

func NewDiceService(diceRepository repositories.DiceRepository,
	diceMapper mappers.DiceMapper) DiceService {
	return &diceService{
		diceRepository: diceRepository,
		diceMapper:     diceMapper,
	}
}

type diceService struct {
	diceRepository repositories.DiceRepository
	diceMapper     mappers.DiceMapper
}

func (d diceService) NewDice(ctx context.Context) (*models.Dice, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dice := models.NewDice(id, 1)

	return dice, nil
}

func (d diceService) NewDices(ctx context.Context) ([]*models.Dice, error) {
	dices := make([]*models.Dice, 0)

	for i := 1; i <= 2; i++ {
		dice, err := d.NewDice(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		dices = append(dices, dice)
	}

	return dices, nil
}

func (d diceService) FindByGameID(ctx context.Context, gameID uuid.UUID) ([]*models.Dice, error) {
	dices := make([]*models.Dice, 0)

	diceEntities, err := d.diceRepository.FindByGameID(ctx, gameID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, diceEntity := range diceEntities {
		dice, err := d.diceMapper.ToModel(ctx, diceEntity)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		dices = append(dices, dice)
	}

	return dices, nil
}

func (d diceService) Save(ctx context.Context, dice *models.Dice, game *models.Game) error {
	diceEntity, err := d.diceMapper.ToEntity(ctx, dice, game)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := d.diceRepository.Save(ctx, diceEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
