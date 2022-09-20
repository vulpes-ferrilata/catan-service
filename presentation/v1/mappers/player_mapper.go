package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toPlayerResponse(player *models.Player) *catan.PlayerResponse {
	if player == nil {
		return nil
	}

	achievementResponses, _ := slices.Map(func(achievement *models.Achievement) (*catan.AchievementResponse, error) {
		return toAchievementResponse(achievement), nil
	}, player.Achievements)

	resourceCardResponses, _ := slices.Map(func(resourceCard *models.ResourceCard) (*catan.ResourceCardResponse, error) {
		return toResourceCardResponse(resourceCard), nil
	}, player.ResourceCards)

	developmentCardResponses, _ := slices.Map(func(developmentCard *models.DevelopmentCard) (*catan.DevelopmentCardResponse, error) {
		return toDevelopmentCardResponse(developmentCard), nil
	}, player.DevelopmentCards)

	constructionResponses, _ := slices.Map(func(construction *models.Construction) (*catan.ConstructionResponse, error) {
		return toConstructionResponse(construction), nil
	}, player.Constructions)

	roadResponses, _ := slices.Map(func(road *models.Road) (*catan.RoadResponse, error) {
		return toRoadResponse(road), nil
	}, player.Roads)

	return &catan.PlayerResponse{
		ID:               player.ID.Hex(),
		UserID:           player.UserID.Hex(),
		Color:            player.Color,
		TurnOrder:        int32(player.TurnOrder),
		IsOffered:        player.IsOffered,
		IsActive:         player.IsActive,
		Score:            int32(player.Score),
		Achievements:     achievementResponses,
		ResourceCards:    resourceCardResponses,
		DevelopmentCards: developmentCardResponses,
		Constructions:    constructionResponses,
		Roads:            roadResponses,
	}
}
