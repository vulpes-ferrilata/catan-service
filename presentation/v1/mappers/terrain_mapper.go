package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/catan"
)

func toTerrainResponse(terrain *models.Terrain) *catan.TerrainResponse {
	if terrain == nil {
		return nil
	}

	return &catan.TerrainResponse{
		ID:     terrain.ID.Hex(),
		Q:      int32(terrain.Q),
		R:      int32(terrain.R),
		Number: int32(terrain.Number),
		Type:   terrain.Type,
	}
}
