package mappers

import (
	"context"
	"sync"

	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/entities"
)

type DiceMapper interface {
	ToEntity(ctx context.Context, dice *models.Dice, game *models.Game) (*entities.Dice, error)
	ToModel(ctx context.Context, diceEntity *entities.Dice) (*models.Dice, error)
}

func NewDiceMapper() DiceMapper {
	return &diceMapper{
		m: make(map[*models.Dice]*entities.Dice),
	}
}

type diceMapper struct {
	m  map[*models.Dice]*entities.Dice
	mu sync.RWMutex
}

func (d *diceMapper) ToEntity(ctx context.Context, dice *models.Dice, game *models.Game) (*entities.Dice, error) {
	if dice == nil {
		return nil, nil
	}

	d.mu.RLock()
	diceEntity, ok := d.m[dice]
	d.mu.RUnlock()
	if !ok {
		diceEntity = new(entities.Dice)

		d.mu.Lock()
		d.m[dice] = diceEntity
		d.mu.Unlock()

		go func(dice *models.Dice, done <-chan struct{}) {
			<-done
			d.mu.Lock()
			delete(d.m, dice)
			d.mu.Unlock()
		}(dice, ctx.Done())
	}

	diceEntity.ID = dice.GetID()
	diceEntity.GameID = game.GetID()
	diceEntity.Number = dice.GetNumber()

	return diceEntity, nil
}

func (d *diceMapper) ToModel(ctx context.Context, diceEntity *entities.Dice) (*models.Dice, error) {
	if diceEntity == nil {
		return nil, nil
	}

	dice := models.NewDice(
		diceEntity.ID,
		diceEntity.Number,
	)

	d.mu.Lock()
	d.m[dice] = diceEntity
	d.mu.Unlock()

	go func(dice *models.Dice, done <-chan struct{}) {
		<-done
		d.mu.Lock()
		delete(d.m, dice)
		d.mu.Unlock()
	}(dice, ctx.Done())

	return dice, nil
}
