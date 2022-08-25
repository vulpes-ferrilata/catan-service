package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toAchievementView(achievementDocument *documents.Achievement) *models.Achievement {
	if achievementDocument == nil {
		return nil
	}

	return &models.Achievement{
		ID:   achievementDocument.ID,
		Type: achievementDocument.Type,
	}
}
