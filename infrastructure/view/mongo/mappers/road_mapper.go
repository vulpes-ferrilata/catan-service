package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toRoadView(roadDocument *documents.Road) *models.Road {
	if roadDocument == nil {
		return nil
	}

	path := toPathView(roadDocument.Path)

	return &models.Road{
		ID:   roadDocument.ID,
		Path: path,
	}
}
