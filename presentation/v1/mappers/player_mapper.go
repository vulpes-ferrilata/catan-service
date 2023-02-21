package mappers

import (
	"github.com/pkg/errors"
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/slices"
)

type playerMapper struct{}

func (p playerMapper) ToResponse(player *models.Player) (*pb_models.Player, error) {
	if player == nil {
		return nil, nil
	}

	achievementResponses, err := slices.Map(func(achievement *models.Achievement) (*pb_models.Achievement, error) {
		return achievementMapper{}.ToResponse(achievement)
	}, player.Achievements...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resourceCardResponses, err := slices.Map(func(resourceCard *models.ResourceCard) (*pb_models.ResourceCard, error) {
		return resourceCardMapper{}.ToResponse(resourceCard)
	}, player.ResourceCards...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	developmentCardResponses, err := slices.Map(func(developmentCard *models.DevelopmentCard) (*pb_models.DevelopmentCard, error) {
		return developmentCardMapper{}.ToResponse(developmentCard)
	}, player.DevelopmentCards...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	constructionResponses, err := slices.Map(func(construction *models.Construction) (*pb_models.Construction, error) {
		return constructionMapper{}.ToResponse(construction)
	}, player.Constructions...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	roadResponses, err := slices.Map(func(road *models.Road) (*pb_models.Road, error) {
		return roadMapper{}.ToResponse(road)
	}, player.Roads...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &pb_models.Player{
		ID:                 player.ID.Hex(),
		UserID:             player.UserID.Hex(),
		Color:              player.Color,
		TurnOrder:          int32(player.TurnOrder),
		ReceivedOffer:      player.ReceivedOffer,
		DiscardedResources: player.DiscardedResources,
		Score:              int32(player.Score),
		Achievements:       achievementResponses,
		ResourceCards:      resourceCardResponses,
		DevelopmentCards:   developmentCardResponses,
		Constructions:      constructionResponses,
		Roads:              roadResponses,
	}, nil
}
