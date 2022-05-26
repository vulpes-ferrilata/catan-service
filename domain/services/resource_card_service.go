package services

import (
	"context"

	"github.com/VulpesFerrilata/catan-service/domain/mappers"
	"github.com/VulpesFerrilata/catan-service/domain/models"
	"github.com/VulpesFerrilata/catan-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service[T any] interface {
	New(ctx context.Context) (T, error)
}

type ResourceCardService interface {
	NewResourceCards(ctx context.Context) ([]*models.ResourceCard, error)
	FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*models.ResourceCard, error)
	Save(ctx context.Context, resourceCard *models.ResourceCard, game *models.Game, player *models.Player) error
}

func NewResourceCardService(resourceCardRepository repositories.ResourceCardRepository,
	resourceCardMapper mappers.ResourceCardMapper) ResourceCardService {
	return &resourceCardService{
		resourceCardRepository: resourceCardRepository,
		resourceCardMapper:     resourceCardMapper,
	}
}

type resourceCardService struct {
	resourceCardRepository repositories.ResourceCardRepository
	resourceCardMapper     mappers.ResourceCardMapper
}

func (r resourceCardService) NewResourceCard(ctx context.Context, resourceCardType models.ResourceCardType) (*models.ResourceCard, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCard := models.NewResourceCard(id, resourceCardType)

	return resourceCard, nil
}

func (r resourceCardService) NewResourceCards(ctx context.Context) ([]*models.ResourceCard, error) {
	resourceCards := make([]*models.ResourceCard, 0)

	longestRoadResourceCard, err := r.NewResourceCard(ctx, models.LongestRoad)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resourceCards = append(resourceCards, longestRoadResourceCard)

	largestArmyResourceCard, err := a.NewResourceCard(ctx, models.LargestArmy)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resourceCards = append(resourceCards, largestArmyResourceCard)

	return resourceCards, nil
}

func (r resourceCardService) FindByGameIDByPlayerID(ctx context.Context, gameID uuid.UUID, playerID *uuid.UUID) ([]*models.ResourceCard, error) {
	resourceCards := make([]*models.ResourceCard, 0)

	resourceCardEntities, err := r.resourceCardRepository.FindByGameIDByPlayerID(ctx, gameID, playerID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, resourceCardEntity := range resourceCardEntities {
		resourceCard, err := r.resourceCardMapper.ToModel(ctx, resourceCardEntity)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		resourceCards = append(resourceCards, resourceCard)
	}

	return resourceCards, nil
}

func (r resourceCardService) Save(ctx context.Context, resourceCard *models.ResourceCard, game *models.Game, player *models.Player) error {
	resourceCardEntity, err := r.resourceCardMapper.ToEntity(ctx, resourceCard, game, player)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.resourceCardRepository.Save(ctx, resourceCardEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
