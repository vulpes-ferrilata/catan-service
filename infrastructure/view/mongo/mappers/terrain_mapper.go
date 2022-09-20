package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toTerrainView(terrainDocument *documents.Terrain) *models.Terrain {
	if terrainDocument == nil {
		return nil
	}

	harbor := toHarborView(terrainDocument.Harbor)

	robber := toRobberView(terrainDocument.Robber)

	return &models.Terrain{
		ID:     terrainDocument.ID,
		Q:      terrainDocument.Q,
		R:      terrainDocument.R,
		Number: terrainDocument.Number,
		Type:   terrainDocument.Type,
		Harbor: harbor,
		Robber: robber,
	}
}
