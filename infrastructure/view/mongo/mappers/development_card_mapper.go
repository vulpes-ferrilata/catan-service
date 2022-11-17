package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type developmentCardMapper struct{}

func (d developmentCardMapper) ToView(developmentCardDocument *documents.DevelopmentCard) (*models.DevelopmentCard, error) {
	if developmentCardDocument == nil {
		return nil, nil
	}

	return &models.DevelopmentCard{
		ID:     developmentCardDocument.ID,
		Type:   developmentCardDocument.Type,
		Status: developmentCardDocument.Status,
	}, nil
}
