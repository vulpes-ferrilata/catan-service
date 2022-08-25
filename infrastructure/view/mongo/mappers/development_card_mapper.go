package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

func toDevelopmentCardView(developmentCardDocument *documents.DevelopmentCard) *models.DevelopmentCard {
	if developmentCardDocument == nil {
		return nil
	}

	return &models.DevelopmentCard{
		ID:     developmentCardDocument.ID,
		Type:   developmentCardDocument.Type,
		Status: developmentCardDocument.Status,
	}
}
