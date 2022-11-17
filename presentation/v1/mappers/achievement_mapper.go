package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type achievementMapper struct{}

func (a achievementMapper) ToResponse(achievement *models.Achievement) (*responses.Achievement, error) {
	if achievement == nil {
		return nil, nil
	}

	return &responses.Achievement{
		ID:   achievement.ID.Hex(),
		Type: achievement.Type,
	}, nil
}
