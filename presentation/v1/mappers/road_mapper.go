package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toRoadResponse(road *models.Road) *responses.Road {
	if road == nil {
		return nil
	}

	path := toPathResponse(road.Path)

	return &responses.Road{
		ID:   road.ID.Hex(),
		Path: path,
	}
}
