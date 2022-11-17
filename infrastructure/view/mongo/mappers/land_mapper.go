package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type landMapper struct{}

func (l landMapper) ToView(landDocument *documents.Land) (*models.Land, error) {
	if landDocument == nil {
		return nil, nil
	}

	return &models.Land{
		ID:       landDocument.ID,
		Q:        landDocument.Q,
		R:        landDocument.R,
		Location: landDocument.Location,
	}, nil
}
