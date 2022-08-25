package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toPathView(pathDocument *documents.Path) *models.Path {
	if pathDocument == nil {
		return nil
	}

	return &models.Path{
		ID:       pathDocument.ID,
		Q:        pathDocument.Q,
		R:        pathDocument.R,
		Location: pathDocument.Location,
	}
}
