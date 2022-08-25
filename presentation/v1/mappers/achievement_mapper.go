package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toAchievementResponse(achievement *models.Achievement) *catan.AchievementResponse {
	if achievement == nil {
		return nil
	}

	return &catan.AchievementResponse{
		ID:   achievement.ID.Hex(),
		Type: achievement.Type,
	}
}
