package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toHarborView(harborDocument *documents.Harbor) *models.Harbor {
	if harborDocument == nil {
		return nil
	}

	return &models.Harbor{
		ID:        harborDocument.ID,
		TerrainID: harborDocument.TerrainID,
		Q:         harborDocument.Q,
		R:         harborDocument.R,
		Type:      harborDocument.Type,
	}
}
