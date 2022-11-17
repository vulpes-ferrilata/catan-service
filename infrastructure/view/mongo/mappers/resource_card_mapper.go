package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type resourceCardMapper struct{}

func (r resourceCardMapper) ToView(resourceCardDocument *documents.ResourceCard) (*models.ResourceCard, error) {
	if resourceCardDocument == nil {
		return nil, nil
	}

	return &models.ResourceCard{
		ID:       resourceCardDocument.ID,
		Type:     resourceCardDocument.Type,
		Offering: resourceCardDocument.Offering,
	}, nil
}
