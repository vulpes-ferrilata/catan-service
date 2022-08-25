package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toLandView(landDocument *documents.Land) *models.Land {
	if landDocument == nil {
		return nil
	}

	land := &models.Land{
		ID:       landDocument.ID,
		Q:        landDocument.Q,
		R:        landDocument.R,
		Location: landDocument.Location,
	}

	return land
}
