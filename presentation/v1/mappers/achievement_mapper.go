package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toAchievementResponse(achievement *models.Achievement) *responses.Achievement {
	if achievement == nil {
		return nil
	}

	return &responses.Achievement{
		ID:   achievement.ID.Hex(),
		Type: achievement.Type,
	}
}
