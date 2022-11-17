package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type playerMapper struct{}

func (p playerMapper) ToView(playerDocument *documents.Player) (*models.Player, error) {
	if playerDocument == nil {
		return nil, nil
	}

	achievements, err := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return achievementMapper{}.ToView(achievementDocument)
	}, playerDocument.Achievements)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCards, err := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return resourceCardMapper{}.ToView(resourceCardDocument)
	}, playerDocument.ResourceCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCards, err := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return developmentCardMapper{}.ToView(developmentCardDocument)
	}, playerDocument.DevelopmentCards)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	constructions, err := slices.Map(func(constructionDocument *documents.Construction) (*models.Construction, error) {
		return constructionMapper{}.ToView(constructionDocument)
	}, playerDocument.Constructions)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	roads, err := slices.Map(func(roadDocument *documents.Road) (*models.Road, error) {
		return roadMapper{}.ToView(roadDocument)
	}, playerDocument.Roads)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.Player{
		ID:                 playerDocument.ID,
		UserID:             playerDocument.UserID,
		Color:              playerDocument.Color,
		TurnOrder:          playerDocument.TurnOrder,
		ReceivedOffer:      playerDocument.ReceivedOffer,
		DiscardedResources: playerDocument.DiscardedResources,
		Score:              playerDocument.Score,
		Achievements:       achievements,
		ResourceCards:      resourceCards,
		DevelopmentCards:   developmentCards,
		Constructions:      constructions,
		Roads:              roads,
	}, nil
}
