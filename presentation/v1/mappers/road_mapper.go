package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toRoadResponse(road *models.Road) *catan.RoadResponse {
	if road == nil {
		return nil
	}

	path := toPathResponse(road.Path)

	return &catan.RoadResponse{
		ID:   road.ID.Hex(),
		Path: path,
	}
}
