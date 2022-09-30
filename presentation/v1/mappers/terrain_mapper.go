package mappers

import (
	"github.com/vulpes-ferrilata/catan-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toTerrainResponse(terrain *models.Terrain) *responses.Terrain {
	if terrain == nil {
		return nil
	}

	harborResponse := toHarborResponse(terrain.Harbor)

	robberResponse := toRobberResponse(terrain.Robber)

	return &responses.Terrain{
		ID:     terrain.ID.Hex(),
		Q:      int32(terrain.Q),
		R:      int32(terrain.R),
		Number: int32(terrain.Number),
		Type:   terrain.Type,
		Harbor: harborResponse,
		Robber: robberResponse,
	}
}
