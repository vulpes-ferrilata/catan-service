package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type robberMapper struct{}

func (r robberMapper) ToView(robberDocument *documents.Robber) (*models.Robber, error) {
	if robberDocument == nil {
		return nil, nil
	}

	return &models.Robber{
		ID: robberDocument.ID,
	}, nil
}
