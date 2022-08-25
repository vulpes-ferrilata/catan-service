package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/utils/slices"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toPlayerView(playerDocument *documents.Player) *models.Player {
	if playerDocument == nil {
		return nil
	}

	achievements, _ := slices.Map(func(achievementDocument *documents.Achievement) (*models.Achievement, error) {
		return toAchievementView(achievementDocument), nil
	}, playerDocument.Achievements)

	resourceCards, _ := slices.Map(func(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
		return toResourceCardView(resourceCardDocument), nil
	}, playerDocument.ResourceCards)

	developmentCards, _ := slices.Map(func(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
		return toDevelopmentCardView(developmentCardDocument), nil
	}, playerDocument.DevelopmentCards)

	constructions, _ := slices.Map(func(constructionDocument *documents.Construction) (*models.Construction, error) {
		return toConstructionView(constructionDocument), nil
	}, playerDocument.Constructions)

	roads, _ := slices.Map(func(roadDocument *documents.Road) (*models.Road, error) {
		return toRoadView(roadDocument), nil
	}, playerDocument.Roads)

	return &models.Player{
		ID:               playerDocument.ID,
		UserID:           playerDocument.UserID,
		Color:            playerDocument.Color,
		TurnOrder:        playerDocument.TurnOrder,
		IsOffered:        playerDocument.IsOffered,
		IsActive:         playerDocument.IsActive,
		IsMe:             playerDocument.IsMe,
		Achievements:     achievements,
		ResourceCards:    resourceCards,
		DevelopmentCards: developmentCards,
		Constructions:    constructions,
		Roads:            roads,
	}
}
