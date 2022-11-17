package mappers

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/catan-service/infrastructure/view/mongo/documents"
	"github.com/vulpes-ferrilata/catan-service/view/models"
)

type terrainMapper struct{}

func (t terrainMapper) ToView(terrainDocument *documents.Terrain) (*models.Terrain, error) {
	if terrainDocument == nil {
		return nil, nil
	}

	harbor, err := harborMapper{}.ToView(terrainDocument.Harbor)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	robber, err := robberMapper{}.ToView(terrainDocument.Robber)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.Terrain{
		ID:     terrainDocument.ID,
		Q:      terrainDocument.Q,
		R:      terrainDocument.R,
		Number: terrainDocument.Number,
		Type:   terrainDocument.Type,
		Harbor: harbor,
		Robber: robber,
	}, nil
}
