package mappers

import (
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type pathMapper struct{}

func (p pathMapper) ToView(pathDocument *documents.Path) (*models.Path, error) {
	if pathDocument == nil {
		return nil, nil
	}

	return &models.Path{
		ID:       pathDocument.ID,
		Q:        pathDocument.Q,
		R:        pathDocument.R,
		Location: pathDocument.Location,
	}, nil
}
