package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toRobberView(robberDocument *documents.Robber) *models.Robber {
	if robberDocument == nil {
		return nil
	}

	return &models.Robber{
		ID: robberDocument.ID,
	}
}
