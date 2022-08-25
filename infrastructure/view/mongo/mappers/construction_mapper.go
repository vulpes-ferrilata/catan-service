package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toConstructionView(constructionDocument *documents.Construction) *models.Construction {
	if constructionDocument == nil {
		return nil
	}

	land := toLandView(constructionDocument.Land)

	return &models.Construction{
		ID:   constructionDocument.ID,
		Type: constructionDocument.Type,
		Land: land,
	}
}
