package mappers

import (
	"context"
	"sync"

	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/entities"
	"github.com/pkg/errors"
)

type ResourceCardMapper interface {
	ToEntity(ctx context.Context, resourceCard *models.ResourceCard, game *models.Game, player *models.Player) (*entities.ResourceCard, error)
	ToModel(ctx context.Context, resourceCardEntity *entities.ResourceCard) (*models.ResourceCard, error)
}

func NewResourceCardMapper() ResourceCardMapper {
	return &resourceCardMapper{
		m: make(map[*models.ResourceCard]*entities.ResourceCard),
	}
}

type resourceCardMapper struct {
	m  map[*models.ResourceCard]*entities.ResourceCard
	mu sync.RWMutex
}

func (r *resourceCardMapper) ToEntity(ctx context.Context, resourceCard *models.ResourceCard, game *models.Game, player *models.Player) (*entities.ResourceCard, error) {
	if resourceCard == nil {
		return nil, nil
	}

	r.mu.RLock()
	resourceCardEntity, ok := r.m[resourceCard]
	r.mu.RUnlock()
	if !ok {
		resourceCardEntity = new(entities.ResourceCard)

		r.mu.Lock()
		r.m[resourceCard] = resourceCardEntity
		r.mu.Unlock()

		go func(resourceCard *models.ResourceCard, done <-chan struct{}) {
			<-done
			r.mu.Lock()
			delete(r.m, resourceCard)
			r.mu.Unlock()
		}(resourceCard, ctx.Done())
	}

	resourceCardEntity.ID = resourceCard.GetID()
	resourceCardEntity.GameID = game.GetID()
	if player != nil {
		*resourceCardEntity.PlayerID = player.GetID()
	} else {
		resourceCardEntity.PlayerID = nil
	}
	resourceCardEntity.Type = string(resourceCard.GetType())

	return resourceCardEntity, nil
}

func (r *resourceCardMapper) ToModel(ctx context.Context, resourceCardEntity *entities.ResourceCard) (*models.ResourceCard, error) {
	if resourceCardEntity == nil {
		return nil, nil
	}

	resourceCardType, err := models.NewResourceCardType(resourceCardEntity.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCard := models.NewResourceCard(
		resourceCardEntity.ID,
		resourceCardType,
	)

	r.mu.Lock()
	r.m[resourceCard] = resourceCardEntity
	r.mu.Unlock()

	go func(resourceCard *models.ResourceCard, done <-chan struct{}) {
		<-done
		r.mu.Lock()
		delete(r.m, resourceCard)
		r.mu.Unlock()
	}(resourceCard, ctx.Done())

	return resourceCard, nil
}
