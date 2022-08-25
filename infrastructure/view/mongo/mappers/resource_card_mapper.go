package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toResourceCardView(resourceCardDocument *documents.ResourceCard) *models.ResourceCard {
	if resourceCardDocument == nil {
		return nil
	}

	return &models.ResourceCard{
		ID:         resourceCardDocument.ID,
		Type:       resourceCardDocument.Type,
		IsSelected: resourceCardDocument.IsSelected,
	}
}
