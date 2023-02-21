package mappers

import (
	pb_models "github.com/vulpes-ferrilata/catan-service-proto/pb/models"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type achievementMapper struct{}

func (a achievementMapper) ToResponse(achievement *models.Achievement) (*pb_models.Achievement, error) {
	if achievement == nil {
		return nil, nil
	}

	return &pb_models.Achievement{
		ID:   achievement.ID.Hex(),
		Type: achievement.Type,
	}, nil
}
