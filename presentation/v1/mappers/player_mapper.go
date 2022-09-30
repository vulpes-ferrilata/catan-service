package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toPlayerResponse(player *models.Player) *responses.Player {
	if player == nil {
		return nil
	}

	achievementResponses, _ := slices.Map(func(achievement *models.Achievement) (*responses.Achievement, error) {
		return toAchievementResponse(achievement), nil
	}, player.Achievements)

	resourceCardResponses, _ := slices.Map(func(resourceCard *models.ResourceCard) (*responses.ResourceCard, error) {
		return toResourceCardResponse(resourceCard), nil
	}, player.ResourceCards)

	developmentCardResponses, _ := slices.Map(func(developmentCard *models.DevelopmentCard) (*responses.DevelopmentCard, error) {
		return toDevelopmentCardResponse(developmentCard), nil
	}, player.DevelopmentCards)

	constructionResponses, _ := slices.Map(func(construction *models.Construction) (*responses.Construction, error) {
		return toConstructionResponse(construction), nil
	}, player.Constructions)

	roadResponses, _ := slices.Map(func(road *models.Road) (*responses.Road, error) {
		return toRoadResponse(road), nil
	}, player.Roads)

	return &responses.Player{
		ID:               player.ID.Hex(),
		UserID:           player.UserID.Hex(),
		Color:            player.Color,
		TurnOrder:        int32(player.TurnOrder),
		IsOffered:        player.IsOffered,
		Score:            int32(player.Score),
		Achievements:     achievementResponses,
		ResourceCards:    resourceCardResponses,
		DevelopmentCards: developmentCardResponses,
		Constructions:    constructionResponses,
		Roads:            roadResponses,
	}
}
